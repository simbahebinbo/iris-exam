package repositories

import "github.com/hqd888/iris-example/datamodels"

type DemoRepository interface {
	GetDemoName() string
}

type DemoManager struct {
}

func NewDemoManager() DemoRepository {
	return &DemoManager{}
}

func (m *DemoManager) GetDemoName() string {
	d := &datamodels.Demo{Name: "this is a test!"}
	return d.Name
}
