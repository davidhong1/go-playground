package main

import (
	"context"
	"fmt"
	"go-study-lib/log"
	"os"
	"os/signal"
	"week02/dao"
)

/*
问:
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
答:
需要。假如是业务层调用了dao层，那么这个sql.ErrNoRows对于业务层来说则是其他层或其他module的错误，并且是本层第一次遇到其他module的错误，
所以应该在sql.ErrNoRows基础上面wrap一些额外信息，比如堆栈信息、当前上下文信息、sql和sql参数信息等等。
*/

var dos = []dao.DoSomething{
	{
		Who:   "davidhong",
		Thing: "learn go",
	}, {
		Who:   "davidhong",
		Thing: "have lunch",
	},
	{
		Who:   "davidhong",
		Thing: "do something",
	},
}

func main() {

	c := make(chan os.Signal)
	signal.Notify(c)

	ctx := context.TODO()

	// init sql connector
	err := dao.InitDB()
	if err != nil {
		panic(fmt.Sprintf("InitDB fail. err: %v", err))
	}

	// init data
	for _, do := range dos {
		err := do.Delete(ctx)
		if err != nil {
			log.ERROR("delete %v fail.", do)
		}
	}
	for _, do := range dos[:len(dos)-1] {
		err := do.Insert(ctx)
		if err != nil {
			log.ERROR("delete %v fail.", do)
		}
	}

	// deal sql.ErrNoRows example
	for _, do := range dos {
		err := GetDoSomething(ctx, &do)
		if err != nil {
			log.ERROR("Get %v err. err: %+v", do, err)
		}
		log.INFO("GetDoSomething %v", do)
	}

	// hold main
	for {
		select {
		case <-ctx.Done():
			log.ERROR("exit by context done.")
			dao.CloseDB()
			os.Exit(1)
		case <-c:
			log.INFO("exit by signal %v.", c)
			dao.CloseDB()
			os.Exit(0)
		}
	}
}

func GetDoSomething(ctx context.Context, do *dao.DoSomething) error {
	err := do.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
