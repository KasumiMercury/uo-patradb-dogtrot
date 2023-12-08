package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"

	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/migrate"
)

func main() {
	loadEnv()

	dsn := os.Getenv("MYSQL_MIGRATE_DSN")

	ctx := context.Background()

	dir, err := atlas.NewLocalDir("ent/migrate/migrations/ent")
	if err != nil {
		log.Fatalf("failed to create migrations directory: %v", err)
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithForeignKeys(false),
		schema.WithDialect(dialect.MySQL),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	err = migrate.NamedDiff(ctx, fmt.Sprintf("mysql://%s", dsn), os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed to generate migrations: %v", err)
	}
}

func loadEnv() {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("failed to load .env file: %v", err)
		}
	}
}
