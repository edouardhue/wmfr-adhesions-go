package iraiser

import "testing"

func TestBadTokenIsRejected(t *testing.T) {
	iRaiser := NewIRaiser(&Config{
		SecureKey: "Key",
	})
	var h = SecureHeader{
		Login: "Test",
		Timestamp: "1483109511",
		Token: []byte{0x10, 0x10},
	}
	if iRaiser.Verify(&h) {
		t.Errorf("Token % x should be rejected.", h.Token)
	}
}

func TestGoodTokenIsApproved(t *testing.T) {
	iRaiser := NewIRaiser(&Config{
		SecureKey: "Key",
	})
	var h = SecureHeader{
		Login: "Test",
		Timestamp: "1483109511",
		Token: []byte{0xBD, 0x14, 0x07, 0x2C, 0x4F, 0x2F, 0xF2, 0x16, 0x79, 0xCD, 0xFC, 0xEA, 0xE2, 0x6E, 0x72, 0x0F},
	}
	if !iRaiser.Verify(&h) {
		t.Errorf("Token % x should be approved.", h.Token)
	}
}