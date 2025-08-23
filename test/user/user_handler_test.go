package user_test

import (
	"bytes"
	"encoding/json"
	"hexagonal/domain/entity"
	router "hexagonal/internal/adapter/http"
	handler "hexagonal/internal/adapter/http/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUsecase implements UserUsecase port
type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) Register(name, email, password string) error {
	args := m.Called(name, email, password)
	return args.Error(0)
}

func (m *MockUsecase) Login(email, password string) (*entity.User, error) {
	args := m.Called(email, password)
	if user, ok := args.Get(0).(*entity.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUsecase) GetAllUsers() ([]*entity.User, error) {
	args := m.Called()
	if users, ok := args.Get(0).([]*entity.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserHandler_Register(t *testing.T) {
	mockUc := new(MockUsecase)
	h := handler.NewUserHandler(mockUc)
	r := router.NewRouter(h)

	payload := map[string]string{
		"name":     "Test",
		"email":    "reg@example.com",
		"password": "pass123",
	}
	body, _ := json.Marshal(payload)

	mockUc.On("Register", "Test", "reg@example.com", "pass123").Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUserHandler_Login(t *testing.T) {
	mockUc := new(MockUsecase)
	h := handler.NewUserHandler(mockUc)
	r := router.NewRouter(h)

	user, _ := entity.NewUser("LoginUser", "log@example.com", "123")
	mockUc.On("Login", "log@example.com", "123").Return(user, nil)

	payload := map[string]string{
		"email":    "log@example.com",
		"password": "123",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got entity.User
	err := json.NewDecoder(w.Body).Decode(&got)
	assert.NoError(t, err)
	assert.Equal(t, "log@example.com", got.Email)
}

func TestUserHandler_GetAll(t *testing.T) {
	mockUc := new(MockUsecase)
	h := handler.NewUserHandler(mockUc)
	r := router.NewRouter(h)

	mockUsers := []*entity.User{
		{Name: "A"}, {Name: "B"},
	}
	mockUc.On("GetAllUsers").Return(mockUsers, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got []*entity.User
	err := json.NewDecoder(w.Body).Decode(&got)
	assert.NoError(t, err)
	assert.Len(t, got, 2)
}
