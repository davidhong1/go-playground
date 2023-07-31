package main02

import (
	"crypto/tls"
	"function-options/model"
	"time"
)

// 使用 Builder 模式
// 不需要 Config 类，只需要多加一个 Builder 类
// 你可能觉得，这个 Builder 类似乎有点多余，
// 我们似乎可以直接在 Server 上进行这样的 Builder 构造，的确是这样。
// 但是，在处理错误的时候可能有点麻烦，不如包装类更好
type ServerBuilder struct {
	model.Server2
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server2.Addr = addr
	sb.Server2.Port = port
	return sb
}

func (sb *ServerBuilder) WithProtocal(protocol string) *ServerBuilder {
	sb.Server2.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithMaxConn(maxconn int) *ServerBuilder {
	sb.Server2.Maxconns = maxconn
	return sb
}

func (sb *ServerBuilder) WithTimeOut(timeout time.Duration) *ServerBuilder {
	sb.Server2.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server2.TLS = tls
	return sb
}

func (sb *ServerBuilder) Build() model.Server2 {
	return sb.Server2
}
