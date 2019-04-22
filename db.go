package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB

func init() {
  // open the database connection
  var err error
  db, err = gorm.Open("mysql", "root:@/go-todo?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
    panic("failed to connect to database")
  }

  // Migrate the schema
  // db.autoMigrate(&todoModel{})
}
