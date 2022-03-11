package cmd

import "fmt"

type Executor struct {
	commandMap map[string]Command
}

func NewExecutor() Executor {
	m := make(map[string]Command)
	return Executor{commandMap: m}
}

func (e Executor) RegisterCommand(str string, command Command) {
	e.commandMap[str] = command
}

func (e Executor) Execute(str string) error {
	command, ok := e.commandMap[str]
	if !ok {
		return fmt.Errorf("incorrect or not exsist")
	}

	command.execute()
	return nil
}
