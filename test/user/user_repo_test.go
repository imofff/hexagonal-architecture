package user_test

import (
	"hexagonal/domain/entity"
	"hexagonal/internal/adapter/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.User{}))
	return db
}

func TestUserRepo_Save(t *testing.T) {
	db := setupTestDB(t)
	repo := postgres.NewUserRepo(db)

	user, _ := entity.NewUser("Test", "test@example.com", "secret")
	err := repo.Save(user)
	assert.NoError(t, err)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := postgres.NewUserRepo(db)

	// arrange
	user, _ := entity.NewUser("Test", "test@example.com", "secret")
	_ = repo.Save(user)

	// act
	found, err := repo.FindByEmail("test@example.com")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, user.Email, found.Email)
	assert.True(t, found.ComparePassword("secret"))
}

func TestUserRepo_GetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := postgres.NewUserRepo(db)

	db.Exec("DELETE FROM users") // clear for safety

	u1, _ := entity.NewUser("A", "a@example.com", "pass")
	u2, _ := entity.NewUser("B", "b@example.com", "pass")
	_ = repo.Save(u1)
	_ = repo.Save(u2)

	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}
