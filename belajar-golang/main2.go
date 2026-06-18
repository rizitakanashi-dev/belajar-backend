package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// PERBAIKAN: Huruf pertama field harus KAPITAL agar bisa diakses oleh Gin dan SQL Scan
type Siswa struct {
	NoAbsen int    `json:"no_absen"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
}

func main() {
	var err error

	db, err = sql.Open("mysql", "root:131109@tcp(127.0.0.1:3306)/Sekolah")
	if err != nil {
		log.Fatal("gagal connect ke database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("database tidak merespon")
	}
	fmt.Println("koneksi sukses")

	r := gin.Default()

	// --- ENDPOINT POST: MENYIMPAN DATA ---
	r.POST("/user", func(c *gin.Context) {
		var data Siswa

		// Perbaikan typo ke ShouldBindJSON
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON salah atau data tidak lengkap"})
			return
		}

		query := "INSERT INTO Murid (no_absen, nama, kelas) VALUES (?, ?, ?)"
		result, err := db.Exec(query, data.NoAbsen, data.Nama, data.Kelas)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
			return
		}
		datamasuk, _ := result.LastInsertId()

		c.JSON(http.StatusCreated, gin.H{
			"status": "Sukses",
			"pesan":  "Data siswa tersimpan",
			"id":     datamasuk,
		})
	})

	// --- ENDPOINT GET: MENGAMBIL SEMUA DATA (SELESAI) ---
	r.GET("/user", func(c *gin.Context) {
		rows, err := db.Query("SELECT no_absen, nama, kelas FROM Murid")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
			return
		}
		defer rows.Close()

		// Membuat slice kosong untuk menampung daftar siswa
		var listSiswa []Siswa

		// Looping lewat setiap baris hasil query
		for rows.Next() {
			var s Siswa
			// Scan data dari kolom database ke dalam variabel struct 's'
			err := rows.Scan(&s.NoAbsen, &s.Nama, &s.Kelas)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca baris data"})
				return
			}
			// Masukkan data siswa ke dalam slice listSiswa
			listSiswa = append(listSiswa, s)
		}

		// Antisipasi jika tabel ternyata kosong agar tidak mengembalikan 'null' di JSON
		if listSiswa == nil {
			listSiswa = []Siswa{}
		}

		// Kirim data dalam bentuk array JSON
		c.JSON(http.StatusOK, listSiswa)
	})

	r.Run(":8080")
}
