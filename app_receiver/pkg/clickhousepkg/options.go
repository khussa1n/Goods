package clickhousepkg

type Option func(*Ch)

func WithHost(host string) Option {
	return func(postgres *Ch) {
		postgres.host = host
	}
}

func WithPort(port string) Option {
	return func(postgres *Ch) {
		postgres.port = port
	}
}

func WithUsername(username string) Option {
	return func(postgres *Ch) {
		postgres.username = username
	}
}

func WithPassword(password string) Option {
	return func(postgres *Ch) {
		postgres.password = password
	}
}

func WithDBName(dbName string) Option {
	return func(postgres *Ch) {
		postgres.dbName = dbName
	}
}
