package models

type User struct {
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	FullName string `bson:"full_name,omitempty" json:"full_name,omitempty"`
}

type UserRequest struct {
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	FullName string `bson:"full_name,omitempty" json:"full_name,omitempty"`
}

type UserResponse struct {
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	FullName string `bson:"full_name,omitempty" json:"full_name,omitempty"`
}

func (u UserRequest) ToUser() User {
	return User{
		Identity: u.Identity,
		Username: u.Username,
		FullName: u.FullName,
	}
}

func (u User) ToUserResponse() UserResponse {
	return UserResponse{
		Identity: u.Identity,
		Username: u.Username,
		FullName: u.FullName,
	}
}