package service 

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"net/http"
	"log"
)

func (t Todo) GetStore(c *gin.Context) {

	todos, err := getAll(t, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot query"})
		return
	}

	c.JSON(http.StatusOK, todos) 
}

func getAll(t Todo, c *gin.Context) ([]Todo, error) {

	var todos []Todo

	session, err := db.GetSession(c)
	if err != nil {
		return []Todo{}, err
	}

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