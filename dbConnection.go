package main

import (
  _ "github.com/bmizerany/pq"
  "database/sql"
  "fmt"
)

func connect() {
  dbInfo := "username=" + dbUser() +
    " dbname=" + dbName() +
    " password=" + dbPass() +
    " sslmode=" + dbSSLMode()
  /* fmt.Println(dbInfo) */
  db, err := sql.Open("postgres", dbInfo)
  errorHandler(err)
  defer db.Close()

  rows, err := db.Query("SELECT name FROM items")
  errorHandler(err)
  var name string
  for rows.Next() {
    err = rows.Scan(&name)
    errorHandler(err)
    fmt.Println(name)
  }

}
