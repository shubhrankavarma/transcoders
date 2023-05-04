package config

type Properties struct {
	Port                  string `env:"TRANSCODERS_SRV_PORT" env-default:"8001" hashicorp:"TRANSCODERS_SRV_PORT"`
	Host                  string `env:"TRANSCODERS_HOST" env-default:"localhost" hashicorp:"TRANSCODERS_HOST"`
	DBUser                string `env:"DB_USER" env-default:"Ananth" hashicorp:"DB_USER"`
	DBPass                string `env:"DB_PASS" env-default:"Ananth1982" hashicorp:"DB_PASS"`
	DBName                string `env:"DB_NAME" env-default:"OnDemand" hashicorp:"DB_NAME"`
	DBURL                 string `env:"DB_URL" env-default:"mongodb+srv://%s:%s@anakordb.azocimz.mongodb.net/?retryWrites=true&w=majority" hashicorp:"DB_URL"`
	TranscodersCollection string `env:"TRANSCODERS_COLLECTION" env-default:"transcoders" hashicorp:"TRANSCODERS_COLLECTION"`
	PageNo                int    `env:"PAGE_NO" env-default:"1" hashicorp:"PAGE_NO"`
	PageSize              int    `env:"PAGE_SIZE" env-default:"15" hashicorp:"PAGE_SIZE"`
	JwtTokenSecret        string `env:"JWT_TOKEN_SECRET" env-default:"CheeseCherryDanishCake" hashicorp:"JWT_TOKEN_SECRET"`
}
