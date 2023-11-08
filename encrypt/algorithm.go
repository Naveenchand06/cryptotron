package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		encryptChar := string(asciiCode + 3)
		encryptedStr += encryptChar
	}
	return encryptedStr
}
