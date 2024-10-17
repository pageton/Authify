package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pageton/authify/config"
)

func getSQLFromSchema(filePath string, identifierType string, name string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sqlBuilder strings.Builder
	inSection := false

	commentPrefix := fmt.Sprintf("-- %s: %s", identifierType, name)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, commentPrefix) {
			inSection = true
			continue
		}

		if inSection && strings.TrimSpace(line) == "" {
			break
		}

		if inSection {
			sqlBuilder.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return sqlBuilder.String(), nil
}

func main() {
	schemaPath := "./db/migrations/schema.sql"

	fmt.Println("Starting database setup...")

	usersTableSQL, err := getSQLFromSchema(schemaPath, "Table", "Users")
	if err != nil {
		log.Fatal("Error fetching Users table SQL:", err)
	}

	authTableSQL, err := getSQLFromSchema(schemaPath, "Table", "Auth")
	if err != nil {
		log.Fatal("Error fetching Auth table SQL:", err)
	}

	authIndexSQL, err := getSQLFromSchema(schemaPath, "Index", "Auth")
	if err != nil {
		log.Fatal("Error fetching Auth index SQL:", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	fmt.Println("Creating Users table...")
	if _, err := db.Exec(usersTableSQL); err != nil {
		log.Fatal("Error creating Users table:", err)
	}
	fmt.Println("Users table created successfully.")

	fmt.Println("Creating Auth table...")
	if _, err := db.Exec(authTableSQL); err != nil {
		log.Fatal("Error creating Auth table:", err)
	}
	fmt.Println("Auth table created successfully.")

	fmt.Println("Creating Auth index...")
	if _, err := db.Exec(authIndexSQL); err != nil {
		log.Fatal("Error creating Auth index:", err)
	}
	fmt.Println("Auth index created successfully.")

	fmt.Println("Users Table SQL:")
	fmt.Println(usersTableSQL)

	fmt.Println("Auth Table SQL:")
	fmt.Println(authTableSQL)

	fmt.Println("Auth Index SQL:")
	fmt.Println(authIndexSQL)

	fmt.Println("Database setup completed successfully.")
}
