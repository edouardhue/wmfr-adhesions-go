package iraiser

import (
	"crypto/md5"
	"bytes"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
)

type SecureHeader struct {
	Login     string
	Timestamp string
	Token     []byte
}

func Verify(h *SecureHeader) bool {
	var expected = md5.Sum([]byte(h.Login + internal.Config.IRaiser.SecureKey + h.Timestamp))
	return bytes.Equal(expected[:], h.Token)
}