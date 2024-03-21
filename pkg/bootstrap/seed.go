package bootstrap

import (
	"demoBlog/internal/database/seeder"
	"demoBlog/pkg/config"
	"demoBlog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}