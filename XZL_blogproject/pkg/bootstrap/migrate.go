package bootstrap

import (
	"xzl_blog/internal/database/migration"
	"xzl_blog/pkg/config"
	"xzl_blog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
