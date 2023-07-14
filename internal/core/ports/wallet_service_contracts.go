package ports

import "github.com/google/uuid"

type WalletServiceContracts interface {
	GenerateWalletService(userID uuid.UUID) error
	ChargeWalletService(walletID, userID uuid.UUID, amount int) (redirectLink string, err error)
	ChargeWalletServiceVerify(refId string, walletID, userID uuid.UUID, SaleReferenceId string) (status bool, err error)
	GetWalletAmountService(walletID, userID uuid.UUID) (int, error)
	HasEnoughCreditService(walletID uuid.UUID, amount int) (bool, error)
}
