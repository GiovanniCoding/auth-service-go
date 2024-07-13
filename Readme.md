## Models from database schema sqlc
1. Generate models from database schema
```bash
sqlc generate
```

## Migrations
1. Load variables from .env file
```bash
export $(grep -v '^#' .env | xargs)
```

2. Create migration from diff
```bash
atlas migrate diff --dir "$MIGRATIONS_PATH" --to "$SCHEMA_PATH" --dev-url "$DB_DEV_URL"
```

3. Apply migrations
```bash
atlas migrate apply --url "$DB_URL"
```
