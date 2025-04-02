package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 computes the MD5 hash of the given data and returns it as a hexadecimal string
func MD5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
