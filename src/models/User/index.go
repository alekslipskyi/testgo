package User

import (
	"core/db"
	"core/db/types"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lib/crypto"
)

type StructUser struct {
	ID        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `access:"private" json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type User struct {
	StructUser
	db.Instance
	Token    string
	instance db.Instance
}

type Token struct {
	FirstName string
	LastName  string
	ID        int64
}

func (user *User) Test() {
	fmt.Println("test are", user.ID)
}

func (user *User) GenerateToken() {
	token := Token{user.FirstName, user.LastName, user.ID}
	test, _ := json.Marshal(token)

	fmt.Println("test", test)

	user.Token = hex.EncodeToString(crypto.EncodeToken(json.Marshal(token)))
}

func getInstance() *db.Instance {
	return &db.Instance{"users", &StructUser{}}
}

func Find(options types.QueryOptions) User {
	return User{StructUser: getInstance().Find(options).(StructUser)}
}

func FindById(id int8, options types.QueryOptions) User {
	return User{StructUser: getInstance().FindById(id, options).(StructUser)}
}
