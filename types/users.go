package types

type Users struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
