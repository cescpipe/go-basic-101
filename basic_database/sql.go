package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Customer struct {
	CustomerCode string `db:"customer_code"`
	CustomerName string `db:"customer_name"`
}

func connectPostgres() error {
	conn := "postgres://azfravbnwomtof:5f04e40f5fece3f0a3245df4402dd66cc29cca642a2ec3e08d7151f04536c185@ec2-54-145-249-177.compute-1.amazonaws.com:5432/defsb694qnpjb0"
	fmt.Println("connection:", conn)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println("open connection error:", err)
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Println("ping error:", err)
		return err
	}

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("version:", version)

	rows, err := db.Query("SELECT cust_code, cust_name FROM customers")
	if err != nil {
		log.Println("query error:", err)
		return err
	}
	defer rows.Close()

	var ags []Customer
	for rows.Next() {
		var data Customer
		rows.Scan(&data.CustomerCode, &data.CustomerName)

		ags = append(ags, data)
	}

	for _, a := range ags {
		log.Println("customer:", a.CustomerCode, a.CustomerName)
	}

	var ag Customer
	if err := db.QueryRow("SELECT cust_name FROM customers WHERE cust_code = $1", "C00013").Scan(&ag.CustomerName); err != nil {
		log.Println("query row error:", err)
		return err
	}
	log.Println("single customer:", ag.CustomerName)
	return nil
}
