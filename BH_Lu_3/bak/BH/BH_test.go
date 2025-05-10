package BH

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест для RegisterHandler
func TestRegisterHandler(t *testing.T) {
	// Создаём тестовый запрос
	userData := UserData{
		Login:    "test_user",
		Password: "test_password",
	}
	body, _ := json.Marshal(userData)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Создаём тестовый HTTP-ответ
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	handler := http.HandlerFunc(RegisterHandler)
	handler.ServeHTTP(rr, req)

	// Проверяем статус ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("RegisterHandler вернул статус %v, ожидался %v", status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := `{"message_0":"Данные успешно сохранены"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("RegisterHandler вернул неожиданный ответ: %v", rr.Body.String())
	}
}

// Тест для LoginHandler
func TestLoginHandler(t *testing.T) {
	// Создаём тестовый запрос
	userData := UserData{
		Login:    "test_user",
		Password: "test_password",
	}
	body, _ := json.Marshal(userData)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Создаём тестовый HTTP-ответ
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)

	// Проверяем статус ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("LoginHandler вернул статус %v, ожидался %v", status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := `{"message":"успешный вход"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("LoginHandler вернул неожиданный ответ: %v", rr.Body.String())
	}
}

// Тест для CalculateHandlerWithAuth
func TestCalculateHandlerWithAuth(t *testing.T) {
	// Создаём тестовый запрос
	expression := Expression_BH_In_Server{
		Expression: "2+2",
	}
	body, _ := json.Marshal(expression)
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test_token")

	// Создаём тестовый HTTP-ответ
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	handler := http.HandlerFunc(CalculateHandlerWithAuth)
	handler.ServeHTTP(rr, req)

	// Проверяем статус ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CalculateHandlerWithAuth вернул статус %v, ожидался %v", status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := `{"message":"выражение обработано"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("CalculateHandlerWithAuth вернул неожиданный ответ: %v", rr.Body.String())
	}
}
