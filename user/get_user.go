package user

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int
	name string
}

func GetUser() {
	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/Test_Mirai")
	if err != nil {
		log.Fatalf("getUser sql.Open error err:%v", err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.name)
		if err != nil {
			panic(err.Error())
		}
		log.Println(user.id, user.name)
	}
}
