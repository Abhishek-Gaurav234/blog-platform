# Blog Platform - SQLite Setup Guide

## ✅ Migration Complete!

Your blog platform has been successfully migrated from PostgreSQL to SQLite.

## Quick Start (Windows - No GCC Required)

The easiest way to run the application on Windows is using Docker:

```powershell
# Build and start the application
docker-compose up --build

# The API will be available at http://localhost:8080
```

## Testing the API

Once the server is running, you can test the endpoints:

### Create a Post
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/posts" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"title":"My First Post","content":"Hello SQLite!","type":"article","author_id":1,"status":"published"}'
```

### Get All Posts
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/posts" -Method GET
```

### Get a Specific Post
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/posts/1" -Method GET
```

## What Changed?

1. **Database**: PostgreSQL → SQLite (single file database)
2. **No separate DB container**: Everything runs in one container
3. **Data persistence**: Database stored in Docker volume `sqlite_data`
4. **Simpler deployment**: No database server to configure

## File Locations

- **Docker**: Database at `/app/data/blog.db` (persisted in volume)
- **Local** (if you install GCC): Database at `./blog.db` in project root

## Next Steps

If you want to run locally without Docker:
1. Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) (provides GCC for Windows)
2. Add GCC to your system PATH
3. Run: 
   ```powershell
   $env:CGO_ENABLED="1"
   go run cmd/api/main.go
   ```

See `MIGRATION_NOTES.md` for detailed technical information.
