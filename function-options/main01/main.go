package main01

import "function-options/model"

func NewServer(addr string, port int, conf *model.Config) (*model.Server, error) {
	// 这段代码算不错的了，大多数情况下，我们可能就止步于此。但是，对于有洁癖、有追求的程序员来说，
	// 他们会看到其中不太好的一点，那就是 Config 并不是必需的，所以，你需要判断是否是 nil 或者 Empty
	// Config{} 会让我们的代码感觉不太干净
	// 这段代码算不错的了，大多数情况下，我们可能就止步于此。但是，对于有洁癖、有追求的程序员来说，
	// 他们会看到其中不太好的一点，那就是 Config 并不是必需的，所以，你需要判断是否是 nil 或者 Empty
	// Config{} 会让我们的代码感觉不太干净
	return &model.Server{
		Addr: addr,
		Port: port,
		Conf: conf,
	}, nil
}
