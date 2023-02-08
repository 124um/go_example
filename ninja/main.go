package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Sex   string
	Phone string
	Email string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	errFatal(err)

	defer db.Close()

	err = db.Ping()	
	errFatal(err)

	fmt.Println("Server connect to DB")

	err = insertUser(db, User{
		Name: "Mirra",
		Sex: "female",
		Phone: "=43424434343434",
		Email: "mirror@mi.com",
	})
	errFatal(err)

    // err = deleteUser(db, 5)
	// errFatal(err)

	// err = updateUser(db, 4, User{
	// 	Name: "Sosyca",
	// 	Sex: "female",
	// 	Phone: "+66666666",
	// 	Email: "heal@mi.com",
	// })

	// errFatal(err)

	users, err := getUsers(db)
	errFatal(err)
	fmt.Println("Select getUsers(db) - ", users)

	// user, err := getUser(db, 2)
	// errFatal(err)
	// fmt.Println("Select getUser(db, 2) - ", user)



}

func errFatal(err error){
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			fmt.Println("No rows")
			return
		}
		log.Fatal(err)
	}
}

func getUsers(db *sql.DB)([]User, error){
	rows, err := db.Query("select * from maq.users")
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		u:= User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Sex, &u.Phone, &u.Email)

		users = append(users, u )
	}

	err = rows.Err()
	return users, err
}

func getUser(db *sql.DB, id int) (User, error){
	var u User
	err := db.QueryRow("select * from maq.users where id = $1", id).
		Scan(&u.ID, &u.Name, &u.Sex, &u.Phone, &u.Email)

	return u, err
}

func insertUser(db *sql.DB, u User) (error){
	tx, err := db.Begin()
	errFatal(err)
	defer tx.Rollback()

	returning, err := tx.Exec("insert into maq.users (name, sex, phone, email) values ( $1, $2, $3, $4 ) RETURNING *",
		u.Name, u.Sex, u.Phone, u.Email)	


	_, err = tx.Exec("insert into maq.logs (entity, action) values ( $1, $2 ) RETURNING *",
	"user", "created")		
	// returning, err := db.Exec("insert into maq.users (name, sex, phone, email) values ( $1, $2, $3, $4 ) RETURNING *",
	// 	 u.Name, u.Sex, u.Phone, u.Email)	
	errFatal(err)
	fmt.Println(" User created - ", returning)
	return tx.Commit()
}

func deleteUser(db *sql.DB, id int) (error) {
	_ , err := db.Exec("delete from maq.users where id = $1 RETURNING *", id)
	errFatal(err)
	fmt.Println(" User deleted - ", fmt.Sprintf("User %s is deleted", id))
	return err
}

func updateUser(db *sql.DB, id int, newUser User) error {
	_ , err := db.Exec("update maq.users set name=$1, sex=$2, phone=$3, email=$4 where id = $5 RETURNING *",
		 newUser.Name, newUser.Sex, newUser.Phone, newUser.Email, id)
	errFatal(err)
	fmt.Println(" User deleted - ", fmt.Sprintf("User %s is update", id))
	return err
}