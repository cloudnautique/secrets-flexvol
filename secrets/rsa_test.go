package secrets

import (
	"encoding/base64"
	"testing"
)

var insecureKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAlqXgL8UtupeafCFVQwckREfGN+KM3M+tiY0CLsd847w3B3MI
rwurSDvBRZMvriYz7LCQIrrXTri8XZC0LNvRdkkHr9HWNPwA1eB8DLRORPIp0H4I
9XwLHP76qaKJY2Af2vL8Oq0paSiSwtCaN983JNwyDXmgGKYv0K+6byUv6AVtiQS8
8kOylCnrSKkui7nzcFuoLR/RwuLCxoK9jmAGBNJCG/16u9eFnaElJ1kCcnS0XsdJ
Biy60lWgnMLwlGel0vGZXjTOdAF1xMHZAHSq2Y0k82brNvxLNQSdnV1TjU70rSYO
Li/hoNep978UR76Fv2ZdBY8Ft06N09N4JNanWwIDAQABAoIBAFDlEdWVFFE2R4aQ
f7BWjWr8/7vSs8F+47kRNzLXfIDt+L7PTsJwibFoJQivWNMzQH7A8SU1H5juKngz
1AyinX/fB3mqPFSHXgt7WCGaUM1FHJ8Qjs8DpRQU95VP6maqn3B7OmZnxezqFKT4
T1fhTUNF2rrRrN6Pnu1476vvVCJKtPJcAqG4IIE01jrvZ/jD1wiZ+s3fpJN0Q/j3
FEkWP0B+KPAbE9viEK+aKX0eO2Jkq7xZYgslQRV1TrCooQ5U2+/xBypGrggHloK/
5/apjteJxwljyZMBRFXoX3Yl6Y2y/TXg2fYTTKo323IVLx/080REYjOXcGujp5Sy
cXJ7SsECgYEAxrzXmfO9E718bjilUBT1t2fy2gch+tubDsQeMwXD57sIgSE4Sr7k
xkaHW6FfgA0rtj94CkMW00509ny7HkyaFNkwrkrC/0R/gUIo0E31fgxTM2cO3urI
QXFw1lmFVsE9/uppgF5L9ktSe8TJz7fMp8iHV+1N7FDyuoNSoFp6/bcCgYEAwg3f
Hni3I5JgRI6MX5j1HquUt76PqI7CYeqRmqcHBSg6d5u1Y0P2Fulh4gdYIX8QrGi2
5viSaTZQt9DVATF4pKs2XMPZc9QooudYTSUhRDAnRfdYFa0E56rtL2L/RXTbZj7S
jYdmMrMBvB9mY+RbLTeWK7yG53IzaidJVp6tY30CgYBo8zbkPRwffZRlXJKoTLlK
BqHv0451PF2RGa5dAXFoQZQHJTTl/BMyRfKbSAf3xnzL/I521OEL68XGmS3znT5N
PjkAAckiJtkyuG53OoQm8XlKjuUCgXgJX0/YUmQg4WHM6ZuXR7TTtwkzBUQR5p00
Cai3nUDmSAU2y7zpo36J1wKBgEZtVGGxu/27/RZEieuUDroP2YyKK4coMKHqyOdQ
4Tpc7ENGjqE1JBYSo4St161oeTupUWAoLLLklIzxzKx/MOLKhJNMPRpNkGX3AlQV
OqqNs2MwLpbHUXVm0mgVTMH/dDT6bd4RmuShlOqalsWANhsGBolfBbLv/nrzQSmf
sxvdAoGALwb3fP9ir2Fs3vHn4fCpiuNoCovWExbhH+UtQ/kDYuXsjt1Th7cxuLPF
FNH/hPpMSf5p6Gl4Ipl12s5U6FVYQlmuVlFgV8iUEKsSkMWdrvvx5X38RlgqQqvU
+7k/Qphbh1dQWKCpMXmeMxRWTtgaftz18zvou6k0CyCSNco6JZ4=
-----END RSA PRIVATE KEY-----`

var testMessage = "Yyaba6uZYkPHLqzzh4n6SB76tU32ugonB8uxdViUhxKpk/tThhPdQQvj4pe1k3advNOMUyIuykbnJ9EUVY4M4KRdilt6KlCQTEPrzTGw9ZxoFdBWlW2Kj3+1BZt/iy36krzryyLS+bNIDE8IRNoafaPmcto1ywQHfBjXiIjoJfYIuXpbQPOLU1ulElMv7ArwG2JbIvYcpIMysoJqaJ7YAauHveMPmAbRB/oGgS/pxIoP9vv1PMPIoP6c6h4raWXZ6uRkMJ7ND6cEq3pXLVlapYgZnOV9lbMBxQGlzApVlDo4BnMsNz/NNiaKYQs5CjO12KySuDjLkRamERL1FaKQhA=="
var expectedMessage = "hello"

func TestDecrypt(t *testing.T) {
	rsaKey, err := loadPrivateKeyFromString(insecureKey)
	if err != nil {
		t.Errorf("Failed to parse private key: %s", err)
		return
	}

	data, err := rsaDecrypt(rsaKey, testMessage)
	if err != nil {
		t.Error(err)
		return
	}

	val, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		t.Error(err)
		return
	}

	if string(val) != expectedMessage {
		t.Errorf("vaules were supposed to match: %s != %s", string(val), expectedMessage)
		return
	}
}
