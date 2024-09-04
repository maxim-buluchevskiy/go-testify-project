package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	// Запрос с параметрами count, превышающим общее количество кафе, и city = "moscow"
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус-код %v, но получен %v", http.StatusOK, status)
	}

	expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	responseBody := responseRecorder.Body.String()

	assert.Equal(t, expectedResponse, responseBody, "Ответ не совпадает с ожидаемым")

	assert.NotEmpty(t, responseBody, "Ответ не должен быть пустым")
	assert.Len(t, strings.Split(responseBody, ","), totalCount, "Количество кафе в ответе должно быть равно totalCount")
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	// Запрос с некорректным городом
	req, err := http.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался статус-код 400")

	expectedResponse := "wrong city value"
	responseBody := responseRecorder.Body.String()

	assert.Equal(t, expectedResponse, responseBody, "Ответ не совпадает с ожидаемым")
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	// Запрос с корректными параметрами
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус-код 200")

	expectedResponse := "Мир кофе,Сладкоежка"
	responseBody := responseRecorder.Body.String()

	assert.Equal(t, expectedResponse, responseBody, "Ответ не совпадает с ожидаемым")
	assert.NotEmpty(t, responseBody, "Ответ не должен быть пустым")
}
