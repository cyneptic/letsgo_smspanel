package ports

import "github.com/google/uuid"

type WalletServiceContracts interface {
	GenerateWalletService(userID uuid.UUID) error
	ChargeWalletService(walletID, userID uuid.UUID, amount int) error
	GetWalletAmountService(walletID, userID uuid.UUID) (int, error)
	HasEnoughCreditService(walletID uuid.UUID, amount int) (bool, error)
}