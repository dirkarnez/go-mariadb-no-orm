package main

import (
    "database/sql"
    "fmt"
    "log"
    "reflect"
    "strings"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID       int
    Username string
    Email    string
}

func (u User) GetID() int {
    return u.ID
}

func (u User) GetUsername() string {
    return u.Username
}

func (u User) GetEmail() string {
    return u.Email
}

func getUserFieldNames() []string {
    return []string{"id", "username", "email"}
}

func getUserFieldValues(u User) []interface{} {
    return []interface{}{u.GetID(), u.GetUsername(), u.GetEmail()}
}

func insertUser(db *sql.DB, user User) error {
    fields := getUserFieldNames()
    values := getUserFieldValues(user)

    query := fmt.Sprintf("INSERT INTO users (%s) VALUES (?%s)", 
        strings.Join(fields, ", "), strings.Repeat(", ?", len(fields)-1))

    _, err := db.Exec(query, values...)
    return err
}

func main() {
    // Connect to the database
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()

    // Create a new user
    user := User{
        Username: "johndoe",
        Email:    "johndoe@example.com",
    }

    // Insert the user into the database
    if err := insertUser(db, user); err != nil {
        log.Fatalf("Failed to insert user: %v", err)
    }

    fmt.Println("User inserted successfully!")
}
