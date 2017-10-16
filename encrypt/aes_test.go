package encrypt

import (
	"encoding/base64"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	key := "test1234test1234"
	content := "this is content"
	ciphertext, err := AESEncrypt([]byte(content), []byte(key))
	if err != nil {
		t.Errorf("failed to encrypt, %v", err)
		return
	}
	t.Logf("ciphertext: %s, len: %d\n", base64.StdEncoding.EncodeToString(ciphertext), len(ciphertext))
	plain, err := AESDecrypt(ciphertext, []byte(key))
	if err != nil {
		t.Errorf("failed to decrypt, %v", err)
		return
	}
	t.Logf("plain: %s, len: %d\n", string(plain), len(plain))
}
