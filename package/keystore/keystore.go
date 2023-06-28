package keystore

import (
	"fmt"
)

type KeyStore struct {
	Pairs map[string]string
}

// make Key value Store
func NewValues() KeyStore {

	ks := KeyStore{
		Pairs: make(map[string]string),
	}
	return ks
}

// Set the  values
func (kv KeyStore) SetValue(key string, value string) {
	kv.Pairs[key] = value
}

// Get the values
func (kv KeyStore) GetValue(key string) (string, error) {
	val, err := kv.Pairs[key]
	if !err {
		return val, fmt.Errorf("the key %s does not exist", key)
	}
	return val, nil
}
