package main

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Student struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

var Students = []Student{
	Student{ID: shared.GetUuid(), Name: "Bruno", Age: 18},
	Student{ID: shared.GetUuid(), Name: "Gabriela", Age: 18},
}

func main() {
	service := gin.Default()
	getRoutes(service)
	service.Run()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHeart)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.POST("/", routePostStudents)
	groupStudents.PUT("/:id", routePutStudents)
	groupStudents.DELETE("/:id", routeDeleteStudents)
	groupStudents.GET("/:id", routeGetStudent)
	return c
}

func routeHeart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	c.Done()
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func routePostStudents(c *gin.Context) {
	var student Student

	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possível obter o payload",
		})
		return
	}

	student.ID = shared.GetUuid()
	Students = append(Students, student)

	c.JSON(http.StatusCreated, student)
}

func routePutStudents(c *gin.Context) {
	var studentPayload Student
	var studentLocal Student
	var newStudents []Student

	err := c.BindJSON(&studentPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possível obter o ID",
		})
		return
	}

	idString := c.Params.ByName("id")
	id, err := shared.GetUuidByString(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Problema com seu o ID",
		})
		return
	}

	for _, studentElement := range Students {
		if studentElement.ID == id {
			studentLocal = studentElement
		}
	}

	if studentLocal.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi encontrar o estudante",
		})
		return
	}

	studentLocal.Name = studentPayload.Name
	studentLocal.Age = studentPayload.Age

	for _, studentElement := range Students {
		if id == studentElement.ID {
			newStudents = append(newStudents, studentLocal)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	Students = newStudents

	c.JSON(http.StatusOK, studentLocal)
}

func routeDeleteStudents(c *gin.Context) {
	var newStudents []Student
	idString := c.Params.ByName("id")
	id, err := shared.GetUuidByString(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Problema com seu o ID",
		})
		return
	}

	for _, studentElement := range Students {
		if id != studentElement.ID {
			newStudents = append(newStudents, studentElement)
		}
	}

	Students = newStudents

	c.JSON(http.StatusOK, gin.H{
		"message": "Estudante excluido com sucesso",
	})
}

func routeGetStudent(c *gin.Context) {
	var student Student
	idString := c.Params.ByName("id")
	id, err := shared.GetUuidByString(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Problema com seu o ID",
		})
		return
	}

	for _, studentElement := range Students {
		if id == studentElement.ID {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi encontrar o estudante",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}