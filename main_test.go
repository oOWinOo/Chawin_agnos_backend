package main

import (
	"bytes"
	"fmt"
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
	gin.SetMode(gin.ReleaseMode)
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
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}
func TestEx1(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "aA1"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":3}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func TestEx2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "1445D1cd"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":0}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test4CharAdd2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "abcd"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":2}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}
func Test4CharAdd3(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "...."}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":3}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test5CharAdd1(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "abcdE"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":1}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test5CharAdd2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "abcde"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":2}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test5CharAdd2_2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "AAAAA"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":2}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test5CharAdd3(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "!.!.!"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":3}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test6CharNoAdd(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "Abc44a"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":0}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test6CharAdd1(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "abcdeF"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":1}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}


func Test6CharAdd1_2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "1beeeF"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":1}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test6CharAdd2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "baaabb"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":2}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test13CharAdd0(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "AbddS2rewf355"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":0}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test13CharAdd1(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "AbddS222rf355"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":1}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test13CharAdd4(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "aaauuuiiioooq"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":4}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test17CharAdd5(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "11122244477755555"}` // replace the value of middle to make its not repeat

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":5}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test28Remove8(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "A234567890abcdefghij12345678"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":8}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test28Remove8_2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "A234567890abcdefghij12345666"}` // 1 repeat 3 in the row and 28 char BUT use only 8 step(remove)

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":8}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test62Remove42(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":42}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}

func Test130Remove110Replace2(t *testing.T) {

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	database.ConnectToDataBase()
	router.POST("/api/strong_password_steps", controllers.SubmitPassword)
	input := `{"init_password": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"}`

	req, err := http.NewRequest("POST", "/api/strong_password_steps",  bytes.NewBufferString(input))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	output := `{"num_of_steps":112}`
	assert.Equal(t, output, w.Body.String())
	fmt.Printf("The input : %s \n",input)
	fmt.Printf("The output : %s \n\n",output)
}
