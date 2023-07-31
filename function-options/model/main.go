package model

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Server struct {
	Addr string
	Port int
	Conf *Config
}

type Server2 struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}
