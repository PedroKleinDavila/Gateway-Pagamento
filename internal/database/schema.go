package database

import (
	"log"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"gorm.io/gorm"
)

func MigrateSchema(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.Account{},
		&domain.Invoice{},
	)

	if err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}
}
