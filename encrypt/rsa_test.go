package encrypt

import (
	"encoding/base64"
	"testing"
)

var (
	_pubKeyPKCS1     = []byte("-----BEGIN PUBLIC KEY-----\n	MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCU9H/6QkUu6AqY7EnKpxAXSanp\np575/9+lcvTMrtgaxqpfkCYv15Qq2oMDGA7ggrGYrLlSs9oW2Klz3vZ3iP0leZ5L\nlmLlc52Wqdp9eeNCSemr9dVIaMOh6J5CR6TkTb/yf64BhVGcBRCgca+Msdx0pmwl\nii+bp0mPFAe5r2bwqQIDAQAB\n-----END PUBLIC KEY-----")
	_privateKeyPKCS1 = []byte("-----BEGIN RSA PRIVATE KEY-----\n	MIICXQIBAAKBgQCU9H/6QkUu6AqY7EnKpxAXSanpp575/9+lcvTMrtgaxqpfkCYv\n15Qq2oMDGA7ggrGYrLlSs9oW2Klz3vZ3iP0leZ5LlmLlc52Wqdp9eeNCSemr9dVI\naMOh6J5CR6TkTb/yf64BhVGcBRCgca+Msdx0pmwlii+bp0mPFAe5r2bwqQIDAQAB\nAoGABYhhr2e5CiK0VFnN7QSC/+y6QKyyXUkmFXzfn3qxYxZeAkWNAP5+pLRG0qaY\n+ScGLSfh1cUmJKVjAVrNo7C5yKd6zntoi9LsIp7OaT4Gvf+wwlRMmf9ybf0BwtqZ\no+pVcN2iqjeFkTRGwdpiAMuD+10Z1ugvzTCm38eANNZf0AECQQC9sFvmmb0b/3iH\nHpfl8YJG3zyYIQXOw9TQJijxuQ4HCO887AFtK4DUWGVRZNOvOnXgSNF+g2bSOZXI\nZ4pXc8xJAkEAyQbKA+Y0Hc/F+BQ6zh3qml4dGzuHw7gy1W5CGfSlivWI05LIN6kn\nnbvpwmFYKNS2TprarudBCzNDgIKF6GhBYQJBALEAuiDKYL6ZevyKva8h1zEO6loC\npoq0DgdVNGurBL3nxmYQy5v1NjfvbEuxTB587LWm/WWSdAXDGlNw2pSUR5ECQQCf\nl6k3kUGczT+wpHbO/gcrSD45SXiOFCUfNip6KrRl7moS9bvIHJ2b5vw82kVAHSci\nnmOwaV1FBAy58GVYzbPhAkBeda2E7/4eyUj4lAEvZCqUe+DX+mUoTQzgG8kZ6RBt\nyWAT+UBgAhy/2yyE0vMzMHYTDS+hRnn9py7WLmyYbnZb\n-----END RSA PRIVATE KEY-----")
	_pubKeyPKCS8     = []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCeWWZD7huChGZ1bg4INW4+aJ/+\nJJ9tGue/wG+8A7wYoJEduHBCWuBnrr3JYhc4JDO6I+N8IhWyKcjYfrDSt0pRJ03C\nfDH0xOrN50Q1PiR25pMplhxXC0GnsOQ/t4chsHuIOpRX1l2F0u2IBcKQPPh3AF6i\nlJmkHx7wj1qeo2r+HwIDAQAB\n-----END PUBLIC KEY-----")
	_privateKeyPKCS8 = []byte("-----BEGIN PRIVATE KEY-----\nMIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJ5ZZkPuG4KEZnVu\nDgg1bj5on/4kn20a57/Ab7wDvBigkR24cEJa4GeuvcliFzgkM7oj43wiFbIpyNh+\nsNK3SlEnTcJ8MfTE6s3nRDU+JHbmkymWHFcLQaew5D+3hyGwe4g6lFfWXYXS7YgF\nwpA8+HcAXqKUmaQfHvCPWp6jav4fAgMBAAECgYAT+htYvpVh5eo3Lmc+yd7aqMBm\ng6BJV3edTV3LTn1QbhFguoAPHv/olLcEzQc44fhUp2pvoq0yCOt5xKHyM4WAfKpF\np8eitq3yMCvTRgfYhHNNr55YClUoDEqXCu4R6khIbaPPE2gJkgHb8jWl3VEzhtsV\nmSiIGqBgzAe6vc7ogQJBANEisUkq+AcUHVuDC5yv+Nx2bOsivm+aAgSy0C4rxGMQ\n2XftgSd2FURM+9xH2ENmpLH6zJbpbeNBbiz73w3yCNECQQDB1Upa3goLuHFwsVOu\nu//Tf+AfZNyUzla8FpSlEKf5QPwbkcMD5K9JScqUr1pgJaxn8kAfrnBYlXIlqrUC\nSVPvAkBIz7xLjhM2W7Fw5+oGDlolX3HCV6FXt5XWTu8d9Az2tKoSD/V0aK80zVyu\nr7DCnpzefkP2kqS4h1I1hSd+cnLRAkB5CZNwwo5LW1WbJWA4ELVjgMqXUAhd86s9\nsGwJ1yjNAMNtA8xfNgIvJaEWz5kDyQKrth5MqkUFS+0HkF2Pm/KJAkA6h2Aumjjg\nW2kNQxnu5yh3aDd1ADRFIAlUEetIA7kBmOhqSik0Qf9cK03WQvH2hKGFh6ebcyBi\n+TlajOm7itPf\n-----END PRIVATE KEY-----")
)

func TestRSAEncryptPKCS1(t *testing.T) {
	plain := []byte("hello")
	d, err := RSAEncryptPKCS1(_pubKeyPKCS1, plain)
	if err != nil {
		t.Errorf("failed to rsa encrypt, error(%v)", err)
		t.FailNow()
	}
	t.Logf("d: %s", base64.StdEncoding.EncodeToString(d))
	p, err := RSADecryptPKCS1(_privateKeyPKCS1, d)
	if err != nil {
		t.Errorf("failed to rsa decrypt, error(%v)", err)
		t.FailNow()
	}
	t.Logf("p: %s", p)
}

func TestRSAEncryptPKCS8(t *testing.T) {
	plain := []byte("hello")
	d, err := RSAEncryptPKCS8(_pubKeyPKCS8, plain)
	if err != nil {
		t.Errorf("failed to rsa encrypt, error(%v)", err)
		t.FailNow()
	}
	t.Logf("d: %s", base64.StdEncoding.EncodeToString(d))
	p, err := RSADecryptPKCS8(_privateKeyPKCS8, d)
	if err != nil {
		t.Errorf("failed to rsa decrypt, error(%v)", err)
		t.FailNow()
	}
	t.Logf("p: %s", p)
}
