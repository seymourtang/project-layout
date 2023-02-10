package option

import (
	"flag"

	"github.com/google/wire"
	"github.com/jinzhu/configor"

	"github.com/seymourtang/project-layout/internal/data/cache/redis"
	"github.com/seymourtang/project-layout/internal/data/db"
	"github.com/seymourtang/project-layout/internal/server/http"
)

var ProvideSet = wire.NewSet(
	New,
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
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Redis struct {
		Addrs    []string
		DB       uint
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"6379"`
		PoolSize int
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}

func New() (*Option, error) {
	var option Option

	config := flag.String("file", "config.yml", "configuration file")
	flag.StringVar(&option.APPName, "name", "", "app name")
	flag.StringVar(&option.MySQL.DBName, "db-name", "", "database name")
	flag.StringVar(&option.MySQL.Username, "db-user", "root", "database user")
	flag.Parse()

	if err := configor.Load(&option, *config); err != nil {
		return nil, err
	}

	return &option, nil
}

func NewRedisConfig(o *Option) []redis.Option {
	return []redis.Option{
		redis.WithDB(int(o.Redis.DB)),
		redis.WithPassword(o.Redis.Password),
		redis.WithEndpoint(o.Redis.Addrs),
		redis.WithPoolSize(o.Redis.PoolSize),
	}
}

func NewMySQLConfig(o *Option) []db.Option {
	return []db.Option{
		db.WithDatabase(o.MySQL.DBName),
		db.WithHost(o.MySQL.Host),
		db.WithUsername(o.MySQL.Username),
		db.WithPassword(o.MySQL.Password),
	}
}

func NewHttpConfig(o *Option) *http.Option {
	return &http.Option{Port: o.Http.Port}
}
