package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	// Запрос с параметрами count, превышающим общее количество кафе, и city = "moscow"
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedResponse, responseRecorder.Body.String())
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	// Запрос с некорректным городом
	req, err := http.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expectedResponse := "wrong city value"
	assert.Equal(t, expectedResponse, responseRecorder.Body.String())
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	// Запрос с корректными параметрами
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedResponse := "Мир кофе,Сладкоежка"
	assert.Equal(t, expectedResponse, responseRecorder.Body.String())
}
