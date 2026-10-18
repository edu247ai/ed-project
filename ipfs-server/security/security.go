package security

import (
	accesscontrol "ipfs-server/security/access-control"
	"ipfs-server/security/audit"
	"ipfs-server/security/encryption"
	"ipfs-server/security/signatures"
)

type Security struct {
	Access    *accesscontrol.AccessControl
	Audit     *audit.Auditor
	Encryptor *encryption.Encryptor
	Signer    *signatures.Signer
}

func New(encKey, sigKey []byte) (*Security, error) {
	enc, err := encryption.New(encKey)
	if err != nil {
		return nil, err
	}

	return &Security{
		Access:    accesscontrol.New(),
		Audit:     audit.New(),
		Encryptor: enc,
		Signer:    signatures.New(sigKey),
	}, nil
}
