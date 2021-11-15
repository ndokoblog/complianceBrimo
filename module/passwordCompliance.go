package module

import (
	"strings"

	"github.com/ndokoblog/complianceBrimo/model"
	"github.com/ndokoblog/complianceBrimo/rule"
)

func ComplyPassword(username string, usernameAlias string, password string, channel string) model.Validation {
	var validationObject model.Validation
	isUsernameAliasExist := false
	if usernameAlias != "" {
		isUsernameAliasExist = true
	}

	isPinPassword := false
	if channel == "EDC" || channel == "ATM" {
		isPinPassword = true
	}

	if isPinPassword {
		if len(password) != 6 {
			validationObject = model.Validation{
				Code: "V6",
				Desc: "Panjang minimal & Maksimal pin adalah 6",
			}
			return validationObject
		}

		//Password must contain only numeric
		if !rule.IsNumeric(password) {
			validationObject = model.Validation{
				Code: "VH",
				Desc: "Pin harus angka",
			}
			return validationObject
		}

		switch password {
		case
			"012345",
			"123456",
			"234567",
			"345678",
			"456789",
			"567890",
			"654321",
			"765432",
			"876543",
			"987654",
			"098765",
			"112233",
			"332211",
			"111111",
			"222222",
			"333333",
			"444444",
			"555555",
			"666666",
			"777777",
			"888888",
			"999999",
			"000000":
			validationObject = model.Validation{
				Code: "VI",
				Desc: "Pin terlalu mudah",
			}
			return validationObject
		}

	}

	if !isPinPassword {
		//Password length must min 8
		if len(password) < 8 {
			validationObject = model.Validation{
				Code: "V6",
				Desc: "Panjang minimal password adalah 8",
			}
			return validationObject
		}

		//Password length must max 12
		if len(password) > 12 {
			validationObject = model.Validation{
				Code: "V7",
				Desc: "Panjang maximal password adalah 12",
			}
			return validationObject
		}

		//Password must not contain space
		if rule.IsSpaceExist(password) {
			validationObject = model.Validation{
				Code: "V8",
				Desc: "Password tidak boleh mengandung karakter spasi",
			}
			return validationObject
		}

		//Password must contain alpha and numeric
		if !rule.IsAsciiChar(password) {
			validationObject = model.Validation{
				Code: "V9",
				Desc: "Password harus mengandung karakter dalam lingkup ASCII x21-x7E",
			}
			return validationObject
		}

		//Password must not contain only numeric
		if rule.IsNumeric(password) {
			validationObject = model.Validation{
				Code: "VA",
				Desc: "Password tidak diperbolehkan hanya angka saja, harus mengandung minimal 1 huruf besar, 1 huruf kecil, 1 angka",
			}
			return validationObject
		}

		//Password must not contain only alpha
		if rule.IsAlphaOnly(password) {
			validationObject = model.Validation{
				Code: "VB",
				Desc: "Password tidak diperbolehkan hanya alfabet saja, harus mengandung minimal 1 huruf besar, 1 huruf kecil, 1 angka",
			}
			return validationObject
		}

		//Password must contain uppercase char
		if !rule.IsUppercaseLetterExist(password) {
			validationObject = model.Validation{
				Code: "VE",
				Desc: "Password harus mengandung minimal 1 huruf besar",
			}
			return validationObject
		}

		//Password must contain uppercase char
		if !rule.IsLowercaseLetterExist(password) {
			validationObject = model.Validation{
				Code: "VJ",
				Desc: "Password harus mengandung minimal 1 huruf kecil",
			}
			return validationObject
		}

		//Password must not contain only alpha
		if !rule.IsValidFormat(password) {
			validationObject = model.Validation{
				Code: "VC",
				Desc: "Format password harus mengandung minimal 1 huruf besar, 1 huruf kecil, 1 angka",
			}
			return validationObject
		}

		//Password must not equal with username
		if strings.Contains(strings.ToUpper(password), strings.ToUpper(username)) {
			validationObject = model.Validation{
				Code: "VF",
				Desc: "Password tidak boleh mengandung unsur username",
			}
			return validationObject
		}

		//Password must not equal with username alias
		if isUsernameAliasExist {
			if strings.Contains(strings.ToUpper(password), strings.ToUpper(usernameAlias)) {
				validationObject = model.Validation{
					Code: "VG",
					Desc: "Password tidak boleh mengandung unsur user alias",
				}
				return validationObject
			}
		}

		//Password must not contain forbidden word
		if rule.IsContainForbbidenWord(password) {
			validationObject = model.Validation{
				Code: "VD",
				Desc: "Password tidak boleh mengandung kata terlarang",
			}
			return validationObject
		}
	}

	validationObject = model.Validation{
		Code: "00",
		Desc: "Format password valid",
	}
	return validationObject
}
