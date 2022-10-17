package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var conn *gorm.DB

type connectionConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func generateConnectionString(c connectionConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
}

func Connect() error {
	config := connectionConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	}

	var err error
	conn, err = gorm.Open(postgres.Open(generateConnectionString(config)), &gorm.Config{})
	return err
}

func Migrate() error {
	return conn.AutoMigrate(&GormProfile{})
}

func GetDb() *gorm.DB {
	return conn
}
