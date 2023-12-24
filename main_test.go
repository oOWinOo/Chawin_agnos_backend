package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oOWinOo/Chawin_agnos_backend/controllers"
	"github.com/oOWinOo/Chawin_agnos_backend/database"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	router := gin.New()
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "value"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":2}`
	assert.Equal(t, output, w.Body.String())
}
