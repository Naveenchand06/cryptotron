package decrypt

func Dimbus(str string ) string {
	decryptedStr := ""
	for _, c := range str {
		asciCode := int(c)
		decryptedStr += string(asciCode - 3)
	}
	return decryptedStr
}