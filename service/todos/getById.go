package service 

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"strconv"
	"log"
)

func (t Todo) GetStoreByID(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Get Session DB"})
		return
	}

	id := c.Param("id")
	
	err = getByID(&t, id, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Find by You ID"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func getByID(t *Todo, id string, session *sql.DB) error {

	num, _ := strconv.Atoi(id)

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
