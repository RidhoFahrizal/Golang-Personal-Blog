package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Ganti dengan kredensial PostgreSQL lokalmu
	connStr := "postgres://postgres:password_kamu@localhost:5432/nama_db_kamu"
	
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Printf("Gagal konek: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Data contoh untuk blog
	title := "Belajar Golang dan Postgres"
	content := "Ini adalah konten artikel blog pertama saya."

	// Query SQL
	query := "INSERT INTO posts (title, content) VALUES ($1, $2)"

	// Eksekusi query
	_, err = conn.Exec(context.Background(), query, title, content)
	if err != nil {
		fmt.Printf("Gagal insert data: %v\n", err)
		return
	}

	fmt.Println("Data blog berhasil disimpan!")
}