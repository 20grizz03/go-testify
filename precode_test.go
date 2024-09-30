package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200")

	expected := ""
	assert.NotEqual(t, expected, responseRecorder.Body.String(), "Тело ответа пустое")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// Проверяем количество кафе
	assert.Len(t, list, totalCount, "Ожидалось %d кафе", totalCount)

}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус-код
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался статус 400")

	// Проверяем тело ответа
	expected := "wrong city value"
	assert.Equal(t, expected, responseRecorder.Body.String(), "Неправильное тело ответа")
}

//поправил
