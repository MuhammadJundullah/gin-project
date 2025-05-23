package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Coba load .env, tapi jangan fatal error jika gagal (misal di Railway)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Gagal memuat file .env, pastikan variabel environment sudah di-set")
	}

	// Ambil variabel environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// Cek variabel wajib
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("❌ Variabel environment database belum lengkap")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal koneksi ke database:", err)
	}

	fmt.Println("✅ Berhasil konek ke database!")
	return db
}
