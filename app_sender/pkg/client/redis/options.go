package redis

type Option func(*Redis)

func WithHost(host string) Option {
	return func(redis *Redis) {
		redis.host = host
	}
}

func WithPort(port string) Option {
	return func(redis *Redis) {
		redis.port = port
	}
}

func WithUsername(username string) Option {
	return func(redis *Redis) {
		redis.username = username
	}
}

func WithPassword(password string) Option {
	return func(redis *Redis) {
		redis.password = password
	}
}

func WithDBName(dbName string) Option {
	return func(redis *Redis) {
		redis.dbName = dbName
	}
}
