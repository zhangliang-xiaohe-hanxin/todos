package service 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"strconv"
	"log"
)

func (t Todo) DeleteStoreByID(c *gin.Context) {
	id := c.Param("id")

	status, err := delete(id, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H { "message": "failed"})
		return
	}

	c.JSON(200, gin.H { "message": "success"})
}

func delete(id string, c *gin.Context) error {
	num, err := strconv.Atoi(id)

	session, err := db.GetSession(c)
	if err != nil {
		return err
	}

	stmt, err := session.Prepare("DELETE from todos WHERE id=$1;")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(num); err != nil {
		return err
	}

	return nil
}