package rule

import (
	"regexp"
)

//Find char [Space]
func IsSpaceExist(s string) bool {
	re := regexp.MustCompile(`[\x20]+`)
	return re.MatchString(s)
}

//All letter inside ASCII 21 - 7E
func IsAsciiChar(s string) bool {
	re := regexp.MustCompile(`^[\x21-\x7E]+$`)
	return re.MatchString(s)
}

//All letter numeric
func IsNumeric(s string) bool {
	re := regexp.MustCompile(`^[\d]+$`)
	return re.MatchString(s)
}

//All letter Alphabet
func IsAlphaOnly(s string) bool {
	re := regexp.MustCompile(`^[a-zA-z]+$`)
	return re.MatchString(s)
}

//Find char Uppercase
func IsUppercaseLetterExist(s string) bool {
	re := regexp.MustCompile(`[A-Z]+`)
	return re.MatchString(s)
}

//Find char lowercase
func IsLowercaseLetterExist(s string) bool {
	re := regexp.MustCompile(`[a-z]+`)
	return re.MatchString(s)
}

//All letter combination alphabet & numeric
func IsAlphaNumeric(s string) bool {
	re1 := regexp.MustCompile(`[a-zA-Z]+`)
	re2 := regexp.MustCompile(`[\d]+`)
	if re1.MatchString(s) && re2.MatchString(s) {
		return true
	}
	return false
}

//All letter combination 1 uppercase, 1 lowercase, 1 numeric
func IsValidFormat(s string) bool {
	re1 := regexp.MustCompile(`[a-z]+`)
	re2 := regexp.MustCompile(`[A-Z]+`)
	re3 := regexp.MustCompile(`[\d]+`)
	if re1.MatchString(s) && re2.MatchString(s) && re3.MatchString(s) {
		return true
	}
	return false
}

//Forbidden word
func IsContainForbbidenWord(s string) bool {
	wordList := []string{
		"Bri12345",
		"Bismillah",
		"P@ssw0rd",
		"Ps12345678",
		"Jakarta123",
		"Jakarta1",
		"Bri123456",
		"Password1",
		"Bismillah123",
		"Bri12345678",
		"Qwerty123",
		"Abc12345",
		"BRI1234567",
		"Abcd1234",
		"Indonesia1",
		"123456",
		"123456789",
		"qwerty",
		"password",
		"1234567",
		"12345678",
		"12345",
		"iloveyou",
		"111111",
		"123123",
		"abc123",
		"qwerty123",
		"1q2w3e4r",
		"admin",
		"qwertyuiop",
		"654321",
		"555555",
		"lovely",
		"7777777",
		"welcome",
		"888888",
		"princess",
		"dragon",
		"password1",
		"123qwe",
	}

	for _, word := range wordList {
		if s == word {
			return true
		}
	}
	return false
}
