package orm

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func generateConnectionString(c ConnectionConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
}

func init() {
	err := godotenv.Load()

	config := ConnectionConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	}

	dsn := fmt.Sprintf(generateConnectionString(config))
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	fmt.Println(db)

}
