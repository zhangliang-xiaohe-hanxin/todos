package main 

import (
	"github.com/zhangliangxiaohehanxin/todos/route"
	"github.com/zhangliangxiaohehanxin/todos/database"
	"fmt"
)

const (
	hostName = "postgres://dpsdgjur:qf4v1Qap7DKwpK3ZySXEWa7rB6B-VsJF@satao.db.elephantsql.com:5432/dpsdgjur"
	port = "1234"
)

func main() {
	todo := &db.Todo{}
	r := handler.Route{ todo, hostName}.Init()
	r.Run(fmt.Sprintf(":%s", port))
}
