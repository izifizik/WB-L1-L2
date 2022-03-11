package chain

import "fmt"

type Updater struct {
	Name string
	Next Chain
}

func (d *Updater) Exec(c *Command) {
	if c.Update {
		fmt.Printf("Updater %s already exec command %s\n", d.Name, c.Name)
		return
	}

	c.Update = true
	fmt.Printf("Updater %s exec command %s\n", d.Name, c.Name)
	d.Next.Exec(c)
}

func (d *Updater) SetNext(chain Chain) {
	d.Next = chain
}
