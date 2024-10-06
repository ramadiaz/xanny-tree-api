package repositories

import (
	"database/sql"
	"log"
	"xanny-tree-api/config"
	"xanny-tree-api/dto"
)

type CompRepository interface{
	RegisterUrl(data dto.Tree) error
	GetUrl(url string) (*string, error) 

	AddLike() (*int64, error) 
}

type compRepository struct {
	DB *sql.DB
}

func NewComponentRepository(DB *sql.DB) *compRepository {
	db := config.InitDB()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS client_track (
			id BIGSERIAL PRIMARY KEY NOT NULL,
			ip VARCHAR(255),
			browser VARCHAR(255),
			version VARCHAR(255),
			os VARCHAR(255),
			device VARCHAR(255),
			origin VARCHAR(255),
			api VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tree_url (
			id BIGSERIAL PRIMARY KEY NOT NULL,
			name VARCHAR(255) NOT NULL,
			short_url VARCHAR(255) NOT NULL,
			original_url VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS like_count (
			id BIGSERIAL PRIMARY KEY NOT NULL,
			name VARCHAR(255) UNIQUE NOT NULL,
			total BIGINT NOT NULL DEFAULT 0,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		INSERT INTO like_count (name)
		VALUES ('profile')
		ON CONFLICT (name) DO NOTHING;
		`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return &compRepository{
		DB: db,
	}
}
