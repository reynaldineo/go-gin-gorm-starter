package cmd

import (
	"log"
	"os"

	"github.com/reynaldineo/go-gin-gorm-starter/migration"
	"gorm.io/gorm"
)

func Commands(db *gorm.DB) {
	migrate := false
	seed := false
	migrateFresh := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
		if arg == "--migrate-fresh" {
			migrateFresh = true
			migrate = false
		}
	}

	if migrateFresh {
		var tables []string
		if err := db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables).Error; err != nil {
			log.Fatalf("error fetching tables: %v", err)
		}

		for _, table := range tables {
			if err := db.Migrator().DropTable(table); err != nil {
				log.Fatalf("error dropping table %s: %v", table, err)
			}
			log.Printf("table %s dropped successfully", table)
		}

		if err := migration.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("fresh migration completed successfully")
	}

	if migrate {
		if err := migration.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if seed {
		if err := migration.Seeder(db); err != nil {
			log.Fatalf("error seeding: %v", err)
		}
		log.Println("seeder completed successfully")
	}

}
