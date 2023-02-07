package redis

type redisOptions struct {
	DB       int
	Password string
	Addrs    []string
	Port     int
	PoolSize int
}

type Option interface {
	apply(*redisOptions)
}

type dbOption int

func (D dbOption) apply(o *redisOptions) {
	o.DB = int(D)
}

func WithDB(db int) Option {
	return dbOption(db)
}

type passwordOption string

func (p passwordOption) apply(o *redisOptions) {
	o.Password = string(p)
}

func WithPassword(password string) Option {
	return passwordOption(password)
}

type addrsOption []string

func (a addrsOption) apply(o *redisOptions) {
	o.Addrs = a
}

func WithEndpoint(addrs []string) Option {
	return addrsOption(addrs)
}

type poolSizeOption int

func (p poolSizeOption) apply(o *redisOptions) {
	o.PoolSize = int(p)
}

func WithPoolSize(poolSize int) Option {
	return poolSizeOption(poolSize)
}
