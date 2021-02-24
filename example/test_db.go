package main

import (
	"database/sql"
	"fmt"

	// "log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Customer struct {
	Id       int
	Email    string
	Balance  float64
	Active   bool
	JoinDate string
}

func main() {
	host := getEnv("DATABASE_HOST", "localhost")
	port, _ := strconv.ParseInt(getEnv("DATABASE_PORT", "5432"), 10, 64)
	user := getEnv("DATABASE_USERNAME", "postgres")
	dbname := "go_rest_api"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password='' dbname=%s sslmode=disable",
		host, port, user, dbname)

	println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select id, email from customers;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.Id, &customer.Email)
		if err != nil {
			panic(err)
		}
		fmt.Println(customer.Email)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// sqlSt := `INSERT INTO customers (id,  email, balance, active, join_date) VALUES ($1, $2, $3, $4, $5);`
	// _, err = db.Exec(sqlSt, 90, "max@gmail.com", 89.99, true, "2020-12-25")
	// if err != nil {
	// 	panic(err)
	// }
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
