package central

import (
	"strconv"

	"github.com/pwn233/golang-short-course/common"
)

type CentralService struct {
}

func NewCentralService() *CentralService {
	return &CentralService{}
}

func (s *CentralService) AddNumber(request AddNumberRequest) (response AddNumberResponse, err error) {
	if err = common.CheckEmptyInput(strconv.Itoa(request.First)); err != nil {
		return AddNumberResponse{}, err
	}
	if err = common.CheckEmptyInput(strconv.Itoa(request.Second)); err != nil {
		return AddNumberResponse{}, err
	}
	return AddNumberResponse{
		Result: request.First + request.Second,
	}, nil
}
