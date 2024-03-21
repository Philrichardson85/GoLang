package bootstrap

import (
	"demoBlog/internal/database/migration"
	"demoBlog/pkg/config"
	"demoBlog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}