package validator

import (
	"encoding/base32"
	"encoding/binary"
	"strings"
)

func IsValid(cid string) bool {
	if len(cid) < 2 || cid[0] != 'b' || cid != strings.ToLower(cid) {
		return false
	}

	decoded, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(strings.ToUpper(cid[1:]))
	if err != nil {
		return false
	}
	if "b"+strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(decoded)) != cid {
		return false
	}

	version, n := binary.Uvarint(decoded)
	if n <= 0 || version != 1 || !isCanonicalVarint(decoded[:n], version) {
		return false
	}
	codec, consumed := binary.Uvarint(decoded[n:])
	if consumed <= 0 || codec != 0x55 || !isCanonicalVarint(decoded[n:n+consumed], codec) {
		return false
	}
	offset := n + consumed

	hashCode, consumed := binary.Uvarint(decoded[offset:])
	if consumed <= 0 || hashCode != 0x12 || !isCanonicalVarint(decoded[offset:offset+consumed], hashCode) {
		return false
	}
	offset += consumed

	hashLength, consumed := binary.Uvarint(decoded[offset:])
	if consumed <= 0 || hashLength != 32 || !isCanonicalVarint(decoded[offset:offset+consumed], hashLength) {
		return false
	}
	offset += consumed

	return len(decoded)-offset == int(hashLength)
}

func isCanonicalVarint(encoded []byte, value uint64) bool {
	canonical := binary.AppendUvarint(nil, value)
	if len(encoded) != len(canonical) {
		return false
	}
	for i := range canonical {
		if encoded[i] != canonical[i] {
			return false
		}
	}
	return true
}
