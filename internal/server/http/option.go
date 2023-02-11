package http

import "time"

type Option struct {
	Port         uint
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
