package main

import (
  _ "github.com/bmizerany/pq"
  "database/sql"
  "fmt"
)

  type dbObject struct {
    Item DbItemsType
  }

  type DbItemsType struct {
    Defindex   sql.NullInt64
    FloatValue sql.NullFloat64 `sql:"float_value"`
    Id         int
    Level      sql.NullInt64
    Name       string
    Quality    sql.NullInt64
  }

func connect() {
  dbInfo := "username=" + dbUser() +
    " dbname=" + dbName() +
    " password=" + dbPass() +
    " sslmode=" + dbSSLMode()
  /* fmt.Println(dbInfo) */
  db, err := sql.Open("postgres", dbInfo)
  errorHandler(err)
  defer db.Close()

  rows, err := db.Query("SELECT defindex, float_value, id, level, name, quality FROM items")
  errorHandler(err)
  var dbType DbItemsType
  for rows.Next() {
    err = rows.Scan(&dbType.Defindex, &dbType.FloatValue, &dbType.Id, &dbType.Level, &dbType.Name, &dbType.Quality)
    errorHandler(err)
    fmt.Println(dbType)
  }

}
