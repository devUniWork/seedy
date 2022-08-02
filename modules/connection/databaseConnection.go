package databaseConnection

import (
    "fmt"
    "database/sql"
    "log"
    "os/user"
    _ "github.com/lib/pq"
   )

func userName() string {
	user, err := user.Current()
	if err != nil {
	  log.Fatalf(err.Error())
	}
	return user.Username
}

func ConnectDb() (DB *sql.DB) {

    connStr := fmt.Sprintf("postgresql://%s:@localhost/fatzebra_development?sslmode=disable", userName())
    // Connect to database
    db, err := sql.Open("postgres", connStr)

    if err != nil {
      log.Fatal(err)
    }

    return db
}