package main 

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
	"strconv"
	"os"
	"net/http"
	_ "github.com/lib/pq"
	"fmt"
)

type APIMethod interface {
	Insert(c *gin.Context) 
	GetStore(c *gin.Context)
	GetStoreByID(c *gin.Context)
}

type Todo struct {
	Id     int `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var (
	hostName = os.Getenv("HOST")
	port = os.Getenv("PORT")
)

func main() {
	var apiMethod APIMethod
	apiMethod = &Todo{}


	r := gin.Default()
	api := r.Group("/api")
	api.GET("/todos", apiMethod.GetStore)
	api.GET("/todos/:id", apiMethod.GetStoreByID)
	api.POST("/todos", apiMethod.Insert)

	r.Run(fmt.Sprintf(":%s", port))
}

func (t *Todo) Insert(c *gin.Context) {
	db, err := sql.Open("postgres", hostName)
	if err != nil {
		log.Fatal("fatal", err.Error())
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot recieve Data"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO todos(title, status) VALUES($1, $2) RETURNING ID;")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "There's Something Wrong on SQL Preparation"})
		return
	}

	stmt.QueryRow(&t.Title, &t.Status)
	c.JSON(http.StatusOK, gin.H{ "message": "Insert Sucessfully"})
}


func (t Todo) GetStoreByID(c *gin.Context) {
	param := c.Param("id")
	num, _ := strconv.Atoi(param)
	db, err := sql.Open("postgres", hostName)
	if err != nil {
		log.Fatal("Cannot connect Database")
	}

	stmt, _ := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	row := stmt.QueryRow(num)
	err = row.Scan(&t.Id, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Couldn't find from your ID"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func (t Todo) GetStore(c *gin.Context) {
	var todos []Todo
	db, _ := sql.Open("postgres", hostName)
	stmt, _ := db.Prepare("SELECT id, title, status FROM todos")

	rows, _ := stmt.Query()
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Status)
		if err != nil {
			log.Fatal(err.Error())
		}
		todos = append(todos, t)
		fmt.Println("data", todos)
	}
	c.JSON(http.StatusOK, todos) 
}