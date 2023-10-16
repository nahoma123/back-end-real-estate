package constant

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DATABASE_URL           string
	DATABASE_NAME          string
	DATABASE_PASSWORD_HASH string
	JWT_SECRET             string
	SERVER_PORT            string
	SERVER_TIMEOUT         string
	CLOUDINARY_URL         string
	CLOUDINARY_FOLDER      string

	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	PORT        string
}

func GetConfig() Config {

	config := Config{}
	config.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	if config.DATABASE_NAME == "" {
		// log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DATABASE_NAME"))
	}
	config.DATABASE_URL = os.Getenv("DATABASE_URL")
	if config.DATABASE_URL == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DATABASE_URL"))
	}

	// config.DATABASE_PASSWORD_HASH = os.Getenv("DATABASE_PASSWORD_HASH")
	// if config.DATABASE_PASSWORD_HASH == "" {
	// 	log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DATABASE_PASSWORD_HASH"))
	// }

	config.JWT_SECRET = os.Getenv("JWT_SECRET")
	config.SERVER_PORT = os.Getenv("SERVER_PORT")
	if config.SERVER_PORT == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "SERVER_PORT"))
	}

	config.SERVER_TIMEOUT = os.Getenv("SERVER_TIMEOUT")
	if config.SERVER_TIMEOUT == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "SERVER_TIMEOUT"))
	}

	config.DB_HOST = os.Getenv("DB_HOST")
	if config.CLOUDINARY_FOLDER == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DB_HOST"))
	}

	config.DB_USER = os.Getenv("DB_USER")
	if config.CLOUDINARY_FOLDER == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DB_USER"))
	}

	config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if config.CLOUDINARY_FOLDER == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DB_PASSWORD"))
	}

	config.PORT = os.Getenv("PORT")
	if config.CLOUDINARY_FOLDER == "" {
		log.Fatal(context.Background(), fmt.Sprintf("environment variable not found %s", "DB_PASSWORD"))
	}

	return config
}
