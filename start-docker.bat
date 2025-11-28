@echo off
echo ==================================
echo Blog Platform - Quick Start
echo ==================================
echo.

REM Check for Docker
docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo Docker is not installed or not in PATH
    echo Please install Docker Desktop from https://www.docker.com/products/docker-desktop
    echo.
    pause
    exit /b 1
)

echo Starting application with Docker...
echo.
echo This will start both backend and frontend
echo.
echo Backend:  http://localhost:8080
echo Frontend: http://localhost:3000
echo.

docker-compose up --build

pause
