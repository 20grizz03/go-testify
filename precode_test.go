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

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200")

	expected := ""
	require.NotEqual(t, expected, responseRecorder.Body.String(), "Тело ответа пустое")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// Проверяем количество кафе
	require.Equal(t, totalCount, len(list), "Ожидалось %d кафе", totalCount)

	// Проверяем содержимое списка кафе
	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expected, body, "Неправильный список кафе")
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус-код
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался статус 400")

	// Проверяем тело ответа
	expected := "wrong city value"
	require.Equal(t, expected, responseRecorder.Body.String(), "Неправильное тело ответа")
}
