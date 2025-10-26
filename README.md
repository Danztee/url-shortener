# URL Shortener

A fast, simple, and lightweight URL shortening service built with Go. This API allows you to create short URLs from long URLs and redirect users to the original URLs.

## Features

- Quick URL shortening with short codes
- Automatic redirect to original URLs
- MongoDB persistent storage
- RESTful API

## Tech Stack

- Go 1.25.3
- Gin web framework
- MongoDB

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Danztee/url-shortener.git
   cd url-shortener
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Create `.env` file:
   ```env
   MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
   BASE_URL=http://localhost:8080
   PORT=8080
   ```

## Running

```bash
go run main.go
```

## API Endpoints

**Shorten a URL:**

```bash
POST /shorten
Content-Type: application/json

{
  "original_url": "https://example.com/very/long/url"
}

Response:
{
  "short_url": "http://localhost:8080/abc123"
}
```

**Redirect to Original URL:**

```bash
GET /:code
```

## Environment Variables

| Variable    | Description               |
| ----------- | ------------------------- |
| `MONGO_URI` | MongoDB connection string |
| `BASE_URL`  | Base URL for short URLs   |
| `PORT`      | Server port number        |

## 📧 Support

If you have any questions or run into any issues, please open an issue on the GitHub repository.

---

Made with ❤️ by Danztee
