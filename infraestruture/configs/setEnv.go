package configs

import "os"

func Environment() {
	os.Setenv("APP_ENV", "local")

	if os.Getenv("APP_ENV") == "production" {
		os.Setenv("DB_HOST", "xplosao-aws.cp0xzrhtdnog.us-east-1.rds.amazonaws.com")
		os.Setenv("DB_PASSWORD", "shO9ocjc19TM")
		os.Setenv("DB_DRIVER", "mysql")
		os.Setenv("DB_DATABASE", "ms-vendaonline")
		os.Setenv("DB_USER", "admin")
		os.Setenv("DB_PORT", "3306")
		os.Setenv("SERVER_PORT", "80")
		return
	}

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PASSWORD", "vendaonline")
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_DATABASE", "ms-vendaonline")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PORT", "3307")
	os.Setenv("SERVER_PORT", "5566")
}
