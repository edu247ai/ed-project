package accesscontrol

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"sync"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleNode   Role = "node"
	RoleClient Role = "client"
)

type AccessControl struct {
	mu      sync.RWMutex
	apiKeys map[[sha256.Size]byte]Role
}

func New() *AccessControl {
	return &AccessControl{
		apiKeys: map[[sha256.Size]byte]Role{},
	}
}

func (ac *AccessControl) RegisterKey(key string, role Role) error {
	if key == "" {
		return errors.New("API key cannot be empty")
	}
	switch role {
	case RoleAdmin, RoleNode, RoleClient:
	default:
		return errors.New("invalid API key role")
	}

	digest := sha256.Sum256([]byte(key))
	ac.mu.Lock()
	ac.apiKeys[digest] = role
	ac.mu.Unlock()
	return nil
}

func (ac *AccessControl) Authorize(key string, required Role) error {
	if key == "" {
		return errors.New("invalid API key")
	}

	digest := sha256.Sum256([]byte(key))
	ac.mu.RLock()
	var role Role
	var ok int
	for registeredDigest, registeredRole := range ac.apiKeys {
		match := subtle.ConstantTimeCompare(digest[:], registeredDigest[:])
		if match == 1 {
			role = registeredRole
			ok = 1
		}
	}
	ac.mu.RUnlock()
	if ok != 1 {
		return errors.New("invalid API key")
	}
	if role != required && role != RoleAdmin {
		return errors.New("insufficient permissions")
	}
	return nil
}

func (ac *AccessControl) RoleFor(key string) (Role, bool) {
	if key == "" {
		return "", false
	}
	digest := sha256.Sum256([]byte(key))
	ac.mu.RLock()
	var role Role
	var ok int
	for registeredDigest, registeredRole := range ac.apiKeys {
		match := subtle.ConstantTimeCompare(digest[:], registeredDigest[:])
		if match == 1 {
			role = registeredRole
			ok = 1
		}
	}
	ac.mu.RUnlock()
	return role, ok == 1
}
