# URL Shortener

A fast, simple, and lightweight URL shortening service built with Go. This API allows you to create short URLs from long URLs and redirect users to the original URLs.

## 🚀 Features

- **Quick URL Shortening**: Generate short codes for long URLs
- **Automatic Redirection**: Seamless redirect to original URLs
- **MongoDB Backend**: Persistent storage with MongoDB
- **RESTful API**: Clean and simple API endpoints
- **Go Performance**: Built with Go for optimal performance
- **Environment-based Configuration**: Easy configuration with environment variables

## 🛠️ Tech Stack

- **Language**: Go 1.25.3
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: MongoDB
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
- **Logging**: [Logrus](https://github.com/sirupsen/logrus)

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.25.3 or higher)
- [MongoDB](https://www.mongodb.com/try/download/community) (local installation or MongoDB Atlas account)

## 🔧 Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Danztee/url-shortener.git
   cd url-shortener
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Set up environment variables:**

   Create a `.env` file in the root directory with the following variables:

   ```env
   MONGO_URI=your_mongodb_connection_string
   BASE_URL=http://localhost:8080
   PORT=8080
   ```

   Example for local MongoDB:

   ```env
   MONGO_URI=mongodb://localhost:27017
   BASE_URL=http://localhost:8080
   PORT=8080
   ```

   For MongoDB Atlas:

   ```env
   MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
   BASE_URL=http://localhost:8080
   PORT=8080
   ```

## 🚦 Running the Application

1. **Start your MongoDB instance** (if running locally):

   ```bash
   mongod
   ```

   Or use MongoDB Atlas if you prefer cloud hosting.

2. **Run the application:**

   ```bash
   go run main.go
   ```

   The server will start on the port specified in your `.env` file (default: 8080).

## 📡 API Documentation

### Base URL

```
http://localhost:8080
```

### Endpoints

#### 1. Shorten a URL

Creates a short code for a long URL.

**Endpoint:** `POST /shorten`

**Request Body:**

```json
{
  "original_url": "https://example.com/very/long/url/path"
}
```

**Response:**

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

**Example using cURL:**

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://example.com/very/long/url"}'
```

#### 2. Redirect to Original URL

Redirects to the original URL using the short code.

**Endpoint:** `GET /:code`

**Example:**

```
http://localhost:8080/abc123
```

**Response:**

- **Status 302**: Redirects to the original URL
- **Status 404**: Short code not found

## 📁 Project Structure

```
url-shortener/
├── config/
│   └── db.go           # Database connection configuration
├── controllers/
│   └── url.go          # URL shortening and redirection logic
├── models/
│   └── url.go          # Data models for URLs
├── routes/
│   └── routes.go       # API route definitions
├── utils/
│   └── shortener.go    # Short code generation utility
├── middleware/         # Middleware (if any)
├── main.go            # Application entry point
├── go.mod             # Go module dependencies
└── README.md          # This file
```

## 🗄️ Database Schema

The application uses MongoDB with the following collection structure:

**Collection:** `urls`  
**Database:** `url_shortener`

```json
{
  "_id": "ObjectId",
  "short_code": "abc123",
  "original": "https://example.com/very/long/url",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## 🎯 How It Works

1. **URL Shortening:**

   - Client sends POST request to `/shorten` with the original URL
   - Server generates a random 6-character short code
   - Short code and original URL are stored in MongoDB
   - Server returns the short URL

2. **URL Redirection:**
   - Client makes GET request to `/:code`
   - Server looks up the short code in MongoDB
   - Server redirects to the original URL

## 🔒 Environment Variables

| Variable    | Description               | Required |
| ----------- | ------------------------- | -------- |
| `MONGO_URI` | MongoDB connection string | Yes      |
| `BASE_URL`  | Base URL for short URLs   | Yes      |
| `PORT`      | Server port number        | Yes      |

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - The web framework used
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - MongoDB driver
- [Logrus](https://github.com/sirupsen/logrus) - Structured logger

## 📧 Support

If you have any questions or run into any issues, please open an issue on the GitHub repository.

---

Made with ❤️ by Danztee
