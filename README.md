Codingan 1_initiate.sql

```sql
-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE Bab7Servo (
    code INTEGER PRIMARY KEY,
    StatusServo INTEGER
);

-- +migrate StatementEnd
```

Codingan database.go

```package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var DBMigrasi embed.FS
var DBKonesi *sql.DB

func DBMigrate(con *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: DBMigrasi,
		Root:       "sql_migrations",
	}

	n, Errs := migrate.Exec(con, "postgres", migrations, migrate.Up)
	if Errs != nil {
		panic(Errs)
	}

	DBKonesi = con

	fmt.Println("Migrasi Sukses", n, migrations)
}
```

