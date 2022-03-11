package chain

import "fmt"

type Saver struct {
	Next Chain
}

func (d *Saver) Exec(c *Command) {
	if !c.Update {
		fmt.Printf("Command %s has not been exec\n", c.Name)
		return
	}
	fmt.Printf("Command is done\n")
}

func (d *Saver) SetNext(chain Chain) {
	d.Next = chain
}
