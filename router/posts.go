package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var DB *sql.DB

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ========= REGISTER ROUTES =========

func RegisterPostRoutes(r *mux.Router) {
	r.HandleFunc("/posts", GetPosts).Methods("GET")
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")
}

// ========= HANDLERS ==========

// GET /posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := DB.Query("SELECT id, title FROM posts")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		posts = append(posts, p)
	}

	json.NewEncoder(w).Encode(posts)
}

// POST /posts
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var data map[string]string
	json.Unmarshal(body, &data)

	title := data["title"]

	stmt, err := DB.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("New post was created"))
}

// GET /posts/{id}
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	row := DB.QueryRow("SELECT id, title FROM posts WHERE id = ?", id)

	var p Post
	err := row.Scan(&p.ID, &p.Title)
	if err == sql.ErrNoRows {
		http.Error(w, "Post not found", 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(p)
}

// PUT /posts/{id}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var data map[string]string
	json.Unmarshal(body, &data)

	newTitle := data["title"]

	stmt, err := DB.Prepare("UPDATE posts SET title = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newTitle, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "Post with ID = %s was updated", id)
}

// DELETE /posts/{id}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stmt, err := DB.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", id)
}
