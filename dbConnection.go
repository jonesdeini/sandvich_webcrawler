package main

import (
  _ "github.com/bmizerany/pq"
  "database/sql"
  "fmt"
)

type sqlObject struct {
  Item ItemType
}

type ItemType struct {
  Id int
  Float_value float32
  Defindex int
  Quality int
  Name string
  Type string
  CreatedAt string
  UpdatedAt string
  level int
}

func connect() {
  dbInfo := "username=" + dbUser() +
    " dbname=" + dbName() +
    " password=" + dbPass() +
    " sslmode=disable"
  fmt.Println(dbInfo)
  db, err := sql.Open("postgres", dbInfo)
  errorHandler(err)
  /* var msg sqlObject */
  var msg string
  err = db.QueryRow("SELECT name FROM items").Scan(&msg)
  errorHandler(err)
  fmt.Println(msg)
}
