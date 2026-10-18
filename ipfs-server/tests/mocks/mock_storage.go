package mocks

import "fmt"

type MockStorage struct {
	Data map[string][]byte
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		Data: make(map[string][]byte),
	}
}

func (m *MockStorage) Put(cid string, data []byte) error {
	m.Data[cid] = data
	return nil
}

func (m *MockStorage) Get(cid string) ([]byte, error) {
	v, ok := m.Data[cid]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return v, nil
}

func (m *MockStorage) Delete(cid string) error {
	delete(m.Data, cid)
	return nil
}
