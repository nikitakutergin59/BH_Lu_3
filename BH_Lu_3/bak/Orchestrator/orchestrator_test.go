package orchestrator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCalculateHandler(t *testing.T) {
	md := NewMemoryData(1 * time.Minute)
	expression := Expression_BH{
		Expression: "2+2",
	}

	body, _ := json.Marshal(expression)
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CalculateHandler(w, r, md)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CalculateHandler вернул статус %v, ожидался %v", status, http.StatusOK)
	}

	expected := `{"result":"4"}`
	if rr.Body.String() != expected {
		t.Errorf("CalculateHandler вернул неожиданный ответ: %v, ожидался %v", rr.Body.String(), expected)
	}
}

func TestLCF_Otvet(t *testing.T) {
	md := NewMemoryData(1 * time.Minute)
	task := Task_BH{
		Task_ID: "123",
	}
	md.Task[task.Task_ID] = &task

	response := LCFOS{
		LCFOS_ID:     "123",
		LCFOS_Result: "4",
	}
	body, _ := json.Marshal(response)
	req := httptest.NewRequest(http.MethodPost, "/lcf_otvet", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LCF_Otvet(w, r, md)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("LCF_Otvet вернул статус %v, ожидался %v", status, http.StatusOK)
	}

	updatedTask := md.Task["123"]
	if updatedTask.Task_Result != "4" || updatedTask.Task_Status != "обработанно" {
		t.Errorf("LCF_Otvet не обновил данные задачи корректно: %+v", updatedTask)
	}
}
