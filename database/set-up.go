package database

import(
  "fmt"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectiDatabase() {
	dsn := "host=localhost user=postgres password=123456789 dbname=pemiluApi-go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database berhasil koneksi broour")

	DB = db
	
}