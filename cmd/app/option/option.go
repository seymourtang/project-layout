package option

import (
	"flag"
	"time"

	"github.com/google/wire"
	"github.com/jinzhu/configor"

	"github.com/seymourtang/project-layout/internal/data/cache/redis"
	"github.com/seymourtang/project-layout/internal/data/db"
	"github.com/seymourtang/project-layout/internal/server/http"
)

var ProviderSet = wire.NewSet(
	NewCmd,
	NewRedisConfig,
	NewMySQLConfig,
	NewHttpConfig,
)

type Option struct {
	APPName string `default:"app name"`

	Http struct {
		Port uint
	}
	MySQL struct {
		Host     string
		DBName   string
		Username string `default:"root"`
		Password string `required:"true"`
		Port     uint   `default:"3306"`
	}

	Redis struct {
		Addrs    string
		DB       uint
		Password string `default:"redispwd"`
		PoolSize int
	}

	Contacts []struct {
		Name  string
		Email string
	}
}

func NewCmd() (*Option, error) {
	var option Option

	flag.StringVar(&option.APPName, "name", "", "app name")
	flag.StringVar(&option.MySQL.Host, "db-host", "127.0.0.1", "database host")
	flag.UintVar(&option.MySQL.Port, "db-port", 3306, "database host")
	flag.StringVar(&option.MySQL.Password, "db-password", "123456", "database password")
	flag.StringVar(&option.MySQL.DBName, "db-name", "test_db", "database name")
	flag.StringVar(&option.MySQL.Username, "db-user", "root", "database user")
	flag.StringVar(&option.Redis.Addrs, "redis-addrs", "127.0.0.1:32769", "redis addrs")
	flag.StringVar(&option.Redis.Password, "redis-password", "redispw", "redis password")
	flag.IntVar(&option.Redis.PoolSize, "redis-poolSize", 55, "redis connection pool size")
	flag.UintVar(&option.Http.Port, "http-port", 8099, "http server port")
	flag.Parse()

	if err := configor.Load(&option); err != nil {
		return nil, err
	}

	return &option, nil
}

func NewRedisConfig(o *Option) []redis.Option {
	return []redis.Option{
		redis.WithDB(int(o.Redis.DB)),
		redis.WithPassword(o.Redis.Password),
		redis.WithAddrs(o.Redis.Addrs),
		redis.WithPoolSize(o.Redis.PoolSize),
	}
}

func NewMySQLConfig(o *Option) []db.Option {
	return []db.Option{
		db.WithDatabase(o.MySQL.DBName),
		db.WithHost(o.MySQL.Host),
		db.WithPort(int(o.MySQL.Port)),
		db.WithUsername(o.MySQL.Username),
		db.WithPassword(o.MySQL.Password),
	}
}

func NewHttpConfig(o *Option) *http.Option {
	return &http.Option{
		Port:         o.Http.Port,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}
}
