# TOTM API Testing with Newman

This directory contains the Newman/Postman collection and configuration for automated API testing of the TOTM API.

## Files

- `TOTM_API_Collection.json` - Postman collection with all API endpoints and tests
- `environment.json` - Newman environment variables for testing
- `reports/` - Directory for test reports (created automatically)
- `README.md` - This file

## Running API Tests

### Prerequisites

1. Docker and Docker Compose installed
2. The TOTM API project built and ready

### Quick Start

1. **Set up the environment** (if needed):
   ```bash
   # Set your database connection string (optional - uses default for testing)
   export TOTM_CONN_STRING="postgres://user:pass@localhost:5432/dbname?sslmode=disable"
   ```

2. **Run the API tests**:
   ```bash
   docker-compose -f docker-compose.apitest.yml up --build
   ```

3. **View results**:
   - CLI output will show test results in the terminal
   - JSON report saved to `postman/reports/newman-report.json`

### What the tests cover

The collection includes tests for:

- **Health Check**: `/health` endpoint
- **Authentication**: 
  - User registration (`/register`)
  - User login (`/login`)
  - Password change (`/changepass`)
- **Blog Management**:
  - Get all blog posts (`/blogposts`)
  - Get specific blog post (`/blogposts/{id}`)
  - Create blog post (`POST /blogposts`)
  - Update blog post (`PUT /blogposts`)
- **AI Features**: OpenAI prompt (`/openai/prompt`)
- **Shop Management**: Shop info (`/shop`)
- **Error Handling**: Invalid requests, unauthorized access, etc.

### Test Flow

1. **Health Check**: Verifies API is running and database is connected
2. **User Registration**: Creates a test user
3. **User Login**: Authenticates and gets JWT token
4. **Protected Endpoints**: Tests endpoints requiring authentication
5. **Error Scenarios**: Tests invalid requests and error handling

### Customization

- **Environment Variables**: Edit `environment.json` to change test data
- **Collection**: Modify `TOTM_API_Collection.json` to add/remove tests
- **Database**: Update connection string in `docker-compose.apitest.yml`

### Reports

After running tests, check:
- `postman/reports/newman-report.json` - Detailed JSON report
- Terminal output - Real-time test results

### Troubleshooting

- **Port conflicts**: Change ports in `docker-compose.apitest.yml`
- **Database issues**: Check PostgreSQL container logs
- **API connection**: Verify the API is running on the expected port 