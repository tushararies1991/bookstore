package user

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			Status:    user.Status,
		}
	}

	usrJSON, _ := json.Marshal(user)
	var privateUsr PrivateUser
	json.Unmarshal(usrJSON, &privateUsr)
	return privateUsr
}

// Wny value in users marsahlling and pointer in User marshall
func (users Users) Marshall(isPublic bool) []interface{} {
	marsheledUsrs := make([]interface{}, len(users))

	for index, user := range users {
		marsheledUsrs[index] = user.Marshall(isPublic)
	}

	return marsheledUsrs
}
