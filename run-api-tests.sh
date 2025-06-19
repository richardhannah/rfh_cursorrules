#!/bin/bash

echo "🚀 Starting TOTM API Testing..."

# Stop any existing containers
echo "📦 Stopping existing containers..."
docker-compose -f docker-compose.apitest.yml down

# Start the infrastructure (database and API)
echo "🏗️  Starting infrastructure services..."
docker-compose -f docker-compose.apitest.yml up -d postgres

# Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready..."
sleep 10

# Start the API
echo "🔧 Starting TOTM API..."
docker-compose -f docker-compose.apitest.yml up -d totm-api

# Wait for API to be ready
echo "⏳ Waiting for API to be ready..."
sleep 15

# Run Newman tests
echo "🧪 Running API tests with Newman..."
docker run --rm \
    --network totmapi_totm-network \
    -v "$(pwd)/postman/TOTM_API_Collection.json:/etc/newman/TOTM_API_Collection.json:ro" \
    -v "$(pwd)/postman/environment.json:/etc/newman/environment.json:ro" \
    -v "$(pwd)/postman/reports:/etc/newman/reports" \
    postman/newman:latest \
    newman run /etc/newman/TOTM_API_Collection.json \
    --environment /etc/newman/environment.json \
    --reporters cli,json \
    --reporter-json-export /etc/newman/reports/newman-report.json \
    --timeout-request 10000 \
    --timeout-script 10000

# Check if tests passed
if [ $? -eq 0 ]; then
    echo "✅ API tests completed successfully!"
else
    echo "❌ API tests failed!"
fi

# Show test results
if [ -f "postman/reports/newman-report.json" ]; then
    echo "📊 Test report saved to: postman/reports/newman-report.json"
fi

echo "🧹 Cleaning up..."
docker-compose -f docker-compose.apitest.yml down

echo "�� Testing complete!" 