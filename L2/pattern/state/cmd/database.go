package cmd

import "fmt"

type database struct {
	dbType string
}

func NewDatabase(dbType string) Database {
	return &database{dbType: dbType}
}

func (d database) Insert() bool {
	fmt.Printf("%s: insert\n", d.dbType)
	return true
}

func (d database) Find() string {
	return fmt.Sprintf("%s: find", d.dbType)
}
