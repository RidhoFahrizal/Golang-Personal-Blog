package repository

import (
	"database/sql"
	"errors"
	_ "errors"

	"github.com/RidhoFahrizal/Golang-Personal-Blog/internal/model"
)

type postRepository struct{
	// this sturct connect the backend with the postgresql
	db *sql.DB
}

func (p *postRepository)GetByID(id string)(*model.Post, error){
	query := `
		SELECT id, title, body, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	var post model.Post


	// using .Scan we insert the post. 
	// at the same thime we checked the error 
	err := p.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Body,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil{
		return nil, err 
	}

	// here we returned the address 
	return &post, nil 
}

func (p *postRepository)Create(post *model.Post)error{

	
	query :=`
		INSERT INTO posts (title, body)
		VALUES ($1,$2)
		RETURNING id, created_at, updated_at
	`


	err := p.db.QueryRow(
		query, 
		post.Title,
		post.Body,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil{
		return err
	}

	return nil 

}

func(p *postRepository)Delete(id string)error{
	result, err := p.db.Exec(
		`DELETE FROM posts WHERE d = $1`, id,
	)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err 
}

func(p *postRepository)List()([]model.Post, error){
	query := `
		SELECT id, title, body, created_at, updated_at
		FROM posts 
		ORDER BY created_at DESC
	`

	rows, err := p.db.Query(query)

	if err != nil {
		return nil, err 
	}

	defer rows.Close()

	posts := make([]model.Post, 0 )

	for rows.Next(){
		var post model.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil{
			return nil, err 
		}
		posts = append(posts, post)
	}

	if err:= rows.Err(); err != nil{
		return nil, err 
	}

	return posts, nil 
}

func(p *postRepository)Update(post *model.Post)error{
	query := `
		UPDATE posts 
		SET title = $1, 
			body = $2
		WHERE id = $3 
		RETURNING updated_at
	`	

	err := p.db.QueryRow(
		query, 
		post.Title,
		post.Body,
		post.ID,
	).Scan(&post.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil 
	}

	return err
}

func NewPostRepository(db *sql.DB)PostRepository{
	return &postRepository{db: db}
}

