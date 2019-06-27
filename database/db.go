package db 

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
	"log"
	"net/http"
)

type Todo struct {
	Id     int `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}


func (t Todo) DeleteStoreByID(c *gin.Context) {
	param := c.Param("id")
	num, err := strconv.Atoi(param)

	db, ok := c.Value("psql_session").(*sql.DB)
	if !ok {
		log.Fatal("can't open DB")
	}


	stmt, err := db.Prepare("DELETE from todos WHERE id=$1;")
	if err != nil {
		log.Fatal("prepare error", err.Error())
	}

	if _, err := stmt.Exec(num); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot Delete Data might because of SQL Prepareation"})
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (t *Todo) UpdateStoreByID(c *gin.Context) {
	param := c.Param("id")
	num, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot recieve Data"})
		return
	}

	db, ok := c.Value("psql_session").(*sql.DB)
	if !ok {
		log.Fatal("can't open DB")
	}



	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "There's Something Wrong on SQL Preparation"})
		return
	}

	stmt.QueryRow(num, t.Status)
	t.Id = num

	c.JSON(http.StatusOK, t)

}

func (t *Todo) Insert(c *gin.Context) {
	db, ok := c.Value("psql_session").(*sql.DB)
	if !ok {
		log.Fatal("can't open DB")
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



	row := stmt.QueryRow(&t.Title, &t.Status)
	err = row.Scan(&t.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot fetch IDS"})
		return
	}

	c.JSON(http.StatusOK, t)
}


func (t Todo) GetStoreByID(c *gin.Context) {
	param := c.Param("id")
	num, _ := strconv.Atoi(param)
	db, ok := c.Value("psql_session").(*sql.DB)
	if !ok {
		log.Fatal("can't open DB")
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{ "message": "There's Something Wrong on SQL Preparation"})
		return
	}
	row := stmt.QueryRow(num)
	err = row.Scan(&t.Id, &t.Title, &t.Status)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Couldn't find from your ID"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func (t Todo) GetStore(c *gin.Context) {
	var todos []Todo
	db, ok := c.Value("psql_session").(*sql.DB)
	if !ok {
		log.Fatal("can't open DB")
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "There's Something Wrong on SQL Preparation"})
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "There's Something Wrong on SQL Preparation"})
		return
	}

	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Status)
		if err != nil {
			log.Fatal(err.Error())
		}
		todos = append(todos, t)
	}
	c.JSON(http.StatusOK, todos) 
}