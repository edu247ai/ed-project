package mocks

type MockCID struct{}

func (m *MockCID) CreateFromBytes(data []byte) string {
	return "mockcid123"
}

func (m *MockCID) IsValid(cid string) bool {
	return cid == "mockcid123"
}
