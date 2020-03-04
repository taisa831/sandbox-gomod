package main

import (
	"github.com/jinzhu/gorm"
)

func main() {
	gorm.Open("sqlite3", "test.db")
}
