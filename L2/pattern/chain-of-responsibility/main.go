package main

import "chain-of-responsibility/chain"

/*
Цепочка дает прохождение определенных данных через цепочку обязательных действий логики
(либо множество несвязанных действий, но цепочка-то есть)
*/

func main() {
	device := &chain.Device{
		Name: "Device-1",
	}
	updater := &chain.Updater{
		Name: "Updater-1",
	}
	saver := &chain.Saver{}

	device.SetNext(updater)
	updater.SetNext(saver)

	command := &chain.Command{
		Name: "command-1",
	}

	device.Exec(command)
}
