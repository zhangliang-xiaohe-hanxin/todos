package handler 

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Route struct {
	API		APIMethod
	DBHost string
}

type APIMethod interface {
	Insert(c *gin.Context) 
	GetStore(c *gin.Context)
	GetStoreByID(c *gin.Context)
	UpdateStoreByID(c *gin.Context)
	DeleteStoreByID(c *gin.Context)
}

func (r Route) psqlPool() gin.HandlerFunc {

	return func (c *gin.Context) {
		db, err := sql.Open("postgres", r.DBHost)
		if err != nil {
			log.Fatal("can't open", err.Error())
		}

		defer db.Close()

		c.Set("session", db)
		c.Next()
	}
}

func (r Route)Init() *gin.Engine {
	apiMethod := r.API
	routes := gin.Default()
	routes.Use(r.psqlPool())
	api := routes.Group("/api")
	api.GET("/todos", apiMethod.GetStore)
	api.GET("/todos/:id", apiMethod.GetStoreByID)
	api.POST("/todos", apiMethod.Insert)
	api.PUT("/todos/:id", apiMethod.UpdateStoreByID)
	api.DELETE("/todos/:id", apiMethod.DeleteStoreByID)

	return routes
}