package main

import (
	"fmt"
	"state/cmd"
)

/*
Состояние позволяет в зависимости от контекста выбирать реализацию под определенный контекст
*/
func main() {
	mongo := cmd.NewDatabase("MongoDB")
	postgres := cmd.NewDatabase("PSQL")

	svc := cmd.NewService(mongo)
	fmt.Println(svc.FindFromRepo(123))
	svc.InsertToRepo(123, "data")

	svc = cmd.NewService(postgres)
	fmt.Println(svc.FindFromRepo(123))
	svc.InsertToRepo(123, "data")
}
