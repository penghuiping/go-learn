package db

import (
	"testing"
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func TestDB(in *testing.T)  {
	db,_ := sql.Open("mysql","root:@(127.0.0.1:33306)/golang123?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	rows,err:= db.Query("select * from users")

	if nil != err {
		fmt.Println(err)
	}

	defer rows.Close()


	for rows.Next(){
		var name string
		var id int
		var age int
		var roleId int
		if err := rows.Scan(&id,&name,&age,&roleId); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name,age,roleId)
	}
}



