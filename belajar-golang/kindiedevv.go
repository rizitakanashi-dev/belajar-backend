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

type Member struct {
	Id    int    `json:"id"`
	Name  string `json:"nama"`
	Roles string `json:"role"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:131109@tcp(127.0.0.1:3306)/Anggota")

	if err != nil {
		log.Fatal("gagal menyambungkan ke database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("database tidak merespon")
	}
	fmt.Println("berhasil connect")

	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var data Member
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "format json salah atau tidak ada"})
			return
		}

		query := "INSERT INTO Info (id, nama, role) VALUES(?,?,?)"
		result, err := db.Exec(query, data.Id, data.Name, data.Roles)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan data ke database"})
			return
		}
		datamasuk, _ := result.LastInsertId()
		c.JSON(http.StatusCreated, gin.H{
			"status": "sukses",
			"pesan":  "data berhasil dimasukkan",
			"id":     datamasuk,
		})
	})

	r.GET("/user", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, nama, role FROM Info")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data dari database"})
			return
		}
		defer rows.Close()

		var listMember []Member

		for rows.Next() {
			var m Member
			err := rows.Scan(&m.Id, &m.Name, &m.Roles)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membaca baris data"})
				return
			}
			listMember = append(listMember, m)
		}
		if listMember == nil {
			listMember = []Member{}
		}
		c.JSON(http.StatusOK, listMember)
	})
	r.Run(":8080")
}
