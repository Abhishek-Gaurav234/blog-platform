package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Type      string    `json:"type" db:"type"` // article, tutorial, review
	AuthorID  int64     `json:"author_id" db:"author_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Status    string    `json:"status" db:"status"` // draft, published, archived
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *Post) error {
	query := `INSERT INTO posts (title, content, type, author_id, status) 
	          VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, post.Title, post.Content, post.Type, post.AuthorID, post.Status)
	if err != nil {
		return err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.ID = id

	// Fetch the created_at and updated_at timestamps
	selectQuery := `SELECT created_at, updated_at FROM posts WHERE id = ?`
	err = r.db.QueryRow(selectQuery, post.ID).Scan(&post.CreatedAt, &post.UpdatedAt)

	return err
}

func (r *PostRepository) FindByID(id int64) (*Post, error) {
	query := `SELECT id, title, content, type, author_id, created_at, updated_at, status 
	          FROM posts WHERE id = ?`

	post := &Post{}
	err := r.db.QueryRow(query, id).Scan(
		&post.ID, &post.Title, &post.Content, &post.Type,
		&post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &post.Status,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return post, err
}

func (r *PostRepository) FindAll(status, contentType string, limit, offset int) ([]*Post, error) {
	query := `SELECT id, title, content, type, author_id, created_at, updated_at, status 
	          FROM posts WHERE 1=1`
	var queryParams []interface{}

	if status != "" {
		query += " AND status = ?"
		queryParams = append(queryParams, status)
	}

	if contentType != "" {
		query += " AND type = ?"
		queryParams = append(queryParams, contentType)
	}

	query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	queryParams = append(queryParams, limit, offset)

	rows, err := r.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		post := &Post{}
		err := rows.Scan(
			&post.ID, &post.Title, &post.Content, &post.Type,
			&post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &post.Status,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) Update(post *Post) error {
	query := `UPDATE posts SET title = ?, content = ?, type = ?, status = ?, updated_at = CURRENT_TIMESTAMP
	          WHERE id = ?`

	_, err := r.db.Exec(query, post.Title, post.Content, post.Type, post.Status, post.ID)
	if err != nil {
		return err
	}

	// Fetch the updated timestamp
	selectQuery := `SELECT updated_at FROM posts WHERE id = ?`
	err = r.db.QueryRow(selectQuery, post.ID).Scan(&post.UpdatedAt)

	return err
}

func (r *PostRepository) Delete(id int64) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
