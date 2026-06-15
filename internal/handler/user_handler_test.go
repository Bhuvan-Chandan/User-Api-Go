package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	db "go-user-api/db/sqlc"

	"github.com/gofiber/fiber/v2"
)

// fake service (NO DB, NO SQLC)
type fakeService struct{}

func (f *fakeService) CreateUser(name, dob string) (db.User, error) {
	return db.User{
		ID:   1,
		Name: name,
	}, nil
}

func (f *fakeService) GetUserWithAge(id int32) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (f *fakeService) ListUsers(page, limit int) ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

func (f *fakeService) UpdateUser(id int32, name string, dob string) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (f *fakeService) DeleteUser(id int32) error {
	return nil
}

func TestCreateUserHandler(t *testing.T) {

	app := fiber.New()

	// inject fake service safely
	h := &UserHandler{
		Service: &fakeService{},
	}

	app.Post("/users", h.CreateUser)

	body := map[string]string{
		"name": "Alice",
		"dob":  "1990-01-01",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
}
