package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type env struct {
	DB_HOST     string `validate:"required"`
	DB_USERNAME string `validate:"required"`
	DB_PASSWORD string `validate:"required"`
	DB_DATABASE string `validate:"required"`
	DB_PORT     string `validate:"required"`
	PORT        string `validate:"required"`
	JWT_SECRET  string `validate:"required"`
}

var Envs env

func EnvValidation() map[string]string {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var validate = validator.New()

	Envs = env{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_PORT:     os.Getenv("DB_PORT"),
		PORT:        os.Getenv("PORT"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
	}
	err := validate.Struct(Envs)

	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("%s is required", err.Field())
		}
		return errors
	}

	return nil
}
