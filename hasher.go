package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)
func Shahasher(userString string) string{
	hash_:=sha256.Sum256([]byte(userString))
	value:=hex.EncodeToString(hash_[:])
	return value
}

func Md5hasher(userString string) string{
	hash_:=md5.Sum([]byte(userString))
	value:=hex.EncodeToString(hash_[:])
	return value
}