package user

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   string
	Pass string
	Name string
}

func GetUser(id string) User { //DBからユーザーのデータ持ってくる関数
	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/AssembRe")
	if err != nil {
		log.Fatalf("getUser sql.Open error err:%v", err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id,password,name FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var user User
	for rows.Next() {

		err := rows.Scan(&user.Id, &user.Pass, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Println(user.Id, user.Pass, user.Name)
	}
	return user
}
