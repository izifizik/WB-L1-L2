package main

import (
	"command/cmd"
	"log"
)

/*
Команда дает функции обратного действия в ответ на определенные действия -> очередь из запросов
(с потенциальной отменой)
*/
func main() {
	exec := cmd.NewExecutor()
	db := cmd.NewRepo()
	iCommand := cmd.NewInsertCommand(db)
	uCommand := cmd.NewUpdateCommand(db)
	dCommand := cmd.NewDeleteCommand(db)
	fCommand := cmd.NewFindCommand(db)

	exec.RegisterCommand("insert", iCommand)
	exec.RegisterCommand("update", uCommand)
	exec.RegisterCommand("delete", dCommand)
	exec.RegisterCommand("find", fCommand)

	err := exec.Execute("insert")
	if err != nil {
		log.Println(err.Error())
	}
	err = exec.Execute("update")
	if err != nil {
		log.Println(err.Error())
	}
	err = exec.Execute("delete")
	if err != nil {
		log.Println(err.Error())
	}
	err = exec.Execute("find")
	if err != nil {
		log.Println(err.Error())
	}
}
