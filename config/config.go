// Package config holds constants + configuration
package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getEnvFileName() string {
	if getEnv() == Const().EnvTest {
		return Const().EnvFileTest
	}
	return Const().EnvFileLocal
}

func init() {
	LoadEnvIfNeeded()
	loadConfig()
}

func LoadEnvIfNeeded() {
	if runningTestOrLocal() {
		name := getEnvFileName()
		err := godotenv.Load(name)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func runningTestOrLocal() bool {
	fmt.Printf("Running in enviroment: %s \n", os.Getenv("ENV"))
	return !(getEnv() == Const().EnvProd)
}

func getEnv() string {
	return os.Getenv("ENV")
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
	if getEnv() == Const().EnvProd {
		return Const().UsersFileCloud
	}

	return Const().UsersFileLocal
}

// Conf returns the Config struct
func Conf() Config {
	return conf
}
