package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

type Car struct {
	Model string `json:"model" validate:"required"`
	Year  uint   `json:"year" validate:"required"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var q User
		err := c.BindJSON(&q)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error to bind"})
			return
		}

		if !IsValidData(c, q) {
			return
		}

		c.Status(http.StatusAccepted)
	})
	r.POST("/car", func(c *gin.Context) {
		var q Car
		err := c.BindJSON(&q)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error to bind"})
			return
		}

		if !IsValidData(c, q) {
			return
		}

		c.Status(http.StatusAccepted)
	})

	r.Run()
}

func IsValidData(c *gin.Context, i interface{}) bool {
	err := validator.New().Struct(i)
	if err != nil {
		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
			return false
		}

		// Outro tipo de erro, como json como json mal formado, ext
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return false
	}
	return true
}

// Nomes nos campo
func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}
