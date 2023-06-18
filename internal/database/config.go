package database

type DBConfig struct {
	PGAddress  string
	PGUser     string
	PGPassword string
	PGDatabase string
}

func GetDefaultDBConfig() *DBConfig {
	return &DBConfig{
		PGAddress:  ":5432",
		PGUser:     "postgres",
		PGPassword: "postgres",
		PGDatabase: "postgres",
	}
}
