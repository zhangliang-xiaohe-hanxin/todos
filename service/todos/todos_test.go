package service 

import (
	"testing"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	hostName = "postgres://dpsdgjur:qf4v1Qap7DKwpK3ZySXEWa7rB6B-VsJF@satao.db.elephantsql.com:5432/dpsdgjur"
)

var testingId int 


func TestTodo(t *testing.T) {
	t.Run("Insert", TestInsert)
	t.Run("Get", TestGetById)
	t.Run("Update", TestUpdate)
	t.Run("Delete", TestDelete)
}

func TestInsert(t *testing.T) {

	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	todo := Todo{ Title: "John Denver", Status: "Dead"} 
	err = insert(&todo, session)
	if err != nil {
		t.Errorf("Cannot Insert Data    Errror: %v", err)
	}

	if todo.Id == 0 {
		t.Errorf("Got Zero Values in Id Field ")
	}

	testingId = todo.Id

}

func TestGetById(t *testing.T) {

	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	var todo Todo

	err = getByID(&todo, testingId, session)
	if err != nil {
		t.Errorf("Cannot get by Id  Errror: %v", err)
	}

	if todo.Id != testingId && todo.Title != "John Denver"  && todo.Status != "Dead" {
		t.Errorf("Data is not the same as Insert's Data")
	}
}

func TestUpdate(t *testing.T) {

	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	todo := Todo{ Title: "John Denver", Status: "Alive"} 
	err = updateByID(&todo, testingId, session)
	if err != nil {
		t.Errorf("Cannot Update  Errror: %v", err)
	}

	if todo.Id != testingId && todo.Title != "John Denver"  && todo.Status != "Alive" {
		t.Errorf("Data hasn't been updated yet")
	}
}

func TestDelete(t *testing.T) {
	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	err = delete(testingId, session)
	if err != nil {
		t.Errorf("Cannot Delete  Errror: %v", err)
	}

	var todo Todo

	err = getByID(&todo, testingId, session)
	if err == nil {
		t.Errorf("Data hasn't been deleted yet: %v", err)
	}

}