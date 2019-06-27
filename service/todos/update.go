package service 

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"strconv"
	"log"
	"net/http"
)

func (t *Todo) UpdateStoreByID(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot Get Session DB"})
		return
	}

	id := c.Param("id")
	num, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID Format"})
		return
	}
	
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot recieve Data"})
		return
	}

	err = updateByID(t, num, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Update"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func updateByID(t *Todo, id int, session *sql.DB) error {

	stmt, err := session.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		return err
	}

	stmt.QueryRow(id, t.Status)
	t.Id = id

	return nil
}