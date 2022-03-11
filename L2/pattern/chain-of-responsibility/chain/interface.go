package chain

type Chain interface {
	Exec(command *Command)
	SetNext(chain Chain)
}

type Command struct {
	Name   string
	Get    bool
	Update bool
}
