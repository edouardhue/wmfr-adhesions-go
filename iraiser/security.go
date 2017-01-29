package iraiser

import (
	"crypto/md5"
	"bytes"
)

func (i *IRaiser) Verify(h *SecureHeader) bool {
	var expected = md5.Sum([]byte(h.Login + i.config.SecureKey + h.Timestamp))
	return bytes.Equal(expected[:], h.Token)
}