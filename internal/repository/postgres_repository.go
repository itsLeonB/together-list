package repository

import (
	"context"
	"database/sql"

	"github.com/itsLeonB/together-list/internal/entity"
	"github.com/itsLeonB/together-list/internal/logging"
	"github.com/rotisserie/eris"
)

type PostgresRepository struct {
	db             *sql.DB
	notionDBsCache []entity.NotionDatabase
}

func NewPostgresRepository(url string) *PostgresRepository {
	db, err := sql.Open("pgx", url)
	if err != nil {
		logging.Fatalf("failed to open database connection: %v", err)
	}

	if pingErr := db.Ping(); err != nil {
		if err = db.Close(); err != nil {
			logging.Errorf("failed to close database connection: %v", err)
		}
		logging.Fatalf("failed to ping database: %v", pingErr)
	}

	return &PostgresRepository{db, nil}
}

func (pr *PostgresRepository) GetNotionDatabases(ctx context.Context) ([]entity.NotionDatabase, error) {
	if len(pr.notionDBsCache) > 0 {
		return pr.notionDBsCache, nil
	}

	rows, err := pr.db.QueryContext(ctx, "SELECT keyword, database_id, api_key FROM notion_databases")
	if err != nil {
		return nil, eris.Wrap(err, "failed to query database IDs")
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logging.Errorf("failed to close rows: %v", err)
		}
	}()

	databases := make([]entity.NotionDatabase, 0)
	for rows.Next() {
		var database entity.NotionDatabase
		if err := rows.Scan(&database.Keyword, &database.DatabaseID, &database.APIKey); err != nil {
			return nil, eris.Wrap(err, "failed to scan row")
		}
		databases = append(databases, database)
	}

	if err := rows.Err(); err != nil {
		return nil, eris.Wrap(err, "error occurred during row iteration")
	}

	if len(databases) > 0 {
		pr.notionDBsCache = databases
	}

	return databases, nil
}

func (pr *PostgresRepository) Close() error {
	return pr.db.Close()
}
