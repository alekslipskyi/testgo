package User

import (
	"core/crypto"
	"core/db"
	"core/db/connect"
	"core/db/types"
	"core/logger"
	"encoding/hex"
	"encoding/json"
)

var log = logger.Logger{"User Model"}

type StructUser struct {
	ID        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `access:"private" json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type User struct {
	StructUser
	Token string
}

type Token struct {
	Username string
	ID       int64
}

func (user *User) GenerateToken() {
	token := Token{user.Username, user.ID}
	_, err := json.Marshal(token)

	if err != nil {
		log.Error(err)
	}

	user.Token = hex.EncodeToString(crypto.EncodeToken(json.Marshal(token)))
}

func (user *User) UpdateAndFind(data map[string]interface{}) User {
	return User{StructUser: getInstance().UpdateAndFind(data, types.Where{"_id": user.ID}).(StructUser)}
}

func (user *User) IsNotExist() bool {
	return user.ID == 0
}

func (user *User) IsIPNotExist(ip string) bool {
	userWithIP := FindWithIP(ip, user.ID)
	return userWithIP.IsNotExist()
}

func (user *User) IsValidPassword(password string) bool {
	return user.Password == crypto.GenerateHash(password)
}

func (user *User) AddAllowIP(ip string) bool {
	_, err := connect.DB.Query("insert into ip_users values ($1, $2)", ip, user.ID)

	if err != nil {
		return false
	}

	return true
}

func (user *User) Drop() bool {
	return getInstance().Drop(types.QueryOptions{Where: types.Where{"_id": user.ID}})
}

func getInstance(name ...string) *db.Instance {
	if len(name) != 0 {
		return &db.Instance{name[0], &StructUser{}}
	}
	return &db.Instance{"users", &StructUser{}}
}

func FindWithIP(ip string, userId int64) User {
	rows, err := connect.DB.Query("select u._id from users u left join ip_users iu on u._id = iu.user_id where iu.ip = $1 and _id=$2", ip, userId)

	if err != nil {
		log.Error("error FindWithIP query", err)
		return User{}
	}

	for rows.Next() {
		str := new(StructUser)
		_ = rows.Scan(&str.ID)
		return User{StructUser: *str}
	}

	return User{}
}

func Find(options types.QueryOptions) User {
	return User{StructUser: getInstance().Find(options).(StructUser)}
}

func CreateAndFind(data map[string]interface{}) User {
	return User{StructUser: getInstance().CreateAndFind(data).(StructUser)}
}

func FindById(id int64, attributes []string) User {
	return User{StructUser: getInstance().FindById(id, attributes).(StructUser)}
}
