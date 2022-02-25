package services

type DBer interface{
	GetAllTest() ([]Test, error)
}

type Servicer interface{
	GetAllService() ([]Test, error)
}

type Service struct{
	db DBer
}

func NewService(db DBer) *Service{
	return &Service{
		db: db,
	}
}

func (s *Service) GetAllService() ([]Test, error){
	result, err := s.db.GetAllTest()
	if err != nil {
		return nil,err
	}
	return result,nil
}