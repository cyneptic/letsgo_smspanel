package validators

import "errors"

func GenerateWalletValidation(id string) error {
	if id == "" {
		return errors.New("enter valid userID")
	}
	return nil
}
func ChargeWalletValidation(userID, walletID, amount string, amountNum int) error {
	if userID == "" {
		return errors.New("enter valid userID")
	}
	if walletID == "" {
		return errors.New("enter valid walletID")
	}
	if amount == "" {
		return errors.New("enter valid amount number")
	}
	if amountNum < 0 {
		return errors.New("amount number should not be negative")
	}
	return nil
}
func GetWalletAmountValidation(userID, walletID string) error {
	if userID == "" {
		return errors.New("please add valid userID")
	}
	if walletID == "" {
		return errors.New("please enter valid walletID")
	}
	return nil
}
func HasWalletCredit(walletID, amount string) error {
	if walletID == "" {
		return errors.New("please enter valid walletID")
	}
	if amount == "" {
		return errors.New("please enter valid walletID")
	}
	return nil
}
