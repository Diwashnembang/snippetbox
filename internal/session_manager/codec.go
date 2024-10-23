package sessionmanager

import "time"

// Codec is the interface for encoding/decoding session data to and from a byte
// slice for use by the session store
type Codec interface {
	Encode(deadline time.Time, values map[string]interface{}) ([]byte, error)
	Decode([]byte) (deadline time.Time, values map[string]interface{}, err error)
}
