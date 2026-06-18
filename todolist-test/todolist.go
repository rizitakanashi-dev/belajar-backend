package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Todo merepresentasikan struktur data untuk Database dan JSON
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"` // Wajib diisi saat POST/PUT
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Variabel global untuk koneksi database
var db *sql.DB

func main() {
	var err error

	// KONEKSI DATABASE
	// Format: username:password@tcp(host:port)/nama_database
	// Silakan sesuaikan dengan kredensial MariaDB Anda
	dsn := "root:131109@tcp(127.0.0.1:3306)/todolist_db"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database: ", err)
	}
	defer db.Close()

	// Memastikan database benar-benar merespon
	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon: ", err)
	}
	fmt.Println("Berhasil terhubung ke MariaDB!")

	// INISIALISASI GIN ROUTER
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// REGISTRASI ROUTING (CRUD)
	r.POST("/todos", createTodo)       // Create
	r.GET("/todos", getAllTodos)       // Read All
	r.GET("/todos/:id", getTodoByID)   // Read One
	r.PUT("/todos/:id", updateTodo)    // Update
	r.DELETE("/todos/:id", deleteTodo) // Delete

	// JALANKAN SERVER
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}

// ==========================================
// HANDLER FUNCTIONS (LOGIKA API)
// ==========================================

// 1. CREATE: Menambah tugas baru
func createTodo(c *gin.Context) {
	var newTodo Todo

	// Validasi input JSON. Jika 'title' kosong, akan otomatis error 400
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kolom 'title' wajib diisi"})
		return
	}

	query := "INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)"
	result, err := db.Exec(query, newTodo.Title, newTodo.Description, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
		return
	}

	// Mengambil ID yang baru saja digenerate oleh database
	id, _ := result.LastInsertId()
	newTodo.ID = int(id)
	newTodo.Completed = false

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo berhasil dibuat",
		"data":    newTodo,
	})
}

// 2. READ ALL: Mengambil semua daftar tugas
func getAllTodos(c *gin.Context) {
	query := "SELECT id, title, description, completed FROM todos"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	// Slice untuk menampung semua todo
	todos := []Todo{}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data"})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// 3. READ ONE: Mengambil satu tugas berdasarkan ID
func getTodoByID(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, title, description, completed FROM todos WHERE id = ?"

	var todo Todo
	err := db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// 4. UPDATE: Mengubah data tugas berdasarkan ID
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo

	// 1. Validasi input JSON dari klien
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid atau 'title' kosong"})
		return
	}

	// 2. Cek apakah data yang mau di-update memang ada di database
	var existingTodo Todo
	checkQuery := "SELECT id FROM todos WHERE id = ?"
	err := db.QueryRow(checkQuery, id).Scan(&existingTodo.ID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan"})
		return
	}

	// 3. Eksekusi update jika data ditemukan
	query := "UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?"
	_, err = db.Exec(query, updatedTodo.Title, updatedTodo.Description, updatedTodo.Completed, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo berhasil diperbarui"})
}

// 5. DELETE: Menghapus tugas berdasarkan ID
func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	// Cek apakah data ada sebelum dihapus
	var existingTodo Todo
	checkQuery := "SELECT id FROM todos WHERE id = ?"
	err := db.QueryRow(checkQuery, id).Scan(&existingTodo.ID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan"})
		return
	}

	query := "DELETE FROM todos WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo berhasil dihapus"})
}
