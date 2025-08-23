package user_test

import (
	"errors"
	"hexagonal/domain/entity"
	"hexagonal/internal/app/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepo implements port.UserRepository
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Save(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) FindByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if user, ok := args.Get(0).(*entity.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepo) GetAll() ([]*entity.User, error) {
	args := m.Called()
	if users, ok := args.Get(0).([]*entity.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserUsecase_Register(t *testing.T) {
	mockRepo := new(MockUserRepo)
	uc := usecase.NewUserUsecase(mockRepo)

	t.Run("should return error if email already exists", func(t *testing.T) {
		mockRepo.On("FindByEmail", "test@example.com").Return(&entity.User{}, nil)

		err := uc.Register("John", "test@example.com", "pass")
		assert.EqualError(t, err, "email already exists")
		mockRepo.AssertCalled(t, "FindByEmail", "test@example.com")
	})

	t.Run("should save user if not exists", func(t *testing.T) {
		mockRepo.On("FindByEmail", "new@example.com").Return(nil, errors.New("not found"))
		mockRepo.On("Save", mock.Anything).Return(nil)

		err := uc.Register("New", "new@example.com", "pass")
		assert.NoError(t, err)
	})
}

func TestUserUsecase_Login(t *testing.T) {
	mockRepo := new(MockUserRepo)
	uc := usecase.NewUserUsecase(mockRepo)

	t.Run("should return user if password matches", func(t *testing.T) {
		user, _ := entity.NewUser("John", "login@example.com", "secret")
		mockRepo.On("FindByEmail", "login@example.com").Return(user, nil)

		u, err := uc.Login("login@example.com", "secret")
		assert.NoError(t, err)
		assert.Equal(t, "login@example.com", u.Email)
	})

	t.Run("should return error if user not found", func(t *testing.T) {
		mockRepo.On("FindByEmail", "missing@example.com").Return(nil, errors.New("not found"))

		_, err := uc.Login("missing@example.com", "pass")
		assert.EqualError(t, err, "user not found")
	})

	t.Run("should return error if password incorrect", func(t *testing.T) {
		user, _ := entity.NewUser("John", "badpass@example.com", "right")
		mockRepo.On("FindByEmail", "badpass@example.com").Return(user, nil)

		_, err := uc.Login("badpass@example.com", "wrong")
		assert.EqualError(t, err, "invalid password")
	})
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepo)
	uc := usecase.NewUserUsecase(mockRepo)

	t.Run("should return all users", func(t *testing.T) {
		users := []*entity.User{
			{Name: "Alice"}, {Name: "Bob"},
		}
		mockRepo.On("GetAll").Return(users, nil)

		result, err := uc.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})
}
