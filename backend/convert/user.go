package convert

import (
	userv1 "github.com/nikola-enter21/wms/backend/api/user/v1"
	"github.com/nikola-enter21/wms/backend/database/model"
)

func CreateUserRequestToModel(in *userv1.CreateUserRequest) (*model.User, error) {
	return &model.User{
		FullName: in.User.FullName,
		Username: in.User.Username,
		Password: in.User.Password,
		Role:     model.Role(in.User.Role),
	}, nil
}

func UserModelToProto(user *model.User) (*userv1.User, error) {
	return &userv1.User{
		Id:       user.Base.ID,
		FullName: user.FullName,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func UsersModelToProto(dbUsers []*model.User) ([]*userv1.User, error) {
	users := []*userv1.User{}

	for _, v := range dbUsers {
		protoUser, err := UserModelToProto(v)
		if err != nil {
			return nil, err
		}

		users = append(users, protoUser)
	}

	return users, nil
}
