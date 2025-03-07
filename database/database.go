package database

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

var DbConnection *sql.DB

func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err, "Error while migrating")
	}
	DbConnection = dbParam

	fmt.Printf("Migrated %d migrations\n", n)
}
