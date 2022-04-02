package models

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "080502"
	dbname   = "novicetask"
)

var db *pg.DB

type User struct {
	tableName  struct{}  `pg:"users"`
	Id         int       `json:"id"`
	Email      string    `json:"email" pg:",unique,notnull"`
	Username   string    `json:"username" pg:",unique,notnull"`
	PwdHash    string    `json:"pwdHash"`
	CreateTime time.Time `json:"createTime" pg:"default:now()"`
	Role       bool      `json:"role" pg:",use_zero"` //0:default 1:admin
}

func (u *User) String() string {
	return fmt.Sprintf("User<%v %v %v %v %v>", u.Id, u.Email, u.Username, u.CreateTime, u.Role)
}

func (u *User) Insert() error {
	_, err := db.Model(u).Insert()
	if err != nil {
		return err
	}
	return nil
}

func CheckUser(email string, pwdHash string) (bool, error) {
	u := new(User)
	err := db.Model(u).
		Where("email = ?", email).
		Where("pwd_hash = ?", pwdHash).
		Select()
	if err != nil {
		return false, err
	}
	if u == nil {
		return false, nil
	}
	return true, nil
}

// SelectId select user by id.
func GetUser(id int) (*User, error) {
	u := &User{Id: id}
	err := db.Model(u).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Update(id int) error {
	user := &User{Id: id}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		return err
	}
	if len(u.Email) != 0 {
		user.Email = u.Email
	}
	if len(u.Username) != 0 {
		user.Username = u.Username
	}
	if len(u.PwdHash) != 0 {
		user.PwdHash = u.PwdHash
	}
	_, err = db.Model(user).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}

func SelectAllUser() []User {
	// Select all users.
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		panic(err)
	}
	return users
}

func (u *User) Delete() error {
	_, err := db.Model(u).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

func Connect() *pg.DB {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: dbname,
	})

	var n int
	if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
		panic(err)
	}

	return db
}

func Close() {
	db.Close()
}

// createSchema creates database schema for User and Story models.
func CreateSchema() error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
