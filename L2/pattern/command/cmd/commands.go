package cmd

type insertCommand struct {
	Repository
}

func NewInsertCommand(repository Repository) Command {
	return &insertCommand{Repository: repository}
}

func (c insertCommand) execute() {
	c.Repository.Insert()
}

/*-----------------------------Update------------------------------*/

type updateCommand struct {
	Repository
}

func NewUpdateCommand(repository Repository) Command {
	return &updateCommand{Repository: repository}
}

func (c updateCommand) execute() {
	c.Repository.Update()
}

/*-----------------------------Delete------------------------------*/

type deleteCommand struct {
	Repository
}

func NewDeleteCommand(repository Repository) Command {
	return &deleteCommand{Repository: repository}
}

func (c deleteCommand) execute() {
	c.Repository.Delete()
}

/*------------------------------Find-------------------------------*/

type findCommand struct {
	Repository
}

func NewFindCommand(repository Repository) Command {
	return &findCommand{Repository: repository}
}

func (c findCommand) execute() {
	c.Repository.Find()
}
