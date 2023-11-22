package core

type (
	HexConfig struct {
		Config HexConfigOptions
	}

	Backend struct {
		Dialect       Dialect
		Username      *string
		Password      *string
		ConnectionUrl *string
		Port          *int32
		Host          *string
	}

	HexConfigOptions struct {
		Connection Backend
	}

	Dialect string
)

const (
	REDIS  Dialect = "redis"
	MONGO  Dialect = "mongo"
	DYNAMO Dialect = "dynamo"
	MYSQL  Dialect = "mysql"
)
