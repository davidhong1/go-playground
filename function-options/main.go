package main

import (
	"fmt"
	"function-options/main01"
	"function-options/main02"
	"function-options/main03"
	"function-options/model"
	"time"
)

func main() {
	srv1, _ := main01.NewServer("localhost", 9000, nil)
	fmt.Printf("%+v\n", srv1)

	conf := model.Config{
		Protocol: "tcp",
		Timeout:  60 * time.Second,
	}
	srv2, _ := main01.NewServer("localhost", 9000, &conf)
	fmt.Printf("%+v\n", srv2)

	sb := main02.ServerBuilder{}
	srv3 := sb.Create("127.0.0.1", 9000).
		WithProtocal("udp").
		WithMaxConn(1024).
		Build()
	fmt.Printf("%+v\n", srv3)

	srv4, _ := main03.NewServer("localhost", 1024)
	fmt.Printf("%+v\n", srv4)
	srv5, _ := main03.NewServer("localhost", 1024, main03.Protocol("udp"))
	fmt.Printf("%+v\n", srv5)
	srv6, _ := main03.NewServer("localhost", 1024, main03.Timeout(300*time.Second), main03.Maxconns(1000))
	fmt.Printf("%+v\n", srv6)
}
