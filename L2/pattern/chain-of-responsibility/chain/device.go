package chain

import "fmt"

type Device struct {
	Name string
	Next Chain
}

func (d *Device) Exec(c *Command) {
	if c.Get {
		fmt.Printf("Command %s already taken\n", c.Name)
		d.Next.Exec(c)
		return
	}

	c.Get = true
	fmt.Printf("Device %s take command %s\n", d.Name, c.Name)
	d.Next.Exec(c)
}

func (d *Device) SetNext(chain Chain) {
	d.Next = chain
}
