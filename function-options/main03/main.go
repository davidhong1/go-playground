package main03

import (
	"crypto/tls"
	"function-options/model"
	"time"
)

// Function Options
type Option func(*model.Server2)

func Protocol(p string) Option {
	return func(s *model.Server2) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *model.Server2) {
		s.Timeout = timeout
	}
}

func Maxconns(maxconns int) Option {
	return func(s *model.Server2) {
		s.Maxconns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *model.Server2) {
		s.TLS = tls
	}
}

// 所以，以后，你要玩类似的代码时，我强烈推荐你使用 Functional Options 这种方式，
// 这种方式至少带来了 6 个好处:
// - 直觉式的编程;
// - 高度的可配置化;
// - 很容易维护和扩展;
// - 自文档;
// - 新来的人很容易上手;
// - 没有什么令人困惑的事(是 nil 还是空)
func NewServer(addr string, port int, options ...Option) (*model.Server2, error) {
	srv := model.Server2{
		Addr: addr,
		Port: port,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}
