package serializers

import "github.com/lilosir/cyticoffee-api/models"

// SerializationUser for user serialization
type SerializationUser struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// UserSerializer serialize user object
func UserSerializer(user models.User) SerializationUser {
	return SerializationUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}
