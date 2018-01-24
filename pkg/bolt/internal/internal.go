package internal

import (
	"encoding/binary"
	"encoding/json"

	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
)

// MarshalUser encodes a user to binary format.
func MarshalUser(user *cloudnativego.User) ([]byte, error) {
	return json.Marshal(user)
}

// UnmarshalUser decodes a user from a binary data.
func UnmarshalUser(data []byte, user *cloudnativego.User) error {
	return json.Unmarshal(data, user)
}

// Itob returns an 8-byte big endian representation of v.
// This function is typically used for encoding integer IDs to byte slices
// so that they can be used as BoltDB keys.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
