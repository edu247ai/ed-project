package postprocessing

import (
	"bytes"
	"compress/gzip"
)

type Postprocessor struct{}

func New() *Postprocessor {
	return &Postprocessor{}
}

func (p *Postprocessor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
