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
		c.JSON(http.StatusBadRequest, gin.H { "message": status})
		return
	}

	c.JSON(200, gin.H { "message": status})
}

func delete(id string, c *gin.Context) (string, error) {
	num, err := strconv.Atoi(id)

	session, err := db.GetSession(c)
	if err != nil {
		return "failed", err
	}

	stmt, err := session.Prepare("DELETE from todos WHERE id=$1;")
	if err != nil {
		return "failed", err
	}

	if _, err := stmt.Exec(num); err != nil {
		return "failed", err
	}

	return "success", nil
}