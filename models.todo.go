package main

import (
  "github.com/jinzhu/gorm"
)

 // todoModel describes a todoModel type
 type todoModel struct {
  gorm.Model
  Title     string `json:"title"`
  Completed int    `json:"completed"`
 }
