package module

import (
	"github.com/ndokoblog/complianceBrimo/model"
	"github.com/ndokoblog/complianceBrimo/rule"
)

func ComplyPin(username string, usernameAlias string, pin string) model.Validation {
	var validationObject model.Validation
	//pin length must min 6 , max 6 if pin true
	if len(pin) != 6 {
		validationObject = model.Validation{
			Code: "HA",
			Desc: "Panjang minimal & maksimal pin adalah 6",
		}
		return validationObject
	}

	//pin must contain only numeric
	if !rule.IsNumeric(pin) {
		validationObject = model.Validation{
			Code: "HB",
			Desc: "Pin harus numerik",
		}
		return validationObject
	}

	switch pin {
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
			Code: "HC",
			Desc: "Pin terlalu mudah",
		}
		return validationObject
	}

	validationObject = model.Validation{
		Code: "00",
		Desc: "Format pin valid",
	}
	return validationObject
}
