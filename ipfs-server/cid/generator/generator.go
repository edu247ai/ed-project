package generator

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"io"
	"strings"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) FromBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return encode(sum[:])
}

func (g *Generator) FromReader(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return encode(h.Sum(nil)), nil
}

func encode(digest []byte) string {
	cidBytes := binary.AppendUvarint(nil, 1)
	cidBytes = binary.AppendUvarint(cidBytes, 0x55)
	cidBytes = binary.AppendUvarint(cidBytes, 0x12)
	cidBytes = binary.AppendUvarint(cidBytes, uint64(len(digest)))
	cidBytes = append(cidBytes, digest...)

	return "b" + strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(cidBytes))
}
