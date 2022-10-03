package accounts

import "encoding/json"

//POST
// model data
type User struct {
	Name     string      `json:"name" binding:"required"`
	Password json.Number `json:"password" binding:"required,number"`
}
