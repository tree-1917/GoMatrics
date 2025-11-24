# GoMatrics

GoMatrics is a RESTful API built with Go (Golang) and MySQL, designed for managing "posts".  
It can optionally integrate with **GoAccess** to provide real-time web log analytics.

---

## Features

- Connects to a MySQL database.
- CRUD operations for "posts":
  - **GET** `/posts` — List all posts
  - **POST** `/posts` — Create a new post
  - **GET** `/posts/{id}` — Get a specific post
  - **PUT** `/posts/{id}` — Update a specific post
  - **DELETE** `/posts/{id}` — Delete a specific post
- Uses **Gorilla Mux** for routing.
- Ready for local development on port `8000`.
- Optional integration with **GoAccess** for real-time web log analysis.

---

## Prerequisites

- Go 1.25 or higher
- MySQL
- Go modules dependencies:

```bash
go mod tidy
````

* (Optional) **GoAccess**:

```bash
sudo apt install goaccess  # Ubuntu/Debian
```

---

## Getting Started

1. **Clone the repository**:

```bash
git clone git@github.com:tree-1917/GoMatrics.git
cd GoMatrics
```

2. **Update database connection** in `main.go`:

```go
db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/gomatric")
```

3. **Run the server**:

```bash
go run main.go
```

4. **Test the API**:

```bash
curl http://localhost:8000/posts
```

---

## Real-Time Web Log Analysis with GoAccess

### What is GoAccess?

GoAccess is an open-source **real-time web log analyzer** and interactive viewer that runs in a terminal on *nix systems or through your browser.
It provides fast and valuable HTTP statistics for system administrators who require a visual server report on the fly.

> ![goaccess](https://goaccess.io/)

### How to use with GoMatrics

1. Enable logging for your Go API (or use Nginx/Apache logs).
2. Run GoAccess pointing to the log file:

```bash
goaccess /path/to/access.log -o /path/to/report.html --real-time-html
```

3. Open the generated `report.html` in your browser for live analytics.

---

## Project Structure

```
GoMatrics/
├── main.go       # Main application entry
├── router/       # Route handlers (posts.go)
├── go.mod        # Go module file
├── go.sum        # Dependencies checksum
└── README.md     # Project documentation
```

---

## License

This project is licensed under the MIT License.

---
