package postgres

import (
	"context"

	"github.com/nikola-enter21/wms/backend/database/model"
)

func (db *serviceDB) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result := db.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUser implements the UserService interface method for retrieving a user.
func (db *serviceDB) GetUser(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	result := db.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db *serviceDB) GetUserByCredentials(ctx context.Context, username, password string) (*model.User, error) {
	var user model.User
	result := db.DB.Where(&model.User{Username: username, Password: password}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db *serviceDB) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result := db.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (db *serviceDB) DeleteUser(ctx context.Context, userID string) error {
	result := db.DB.Where("id = ?", userID).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *serviceDB) ListUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
