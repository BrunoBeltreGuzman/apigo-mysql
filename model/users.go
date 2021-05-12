package model

import (
	"gomysql/types"
)

func GetAllUsers() (AllUsers []types.Users) {
	AllUsers = []types.Users{
		{Id: 1, Name: "Bobo", Email: "bobo@gmail.com", Password: "123", Created_at: "51/564/54", Updated_at: "52/654/254"},
		{Id: 2, Name: "Bobo", Email: "bobo@gmail.com", Password: "123", Created_at: "51/564/54", Updated_at: "52/654/254"}, {Id: 3, Name: "Bobo", Email: "bobo@gmail.com", Password: "123", Created_at: "51/564/54", Updated_at: "52/654/254"},
	}
	return AllUsers
}

func GetByIdUsers(id int) (AllUsers []types.Users) {
	AllUsers = []types.Users{
		{Id: 1, Name: "Bobo", Email: "bobo@gmail.com", Password: "123", Created_at: "51/564/54", Updated_at: "52/654/254"},
		{Id: 2, Name: "Bobo", Email: "bobo@gmail.com", Password: "123", Created_at: "51/564/54", Updated_at: "52/654/254"}, {Id: 3, Name: "Bobo", Email: "bobo@gmail.com", Password: "123", Created_at: "51/564/54", Updated_at: "52/654/254"},
	}

	for _, user := range AllUsers {
		if user.Id == id {
			arrUser := []types.Users{
				user,
			}
			return arrUser
		}
	}

	return nil
}
