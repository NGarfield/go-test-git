package services

type Servicer interface {
	InsertTestDataService(req *Test) error
}
type Repositorer interface {
	InsertTestData(name, lastname string, age uint) error
}
type Service struct {
	mysql Repositorer
}

func NewService(db Repositorer) *Service {
	return &Service{
		mysql: db,
	}
}

func (s *Service) InsertTestDataService(req *Test) error {
	return s.mysql.InsertTestData(req.Name, req.Lastname, req.Age)
}
