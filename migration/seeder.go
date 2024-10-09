package migration

import (
	"github.com/reynaldineo/go-gin-gorm-starter/migration/seed"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seed.ListUserSeeder(db); err != nil {
		return err
	}

	return nil
}
