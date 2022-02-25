package services

type DBer interface {
	GetAllTest() ([]Test, error)
	InsertTestData(name, lastname string, age uint) error
}

type Servicer interface {
	GetAllService() ([]Test, error)
	InsertTestDataService(req *Test) error
}

type Service struct {
	db DBer
}

func NewService(db DBer) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetAllService() ([]Test, error) {
	result, err := s.db.GetAllTest()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) InsertTestDataService(req *Test) error {
	return s.db.InsertTestData(req.Name, req.Lastname, req.Age)
}
