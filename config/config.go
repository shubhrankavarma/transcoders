package config

type Properties struct {
	Port                  string `env:"TRANSCODERS_SRV_PORT" env-default:"51000"`
	Host                  string `env:"TRANSCODERS_HOST" env-default:"localhost"`
	DBUser                string `env:"DB_USER" env-default:"Ananth"`
	DBPass                string `env:"DB_PASS" env-default:"Ananth1982"`
	DBName                string `env:"DB_NAME" env-default:"OnDemand"`
	DBURL                 string `env:"DB_URL" env-default:"mongodb+srv://%s:%s@anakordb.azocimz.mongodb.net/?retryWrites=true&w=majority"`
	TranscodersCollection string `env:"TRANSCODERS_COLLECTION" env-default:"transcoders"`
	JwtTokenSecret        string `env:"JWT_TOKEN_SECRET" env-default:"CheeseCherryDanishCake"`
}
