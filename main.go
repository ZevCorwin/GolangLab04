package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Định nghĩa cấu trúc cho Student
type Student struct {
	ID    string `json:"id"`    // Mã sinh viên
	Name  string `json:"name"`  // Tên sinh viên
	Age   int    `json:"age"`   // Tuổi sinh viên
	Email string `json:"email"` // Email sinh viên
}

// Khai báo biến lưu trữ dữ liệu sinh viên dưới dạng map
var students = make(map[string]Student)

// Tạo mới sinh viên (Create)
func createStudent(c *gin.Context) {
	var newStudent Student

	// Parse dữ liệu JSON từ body của yêu cầu
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lưu sinh viên vào map với ID làm key
	students[newStudent.ID] = newStudent
	c.JSON(http.StatusCreated, newStudent)
}

// Lấy danh sách sinh viên (Read)
func getStudents(c *gin.Context) {
	// Convert map thành một slice để trả về dạng JSON
	studentList := make([]Student, 0, len(students))
	for _, student := range students {
		studentList = append(studentList, student)
	}

	c.JSON(http.StatusOK, studentList)
}

// Cập nhật thông tin sinh viên (Update)
func updateStudent(c *gin.Context) {
	id := c.Param("id") // Lấy ID từ URL

	var updatedStudent Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra sinh viên có tồn tại hay không
	if _, exists := students[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Cập nhật thông tin sinh viên
	updatedStudent.ID = id
	students[id] = updatedStudent
	c.JSON(http.StatusOK, updatedStudent)
}

// Xóa sinh viên (Delete)
func deleteStudent(c *gin.Context) {
	id := c.Param("id")

	// Kiểm tra sinh viên có tồn tại hay không
	if _, exists := students[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Xóa sinh viên khỏi map
	delete(students, id)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

func main() {
	r := gin.Default() // Tạo một router mặc định của Gin

	// Định nghĩa route cho "Hello, world"
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world")
	})

	// CRUD routes cho Student
	r.POST("/students", createStudent)       // Tạo mới sinh viên
	r.GET("/students", getStudents)          // Lấy danh sách sinh viên
	r.PUT("/students/:id", updateStudent)    // Cập nhật sinh viên
	r.DELETE("/students/:id", deleteStudent) // Xóa sinh viên

	// Chạy server ở cổng mặc định 8080
	r.Run()
}
