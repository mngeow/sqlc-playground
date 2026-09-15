# sqlc-playground

## Local SQLite database and migrations

Run the commands below from the project root so Goose can find `.env` and the
relative database and migration paths.

### Setup

Install the Goose CLI if needed (this setup was verified with v3.28.0):

```powershell
go install github.com/pressly/goose/v3/cmd/goose@v3.28.0
```

Ensure Go's binary directory (normally `$HOME\go\bin` on Windows) is on your `PATH`.

A local `.env` is configured. For a fresh checkout, copy `.env.example` to `.env`.
Goose loads `.env` automatically:

```dotenv
GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING="./playground.db?_pragma=foreign_keys(1)"
GOOSE_MIGRATION_DIR=./migrations
```

`_pragma=foreign_keys(1)` enables foreign-key enforcement on each SQLite
connection opened by Goose. Other clients must enable it on their own connections.
The database file, SQLite sidecar files, and local `.env` are git-ignored.

### Apply migrations

```powershell
goose up
goose status
```

The first `goose up` creates `playground.db` automatically and applies the
migrations in order: `authors`, then `books`. Goose records migration history in
`goose_db_version`; subsequent runs apply only pending migrations.

### Create a new database with Go

Use the standalone command to create an initialized, empty SQLite database:

```powershell
go run ./cmd/createdb -db ./my-database.db
```

The `-db` flag defaults to `./new.db`. The parent directory must exist, and the
command returns an error if the file already exists. Database files ending in
`.db` in the project root are git-ignored.

To create the `authors` and `books` tables in this new database, run Goose with
the same path explicitly. `-env=none` disables loading the default database
settings from `.env` for this command:

```powershell
goose -env=none -dir ./migrations sqlite3 "./my-database.db?_pragma=foreign_keys(1)" up
```

### Common commands

```powershell
goose version                        # Show the current schema version
goose down                           # Undo the latest migration
goose -s create add_book_isbn sql     # Create the next numbered SQL migration
```

The current rollback migrations drop their respective tables, including any data.
