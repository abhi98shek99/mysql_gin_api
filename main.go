package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	router.GET("/person/:id/info", getdata)
	router.POST("/person/create", postdata)
	router.Run(":1000")
}

//Task 1 to create a GET API Using Gin Mysql
func getdata(c *gin.Context) {
	db, err := sql.Open("mysql", "root:abhishek@tcp(127.0.0.1:3306)/cetec")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// To make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Person struct {
		Name         string
		Phone_number string
		City         string
		State        string
		Street1      string
		Street2      string
		Zip_code     string
	}
	var (
		person Person
		result Person
	)
	id := c.Param("id")
	fmt.Println(id)
	row := db.QueryRow("select person.name,phone.number,address.city,address.state,address.street1,address.street2,address.zip_code from person join phone on person.id=phone.person_id join address_join on address_join.person_id=person.id join address on address.id=address_join.address_id where person.id = ?;", id)
	err = row.Scan(&person.Name, &person.Phone_number, &person.City, &person.State, &person.Street1, &person.Street2, &person.Zip_code)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		result = person
	}
	c.JSON(http.StatusOK, result)
}

// Task 2 to create a POST Api using Gin and Mysql
func postdata(c *gin.Context) {
	db, err := sql.Open("mysql", "root:abhishek@tcp(127.0.0.1:3306)/cetec")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	name := c.PostForm("abhishek")
	number := c.PostForm("12")
	address := c.PostForm("4")
	city := c.PostForm("Sacramento")
	state := c.PostForm("CA")
	street1 := c.PostForm("112 Main St")
	street2 := c.PostForm("Apt 12")
	zip_code := c.PostForm("12345")
	insert, err := db.Query("insert into from person join phone on person.id=phone.person_id join address_join on address_join.person_id=person.id join address on address.id=address_join.address_id (person.name,phone.number,address.city,address.state,address.street1,address.street2,address.zip_code) values(?,?);", name, number, city, address, state, street1, street2, zip_code)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer insert.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": "200",
	})
}
