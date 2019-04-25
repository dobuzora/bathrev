package database

import (
	"github.com/dobuzora/bathrev/internal/app/auth/password"
	"github.com/dobuzora/bathrev/internal/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // enable the sqlite3 dialect
	"os"
	"path/filepath"
)

var mkdirAll = os.MkdirAll

// New creates a new wrapper for the gorm database framework.
func New(dialect, connection, defaultUser, defaultPass string, strength int, createDefaultUserIfNoExist bool) (*GormDatabase, error) {
	createDirectoryIfSqlite(dialect, connection)

	db, err := gorm.Open(dialect, connection)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(10)

	if dialect == "sqlite3" {
		db.DB().SetMaxOpenConns(1)
	}

	db.AutoMigrate(new(models.User))
	userCount := 0
	db.Find(new(models.User)).Count(&userCount)

	if createDefaultUserIfNoExist && userCount == 0 {
		db.Create(&models.User{Name: defaultUser, Pass: password.CreatePassword(defaultPass, strength), Admin: true})
	}

	return &GormDatabase{DB: db}, nil
}

func createDirectoryIfSqlite(dialect string, connection string) {
	if dialect == "sqlite3" {
		if _, err := os.Stat(filepath.Dir(connection)); os.IsNotExist(err) {
			if err := mkdirAll(filepath.Dir(connection), 0777); err != nil {
				panic(err)
			}
		}
	}
}

// GormDatabase is a wrapper for the gorm framework.
type GormDatabase struct {
	DB *gorm.DB
}

// Close closes the gorm database connection.
func (d *GormDatabase) Close() {
	d.DB.Close()
}
