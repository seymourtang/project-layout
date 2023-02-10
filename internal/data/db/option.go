package db

import "fmt"

type options struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type Option interface {
	apply(*options)
}

type UsernameOption string

func (u UsernameOption) apply(o *options) {
	o.Username = string(u)
}

func WithUsername(username string) Option {
	return UsernameOption(username)
}

type PasswordOption string

func (u PasswordOption) apply(o *options) {
	o.Password = string(u)
}

func WithPassword(Password string) Option {
	return PasswordOption(Password)
}

type HostOption string

func (u HostOption) apply(o *options) {
	o.Host = string(u)
}

func WithHost(host string) Option {
	return HostOption(host)
}

type PortOption int

func (u PortOption) apply(o *options) {
	o.Port = int(u)
}

func WithPort(port int) Option {
	return PortOption(port)
}

type DatabaseOption string

func (u DatabaseOption) apply(o *options) {
	o.Database = string(u)
}

func WithDatabase(database string) Option {
	return DatabaseOption(database)
}

func (o options) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", o.Username, o.Password, o.Host, o.Port, o.Database)
}
