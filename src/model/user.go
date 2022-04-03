package model

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"src/config"
	"time"
)

var db *pg.DB

type User struct {
	tableName  struct{}  `pg:"users"`
	Id         int       `json:"id"`
	Email      string    `json:"email" pg:",unique,notnull"`
	Username   string    `json:"username" pg:",unique,notnull"`
	PwdHash    string    `json:"pwd"`
	CreateTime time.Time `json:"createTime" pg:"default:now()"`
	Role       bool      `json:"role" pg:",use_zero"` //0:default 1:admin
}

func (u *User) String() string {
	return fmt.Sprintf("User<%v %v %v %v %v>", u.Id, u.Email, u.Username, u.CreateTime, u.Role)
}

// Insert user data
func (u *User) Insert() error {
	_, err := db.Model(u).Insert()
	if err != nil {
		return err
	}
	return nil
}

// Update user data
func Update(id int, email string, username string, pwdHash string) error {
	if email != "" {
		_, err := db.Model(User{}).Set("email = ?", email).Where("id = ?", id).Update()
		if err != nil {
			return err
		}
	}
	if username != "" {
		_, err := db.Model(User{}).Set("username = ?", username).Where("id = ?", id).Update()
		if err != nil {
			return err
		}
	}
	if pwdHash != "" {
		_, err := db.Model(User{}).Set("pwdHash = ?", pwdHash).Where("id = ?", id).Update()
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckUser checks email & password_hash
func CheckUser(email string, pwdHash string) (*User, error) {
	u := new(User)
	err := db.Model(u).
		Where("email = ?", email).
		Where("pwd_hash = ?", pwdHash).
		Select()
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	return u, nil
}

// GetUser returns user info by id.
func GetUser(id int) (*User, error) {
	u := &User{Id: id}
	err := db.Model(u).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return u, nil
}

// SelectAllUser returns all users' info
func SelectAllUser() ([]User, error) {
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUser user data
func DeleteUser(id int) error {
	u := &User{Id: id}
	_, err := db.Model(u).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

// Connect database
func Connect() *pg.DB {
	db = pg.Connect(&pg.Options{
		User:     config.User,
		Password: config.Password,
		Database: config.Dbname,
	})

	var n int
	if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
		panic(err)
	}

	return db
}

// Close database
func Close() {
	db.Close()
}

// CreateSchema creates database schema for User model
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
