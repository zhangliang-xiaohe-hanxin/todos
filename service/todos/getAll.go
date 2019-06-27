package service 

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"net/http"
	"log"
)

func (t Todo) GetStore(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot Get Session DB"})
		return
	}

	todos, err := getAll(t, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot query"})
		return
	}

	c.JSON(http.StatusOK, todos) 
}

func getAll(t Todo, session *sql.DB) ([]Todo, error) {

	var todos []Todo

	stmt, err := session.Prepare("SELECT id, title, status FROM todos")

	if err != nil {
		return []Todo{}, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return []Todo{}, err
	}

	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Status)
		if err != nil {
			return []Todo{}, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}