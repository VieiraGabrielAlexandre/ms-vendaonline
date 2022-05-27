package configs

import "os"

func Environment() {
	os.Setenv("APP_ENV", "local")
	os.Setenv("DB_HOST", "xplosao-aws.cp0xzrhtdnog.us-east-1.rds.amazonaws.com")
	os.Setenv("DB_PASSWORD", "shO9ocjc19TM")
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_DATABASE", "ms-vendaonline")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PORT", "3306")
}
