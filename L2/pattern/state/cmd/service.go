package cmd

type service struct {
	database Database
}

func NewService(db Database) Service {
	return &service{database: db}
}

func (s service) FindFromRepo(uuid int) string {
	return s.database.Find()
}

func (s service) InsertToRepo(uuid int, data string) bool {
	return s.database.Insert()
}
