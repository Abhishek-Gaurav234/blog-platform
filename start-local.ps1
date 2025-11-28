# Blog Platform - Quick Start Script for Windows
# This script helps you run the application locally

Write-Host "==================================" -ForegroundColor Cyan
Write-Host "Blog Platform - Quick Start" -ForegroundColor Cyan
Write-Host "==================================" -ForegroundColor Cyan
Write-Host ""

# Check if Node.js is installed
Write-Host "Checking prerequisites..." -ForegroundColor Yellow
$nodeInstalled = Get-Command node -ErrorAction SilentlyContinue
if (-not $nodeInstalled) {
    Write-Host "❌ Node.js is not installed!" -ForegroundColor Red
    Write-Host "Please install Node.js from https://nodejs.org/" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Alternatively, use Docker:" -ForegroundColor Green
    Write-Host "  docker-compose up --build" -ForegroundColor White
    exit 1
}

# Check if Go is installed
$goInstalled = Get-Command go -ErrorAction SilentlyContinue
if (-not $goInstalled) {
    Write-Host "❌ Go is not installed!" -ForegroundColor Red
    Write-Host "Please install Go from https://golang.org/" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Alternatively, use Docker:" -ForegroundColor Green
    Write-Host "  docker-compose up --build" -ForegroundColor White
    exit 1
}

Write-Host "✅ Node.js is installed" -ForegroundColor Green
Write-Host "✅ Go is installed" -ForegroundColor Green
Write-Host ""

# Install frontend dependencies if needed
if (-not (Test-Path "frontend\node_modules")) {
    Write-Host "Installing frontend dependencies..." -ForegroundColor Yellow
    Push-Location frontend
    npm install
    Pop-Location
    Write-Host "✅ Frontend dependencies installed" -ForegroundColor Green
    Write-Host ""
}

# Start backend in background
Write-Host "Starting backend server..." -ForegroundColor Yellow
$env:DB_PATH = "./blog.db"
Start-Process powershell -ArgumentList "-NoExit", "-Command", "Write-Host 'Backend Server Running on http://localhost:8080' -ForegroundColor Green; cd backend; go run cmd/api/main.go"
Start-Sleep -Seconds 3

# Start frontend
Write-Host "Starting frontend server..." -ForegroundColor Yellow
Push-Location frontend
Start-Process powershell -ArgumentList "-NoExit", "-Command", "Write-Host 'Frontend Running on http://localhost:3000' -ForegroundColor Green; npm run dev -- --config vite.config.local.js"
Pop-Location

Write-Host ""
Write-Host "==================================" -ForegroundColor Cyan
Write-Host "🚀 Application Started!" -ForegroundColor Green
Write-Host "==================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Frontend: http://localhost:3000" -ForegroundColor White
Write-Host "Backend:  http://localhost:8080" -ForegroundColor White
Write-Host ""
Write-Host "Press Ctrl+C in the server windows to stop" -ForegroundColor Yellow
Write-Host ""
