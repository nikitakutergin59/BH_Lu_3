package Lu

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLCF(t *testing.T) {
	task := TaskInLu{
		TIL_ID:    "123",
		TIL_Value: "2+2",
	}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest(http.MethodPost, "/LCF", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(LCF)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("LCF вернул статус %v, ожидался %v", status, http.StatusOK)
	}

	expected := `{"in_lcf_id":"123","in_lcf_result":"4"}`
	if rr.Body.String() != expected {
		t.Errorf("LCF вернул неожиданный ответ: %v, ожидался %v", rr.Body.String(), expected)
	}
}
