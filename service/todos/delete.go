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
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Cannot Get Session DB"})
		return
	}

	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H { "status": "failed"})
		return
	}

	err = delete(num, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H { "status": "failed"})
		return
	}

	c.JSON(200, gin.H { "status": "success"})
}

func delete(id int, session *sql.DB) error {

	stmt, err := session.Prepare("DELETE from todos WHERE id=$1;")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}