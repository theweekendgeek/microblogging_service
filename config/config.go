package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getEnvFileName() string {
	if isTest() {
		return Const().EnvFileTest
	}
	return Const().EnvFileLocal
}

func init() {
	if !IsProd() {
		name := getEnvFileName()
		err := godotenv.Load(name)
		if err != nil {
			log.Fatal(err)
		}
	}
	loadConfig()
}

func setEnvironment() string {
	env := os.Getenv("ENV")
	if env == Const().EnvLocal || env == Const().EnvTest {
		return env
	}
	return Const().EnvProd
}

type Config struct {
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
	Bearer   string
	Env      string
	FilePath string
}

var conf Config

func loadConfig() {
	conf = Config{
		DbHost:   os.Getenv("DB_HOST"),
		DbPort:   os.Getenv("DB_PORT"),
		DbUser:   os.Getenv("POSTGRES_USER"),
		DbPass:   os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Bearer:   os.Getenv("BEARER"),
		FilePath: setUserFilePath(),
		Env:      setEnvironment(),
	}
}

func setUserFilePath() string {
	if IsProd() {
		return Const().UsersFileCloud
	}

	return Const().UsersFileLocal
}

func Conf() Config {
	return conf
}
