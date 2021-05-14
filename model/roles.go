package model

import (
	"gomysql/database"
	"gomysql/types"
	"log"
)

func GetRole(id int) []types.Roles {
	connection := database.GetConnection()
	const sql = "SELECT * FROM roles WHERE id=?"
	results, error := connection.Query(sql, 1)
	if error != nil {
		log.Fatal(error)
		return []types.Roles{}
	}

	var role = types.Roles{}
	var roles = []types.Roles{}
	for results.Next() {
		err := results.Scan(&role.Id, role.Role, &role.Created_at, &role.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		roles = append(roles, role)
	}

	defer results.Close()
	return roles
}
