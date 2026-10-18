package preprocessing

type Preprocessor struct{}

func New() *Preprocessor {
	return &Preprocessor{}
}

func (p *Preprocessor) Clean(data []byte) []byte {
	return data
}
