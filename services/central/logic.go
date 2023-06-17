package central

type CentralService struct {
}

func NewCentralService() *CentralService {
	return &CentralService{}
}

func (s *CentralService) AddNumber(request AddNumberRequest) (response AddNumberResponse, err error) {
	response = AddNumberResponse{
		Result: request.First + request.Second,
	}
	return response, nil
}
