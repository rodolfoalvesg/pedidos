package migration

import "gorm.io/gorm"

func EnableUUIDPostgres(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
}
