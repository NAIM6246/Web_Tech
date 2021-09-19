package services

import "fmt"

type UserService struct{}

func (h *UserService) Hello() string {
	return fmt.Sprintf("from service layer")
}
