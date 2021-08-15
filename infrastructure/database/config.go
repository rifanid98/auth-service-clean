package database

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	// ctxTimeout time.Duration
}

func newConfigPostgres() *config {
	return &config{
		host:     "localhost",
		database: "okedok_golang",
		port:     "5433",
		driver:   "postgres",
		user:     "postgres",
		password: "postgres",
	}
}
