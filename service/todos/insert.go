package service 

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"net/http"
	"log"
)

func (t *Todo) Insert(c *gin.Context) {
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot recieve Data"})
		return
	}

	err := insert(t, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Insert"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func insert(t *Todo, c *gin.Context) error {

	session, err := db.GetSession(c)
	if err != nil {
		return err
	}

	stmt, err := session.Prepare("INSERT INTO todos(title, status) VALUES($1, $2) RETURNING ID;")
	if err != nil {
		return err
	}

	row := stmt.QueryRow(t.Title, t.Status)
	err = row.Scan(&t.Id)
	if err != nil {
		return err
	}
	
	return nil
}