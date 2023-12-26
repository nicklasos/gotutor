package services

type MenuService struct {
	Name string
}

func NewMenuService(name string) *MenuService {
	return &MenuService{Name: name}
}
