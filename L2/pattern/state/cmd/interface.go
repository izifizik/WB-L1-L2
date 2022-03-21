package cmd

type Database interface {
	Insert() bool
	Find() string
}

type Service interface {
	InsertToRepo(uuid int, data string) bool
	FindFromRepo(uuid int) string
}
