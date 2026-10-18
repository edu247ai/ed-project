package unit

import (
	"bytes"
	"testing"

	"ipfs-server/cid/generator"
	"ipfs-server/cid/validator"
)

func TestCIDGeneration(t *testing.T) {
	gen := generator.New()
	cid := gen.FromBytes([]byte("hello"))

	if !validator.IsValid(cid) {
		t.Fatalf("expected valid CID, got %s", cid)
	}
	const expected = "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq"
	if cid != expected {
		t.Fatalf("expected CID %s, got %s", expected, cid)
	}
}

func TestCIDRejectsNonCIDAndNonSHA256Multihashes(t *testing.T) {
	for _, value := range []string{
		"",
		"bafkreia",
		"0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		"BAFKREIBM6JG3UX5QUMHCN2B3FLC3TYU6D6WFEZQ6V4J6HKN5H7G7IQ",
	} {
		if validator.IsValid(value) {
			t.Errorf("expected invalid CID %q", value)
		}
	}
}

func TestCIDFromReaderMatchesBytes(t *testing.T) {
	gen := generator.New()
	data := []byte("streamed content")

	got, err := gen.FromReader(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if want := gen.FromBytes(data); got != want {
		t.Fatalf("reader CID %s does not match byte CID %s", got, want)
	}
}
