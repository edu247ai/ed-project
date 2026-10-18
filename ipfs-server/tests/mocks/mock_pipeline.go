package mocks

type MockPipeline struct {
	Store *MockStorage
}

func NewMockPipeline() *MockPipeline {
	return &MockPipeline{
		Store: NewMockStorage(),
	}
}

func (p *MockPipeline) Upload(data []byte) (string, error) {
	cid := "mockcid123"
	p.Store.Put(cid, data)
	return cid, nil
}

func (p *MockPipeline) Download(cid string) ([]byte, error) {
	return p.Store.Get(cid)
}
