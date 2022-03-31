package models

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "080502"
	dbname   = "test"
)

var db *pg.DB

type User struct {
	Id    int64
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email)
}

func InsertUser(u *User) {
	//u := &User{
	//	Name:  "admin",
	//	Email: "admin1@admin",
	//}
	_, err := db.Model(u).Insert()
	if err != nil {
		print("insert error")
		panic(err)
	}
	fmt.Printf("success")
}

func SelectName(name string) *User {
	// Select user by primary key.
	user := &User{Name: name}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	return user
}

func SelectId(id int64) *User {
	// Select user by primary key.
	user := &User{Id: id}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	return user
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

func Insert() {
	user1 := &User{
		Name:  "admin",
		Email: "admin1@admin",
	}
	_, err := db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}
