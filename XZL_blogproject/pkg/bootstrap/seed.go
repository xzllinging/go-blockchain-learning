package bootstrap

import (
	"xzl_blog/internal/database/seeder"
	"xzl_blog/pkg/config"
	"xzl_blog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
