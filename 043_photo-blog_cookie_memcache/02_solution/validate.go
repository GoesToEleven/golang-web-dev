package mem

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func tampered(s string) bool {
	xs := strings.Split(s, "|")
	// usrToken := xs[0]
	usrData := xs[1]
	usrCode := xs[2]
	if usrCode != getCode(usrData) {
		return true
	}
	return false
}
