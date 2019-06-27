package service 

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"strconv"
	"log"
)

func (t Todo) DeleteStoreByID(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Get Session DB"})
		return
	}

	id := c.Param("id")

	err = delete(id, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H { "message": "failed"})
		return
	}

	c.JSON(200, gin.H { "message": "success"})
}

func delete(id string, session *sql.DB) error {
	num, err := strconv.Atoi(id)

	stmt, err := session.Prepare("DELETE from todos WHERE id=$1;")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(num); err != nil {
		return err
	}

	return nil
}