package service 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"strconv"
	"log"
)

func (t Todo) GetStoreByID(c *gin.Context) {
	id := c.Param("id")
	
	err := getByID(&t, id, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Query"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func getByID(t *Todo, id string, c *gin.Context) error {

	num, _ := strconv.Atoi(id)
	session, err := db.GetSession(c)
	if err != nil {
		return err
	}

	stmt, err := session.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		return err
	}
	row := stmt.QueryRow(num)
	err = row.Scan(&t.Id, &t.Title, &t.Status)
	if err != nil {
		return err
	}

	return nil
}
