package services

import (
	"fmt"

	"github.com/hqd888/demo_iris/repositories"
)

type DemoService interface {
	ShowDemoName() string
}

type DemoServiceManager struct {
	repo repositories.DemoRepository // interface
}

func NewDemoServiceManager(repo repositories.DemoRepository) DemoService {
	return &DemoServiceManager{repo: repo}
}

func (m *DemoServiceManager) ShowDemoName() string {
	fmt.Println(m.repo.GetDemoName())
	return m.repo.GetDemoName()
}
