package iraiser

import (
	"crypto/md5"
	"bytes"
)

type SecureHeader struct {
	Login     string
	Timestamp string
	Token     []byte
}

func (i *IRaiser) Verify(h *SecureHeader) bool {
	var expected = md5.Sum([]byte(h.Login + i.config.SecureKey + h.Timestamp))
	return bytes.Equal(expected[:], h.Token)
}