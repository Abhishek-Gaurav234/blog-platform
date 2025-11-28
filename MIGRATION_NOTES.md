# SQLite Migration Notes

## Changes Made

The blog platform has been successfully migrated from PostgreSQL to SQLite.

### Files Modified

1. **internal/models/database.go**
   - Changed driver from `github.com/lib/pq` to `github.com/mattn/go-sqlite3`
   - Updated connection string from PostgreSQL to SQLite format
   - Added automatic table creation on startup
   - Adjusted connection pool settings (SQLite works best with 1 connection)
   - Database file path configurable via `DB_PATH` environment variable (default: `./blog.db`)

2. **internal/models/post.go**
   - Changed SQL parameter placeholders from `$1, $2, ...` (PostgreSQL) to `?` (SQLite)
   - Updated `Create()` method to use `Exec()` with `LastInsertId()` instead of `RETURNING` clause
   - Updated `Update()` method to use `CURRENT_TIMESTAMP` instead of `NOW()`
   - Removed `fmt` package import (no longer needed)

3. **docker-compose.yml**
   - Removed PostgreSQL service
   - Simplified to single app service
   - Added volume for SQLite database persistence
   - Set `DB_PATH` environment variable to `/app/data/blog.db`

4. **Dockerfile**
   - Enabled CGO compilation (required for SQLite)
   - Added build dependencies: `gcc`, `musl-dev`, `sqlite-dev`
   - Added runtime dependency: `sqlite-libs`
   - Created data directory for database file

5. **go.mod / go.sum**
   - Replaced `github.com/lib/pq` with `github.com/mattn/go-sqlite3`

## Running the Application

### Local Development (Windows)

**Important**: SQLite requires CGO, which needs a C compiler (GCC) on Windows.

**Option 1: Install GCC (Recommended)**
1. Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or [MinGW-w64](https://www.mingw-w64.org/)
2. Add GCC to your PATH
3. Run the application:
```powershell
# Set database path (optional, defaults to ./blog.db)
$env:DB_PATH="./blog.db"

# Enable CGO and run
$env:CGO_ENABLED="1"
go run cmd/api/main.go
```

**Option 2: Use Docker (Easiest)**
```powershell
# Build and run with docker-compose (no GCC required on host)
docker-compose up --build
```

### Local Development (Linux/Mac)

```bash
# Set database path (optional, defaults to ./blog.db)
export DB_PATH="./blog.db"

# CGO is enabled by default on Linux/Mac
go run cmd/api/main.go
```

### Docker

```bash
# Build and run with docker-compose
docker-compose up --build

# The SQLite database will be persisted in the sqlite_data volume
```

## Database Location

- **Local Development**: `./blog.db` (or path specified in `DB_PATH` env var)
- **Docker**: `/app/data/blog.db` (persisted in `sqlite_data` volume)

## Key Differences from PostgreSQL

1. **Single Connection**: SQLite uses a single connection for best performance
2. **File-Based**: Database is stored in a single file, no separate server needed
3. **Simpler Deployment**: No need to manage a separate database container
4. **Automatic Schema Creation**: Tables are created automatically on first run

## Migration Benefits

- Simpler deployment architecture
- No separate database server required
- Reduced resource consumption
- Easier local development setup
- Database is portable (single file)
