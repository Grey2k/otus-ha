package conf

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultServerPort = "8080"
	defaultDbPort     = "3306"
	defaultDbName     = "otus"
	defaultDbOpenConn = 10
	defaultDbIdleConn = 10
	defaultDbConnLife = time.Minute * 2
	defaultReqTimeout = time.Second * 2
	defaultJwtTtl     = time.Hour * 1
	defaultBusTopic   = "user.events"
)

// Config struct holds application's parameters
type Config struct {
	DbUser           string
	DbPassword       string
	DbHost           string
	DbPort           string
	DbName           string
	DbMaxOpenConn    int
	DbMaxIdleConn    int
	DbConnMaxLife    time.Duration
	DbMigrationsPath string
	BusDSN           string
	BusTopic         string

	NewsServiceToken string
	NewsServiceUrl   string

	ServerPort     string
	RequestTimeout time.Duration // in seconds

	JwtSecret string
	JwtTtl    time.Duration

	DebugMode bool
}

func PopulateConfig() (*Config, error) {
	var (
		cfg   Config
		exist bool
	)

	if cfg.JwtSecret, exist = os.LookupEnv("JWT_SECRET"); !exist {
		return nil, errors.New("ENV `JWT_SECRET` should be specified")
	}
	cfg.JwtTtl = defaultJwtTtl

	if cfg.DbHost, exist = os.LookupEnv("DB_HOST"); !exist {
		return nil, errors.New("ENV `DB_HOST` should be specified")
	}
	if cfg.DbUser, exist = os.LookupEnv("DB_USER"); !exist {
		return nil, errors.New("ENV `DB_USER` should be specified")
	}
	if cfg.DbPassword, exist = os.LookupEnv("DB_PASSWORD"); !exist {
		return nil, errors.New("ENV `DB_PASSWORD` should be specified")
	}

	if cfg.DbName, exist = os.LookupEnv("DB_NAME"); !exist {
		cfg.DbName = defaultDbName
	}
	if cfg.DbPort, exist = os.LookupEnv("DB_PORT"); !exist {
		cfg.DbPort = defaultDbPort
	}

	cfg.DbMaxOpenConn = defaultDbOpenConn
	cfg.DbMaxIdleConn = defaultDbIdleConn
	cfg.DbConnMaxLife = defaultDbConnLife

	if cfg.DbMigrationsPath, exist = os.LookupEnv("DB_MIGRATIONS_PATH"); !exist {
		cfg.DbMigrationsPath = filepath.Dir(filepath.Dir(os.Args[0])) + "/backend/migrations"
	}

	if cfg.ServerPort, exist = os.LookupEnv("SERVER_PORT"); !exist {
		cfg.ServerPort = defaultServerPort
	}

	if cfg.BusDSN, exist = os.LookupEnv("BUS_DSN"); !exist {
		return nil, errors.New("ENV `BUS_DSN` should be specified")
	}
	if cfg.BusTopic, exist = os.LookupEnv("BUS_TOPIC"); !exist {
		cfg.BusTopic = defaultBusTopic
	}

	if cfg.NewsServiceToken, exist = os.LookupEnv("NEWS_TOKEN"); !exist {
		return nil, errors.New("ENV `NEWS_TOKEN` should be specified")
	}
	if cfg.NewsServiceUrl, exist = os.LookupEnv("NEWS_URL"); !exist {
		return nil, errors.New("ENV `NEWS_URL` should be specified")
	}

	cfg.RequestTimeout = defaultReqTimeout

	tmp, exist := os.LookupEnv("DEBUG")
	cfg.DebugMode = exist && tmp == "true"

	return &cfg, nil
}
