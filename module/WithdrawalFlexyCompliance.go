package module
import (
	//"encoding/json"

	//"github.com/ndokoblog/complianceBrimo/library"
	"github.com/ndokoblog/complianceBrimo/model"
	//"github.com/ndokoblog/complianceBrimo/constant"
)

func ComplyWithdrawalFlexyRequest(
	withdrawalDate string, 
	minBalAfterWithdrawal string, 
	maxPencairanPersen string,
	withdrawalAmount string,
	minPenaltyWithdrawal string,
	percentagePenaltyWithdrawal string,
	availableBal string,
) model.Validation {
	//var flexyParameter map[string]interface{}
	var oke model.Validation
	//err := json.Unmarshal([]byte(parameter), &flexyParameter)

	//if err != nil {
	//	return library.GetValidationResult(constant.RC_J0, constant.DESC_J0)
	//}

	return oke
}