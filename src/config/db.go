package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	// host := "127.0.0.1"
	// port := "3306"
	// dbname := "go"
	// username := "root"
	// password := ""

	// dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"

	dsn := "host=aws-0-ap-southeast-1.pooler.supabase.com " +
		"user=postgres.aaoeugvpknmcamicqgwu " +
		"password=NIOYcrDi5V5iofN4 " +
		"dbname=golang " +
		"port=5432 " +
		"sslmode=require " + // karena Supabase pakai SSL
		"TimeZone=Asia/Jakarta"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


	if err != nil {
		panic("Tidak dapat terkoneksi ke database")
	}

	return db
}