package utils

import (
	"encoding/json"
	"net/http"
)

// Response adalah struktur standar untuk semua respon API JSON
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  // omitempty agar tidak muncul jika nil
	Error   string      `json:"error,omitempty"` // omitempty agar tidak muncul jika nil
}

// writeJSON adalah helper internal untuk menulis respon JSON
func writeJSON(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// RespondSuccess mengirimkan respon sukses (HTTP 200-299)
func RespondSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	writeJSON(w, statusCode, resp)
}

// RespondError mengirimkan respon error (HTTP 400-599)
func RespondError(w http.ResponseWriter, statusCode int, message string, err string) {
	resp := Response{
		Success: false,
		Message: message,
		Error:   err,
	}
	writeJSON(w, statusCode, resp)
}