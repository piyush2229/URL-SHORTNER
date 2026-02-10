# URL Shortener

A full-stack URL shortening service built with Go backend and React frontend. This application allows users to convert long URLs into short, shareable links with persistence using MongoDB.

## Features

✨ **Core Features:**
- **URL Shortening**: Convert long URLs into short, unique codes
- **URL Retrieval**: Redirect from short URLs to original long URLs
- **MongoDB Persistence**: All shortened URLs are stored in MongoDB
- **RESTful API**: Clean and simple API endpoints for URL operations
- **CORS Enabled**: Frontend can communicate with backend seamlessly
- **Responsive UI**: Modern React-based user interface

## Technology Stack

### Backend
- **Language**: Go (46.6%)
- **Database**: MongoDB
- **Framework**: Net/HTTP (Go standard library)
- **Dependencies**:
  - go.mongodb.org/mongo-driver - MongoDB driver for Go
  - github.com/joho/godotenv - Environment variable management

### Frontend
- **Framework**: React 19.2.0 (38.2%)
- **Build Tool**: Vite 7.3.1
- **Styling**: CSS (11.3%)
- **HTML**: HTML5 (3.9%)
- **Other**: React DOM for DOM manipulation

## Project Structure

```
URL-SHORTNER/
├── main.go                 # Backend server and API handlers
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── frontend/               # React frontend application
│   ├── src/               # React source files
│   ├── index.html         # HTML entry point
│   ├── vite.config.js     # Vite configuration
│   ├── package.json       # NPM dependencies
│   └── eslint.config.js   # ESLint configuration
└── .gitignore             # Git ignore rules
```

## Setup Instructions

### Prerequisites
- Go 1.16+ installed
- Node.js and npm installed
- MongoDB instance running

### Backend Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/piyush2229/URL-SHORTNER.git
   cd URL-SHORTNER
   ```

2. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   MONGO_URI=mongodb://localhost:27017
   MONGO_DB=url_shortener
   MONGO_COLLECTION=urls
   BASE_URL=http://localhost:3000
   ```

3. **Install Go dependencies**
   ```bash
   go mod download
   ```

4. **Run the backend server**
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Run development server**
   ```bash
   npm run dev
   ```
   The frontend will be available at `http://localhost:5173`

4. **Build for production**
   ```bash
   npm run build
   ```

## API Endpoints

### POST /shorten
**Shortens a URL**

Request:
```json
{
  "url": "https://example.com/very/long/url"
}
```

Response:
```json
{
  "id": 1,
  "original_url": "https://example.com/very/long/url",
  "short_url": "abc12345",
  "creation_date": "2026-02-10"
}
```

### GET /:shortUrl
**Redirects to the original URL**

## Frontend URL

🌐 **Frontend Application**: `http://localhost:5173`

The frontend provides a user-friendly interface to:
- Enter long URLs to be shortened
- Copy shortened URLs to clipboard
- View URL history and details

## Configuration

The application supports environment variables for customization:

| Variable | Default | Description |
|----------|---------|-------------|
| `MONGO_URI` | `mongodb://localhost:27017` | MongoDB connection URI |
| `MONGO_DB` | `url_shortener` | MongoDB database name |
| `MONGO_COLLECTION` | `urls` | MongoDB collection name |
| `BASE_URL` | `http://localhost:3000` | Base URL for shortened links |

## Development

### Backend Development
- The backend uses Go's `net/http` package for HTTP server
- MD5 hashing is used to generate short URL codes
- CORS headers are enabled for cross-origin requests

### Frontend Development
- React with Vite for fast development experience
- ESLint configured for code quality
- React Hooks and modern JavaScript features

### Running Tests

Frontend linting:
```bash
cd frontend
npm run lint
```

## Deployment

### Backend Deployment
```bash
go build -o url-shortener main.go
./url-shortener
```

### Frontend Deployment
```bash
cd frontend
npm run build
# Deploy the dist/ folder to your hosting service
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.

## Author

**Piyush Sharma** - [@piyush2229](https://github.com/piyush2229)

## Support

For issues and questions, please open an issue on GitHub: [Issues](https://github.com/piyush2229/URL-SHORTNER/issues)