#!/usr/bin/env pwsh

Write-Host "Starting TOTM API Testing..." -ForegroundColor Green

# Stop any existing containers
Write-Host "Stopping existing containers..." -ForegroundColor Yellow
docker-compose -f docker-compose.apitest.yml down --remove-orphans 2>$null

# Start the infrastructure (database and API)
Write-Host "Starting infrastructure services..." -ForegroundColor Yellow
docker-compose -f docker-compose.apitest.yml up -d postgres

# Wait for PostgreSQL to be ready
Write-Host "Waiting for PostgreSQL to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 10

# Start the API
Write-Host "Starting TOTM API..." -ForegroundColor Yellow
docker-compose -f docker-compose.apitest.yml up -d --build totm-api

# Wait for API to be ready
Write-Host "Waiting for API to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 15

# Run Newman tests
Write-Host "Running API tests with Newman..." -ForegroundColor Green
docker run --rm `
    --network totmapi_totm-network `
    -v "${PWD}/postman/TOTM_API_Collection.json:/etc/newman/TOTM_API_Collection.json:ro" `
    -v "${PWD}/postman/environment.json:/etc/newman/environment.json:ro" `
    -v "${PWD}/postman/reports:/etc/newman/reports" `
    postman/newman:latest run /etc/newman/TOTM_API_Collection.json `
    --environment /etc/newman/environment.json `
    --reporters cli,json `
    --reporter-json-export /etc/newman/reports/newman-report.json `
    --timeout-request 10000 `
    --timeout-script 10000

# Check if tests passed
if ($LASTEXITCODE -eq 0) {
    Write-Host "API tests completed successfully!" -ForegroundColor Green
} else {
    Write-Host "API tests failed!" -ForegroundColor Red
}

# Show test results
if (Test-Path "postman/reports/newman-report.json") {
    Write-Host "Test report saved to: postman/reports/newman-report.json" -ForegroundColor Cyan
}

Write-Host "Cleaning up..." -ForegroundColor Yellow
docker-compose -f docker-compose.apitest.yml down --remove-orphans

Write-Host "Testing complete!" -ForegroundColor Green 