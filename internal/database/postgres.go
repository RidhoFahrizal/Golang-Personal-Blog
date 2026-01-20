package database 

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq" 
)


func MustConnect() * sql.DB{
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("✅ PostgreSQL connected")
	return db
}

