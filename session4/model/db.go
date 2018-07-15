package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// buat variabel global untuk koneksi ke database
var DB *gorm.DB

func init(){
	var err error
	// assign value ke variabel DB, jangan gunakan := karena value tidak akan terassign secara global
	DB, err = gorm.Open("mysql", "root@/training?charset=utf8&parseTime=True&loc=Local")
	// contoh koneksi dengan password dan custom ip + port
	//DB, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// untuk membuat pembacaan struct menjadi singular, contoh User akan mencari tabel user, jika tidak akan mencari tabel users
	DB.SingularTable(true)
}