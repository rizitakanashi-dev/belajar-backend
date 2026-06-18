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

type Siswa struct {
	No_absen int    `json:"no_absen"`
	Nama     string `json:"nama"`
	Kelas    string `json:"kelas"`
}

func main() {
	var err error // tipe data error

	db, err = sql.Open("mysql", "root:131109@tcp(127.0.0.1:3306)/Sekolah") // jika tidak ada error, maka akan konek ke database
	// nil itu adalah null / kosong
	if err != nil {
		log.Fatal("gagal connect ke database") //jika error konek ke database
	}
	defer db.Close() //akan langsung tertutup

	// untuk mengecek ping database
	err = db.Ping()
	if err != nil {
		log.Fatal("database tidak merespon")
	}
	fmt.Println("koneksi sukses")
	// fmt adalah module untuk input output

	r := gin.Default()

	r.POST("/user", func(c gin.Context) {
		var data Siswa

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Gagal menyimpan ke database",
			})
			return
		}
		query := "INSERT INTO Murid (no_absen, nama, kelas) VALUES (?, ?, ?)"

		// memasukkan data query nya ke dalam database
		result, err := db.Exec(query, data.No_absen, data.Nama, data.Kelas)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Gagal menyimpan data ke database",
			})
			return
		}
		datamasuk, _ := result.LastInsertId()

		c.JSON(http.StatusCreated, gin.H{
			"status": "Sukses",
			"pesan":  "Data siswa tersimpan",
			"absen":  datamasuk,
		})
	})

	r.GET("/murid", func(c *gin.Context) {
		rows, err := db.Query("SELECT no_absen, nama, kelas FROM Murid")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
			return
		}
		defer rows.Close()

		var siswa []Siswa.

		for rows.Next() {
			var murid Siswa
			if err := rows.Scan(&murid.No_absen, &murid.Nama, &murid.Kelas); err != nil {
				continue
			}
			murid = append(siswa, murid)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "sukses",
			"data":   siswa,
		})
	})
	r.Run(":8080")
}
