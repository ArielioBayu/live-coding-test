package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	AddStudent(c *gin.Context)
	GetStudents(c *gin.Context)
	GetStudentByID(c *gin.Context)
}

type studentAPI struct {
	studentRepo repo.StudentRepository
}

func NewStudentAPI(studentRepo repo.StudentRepository) *studentAPI {
	return &studentAPI{studentRepo}
}

func (s *studentAPI) AddStudent(c *gin.Context) {
	var student model.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err = s.studentRepo.Store(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add student success"})
}

func (s *studentAPI) GetStudents(c *gin.Context) {
	students, err := s.studentRepo.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (s *studentAPI) GetStudentByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	students, err := s.studentRepo.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "student not found"})

}
