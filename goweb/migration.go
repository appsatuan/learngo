package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Jurnal struct {
	IDJurnal      int `gorm:"primaryKey;autoIncrement"`
	Kategori      string
	NPSN          string
	NIK           string
	WaktuMulai    string
	WaktuSelesai  string
	Durasi        int
	JamKeMulai    int
	JamKeSelesai  int
	JenisKegiatan string
	Lokasi        string
	Uraian        string
	Keterangan    string
}

type Absensi struct {
	IDAbsensi  int `gorm:"primaryKey;autoIncrement"`
	IDJurnal   int `gorm:"index"` // Foreign key relationship to Jurnal
	NPSN       string
	NIKSiswa   string
	NamaSiswa  string
	Keterangan string
}

func migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Jurnal{}, &Absensi{})
	if err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
	}

	log.Println("Migration completed successfully.")
}

func main() {
	migrate()
}
