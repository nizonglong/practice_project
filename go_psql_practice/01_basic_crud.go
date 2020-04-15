package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func sqlOpen() {
	var err error
	db, err = sql.Open("postgres", "port=54321 user=postgres password=123456 dbname=study sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	//还可以是这种方式打开
	//db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	checkErr(err)
}
func sqlInsert() {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO iuser(id,name,age) VALUES($1,$2,$3) RETURNING id")
	checkErr(err)

	res, err := stmt.Exec(4, "李4", 23)
	res, err = stmt.Exec(5, "李5", 25)
	//这里的三个参数就是对应上面的$1,$2,$3了

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func sqlDelete() {
	//删除数据
	stmt, err := db.Prepare("delete from iuser where id=$1")
	checkErr(err)

	res, err := stmt.Exec(4)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func sqlSelect() {
	//查询数据
	rows, err := db.Query("SELECT * FROM iuser")
	checkErr(err)

	fmt.Println("-----------")
	for rows.Next() {
		var id int
		var name sql.NullString
		var age sql.NullInt64
		var birthday sql.NullString
		var motto sql.NullString
		var email sql.NullString
		err = rows.Scan(&id, &name, &age, &birthday, &motto, &email)
		checkErr(err)
		fmt.Printf("%d, %s, %d, %s, %s, %s -----------\n", id, name, age, birthday, motto, email)
	}
}
func sqlUpdate() {
	//更新数据
	stmt, err := db.Prepare("update iuser set motto=$1 where id=$2")
	checkErr(err)

	res, err := stmt.Exec("li 5 motto", 5)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func sqlClose() {
	_ = db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sqlTest() {

	sep := "----------\n"
	sqlOpen()
	fmt.Println(sep, "*sqlOpen")

	sqlSelect()
	fmt.Println(sep, "*sqlSelect")

	sqlInsert()
	sqlSelect()
	fmt.Println(sep, "*sqlInsert")

	sqlUpdate()
	sqlSelect()
	fmt.Println(sep, "*sqlUpdate")

	sqlDelete()
	sqlSelect()
	fmt.Println(sep, "*sqlDelete")

	sqlClose()
	fmt.Println(sep, "*sqlClose")
}

func main() {

	sqlTest()
}
