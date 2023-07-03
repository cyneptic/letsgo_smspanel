package ports

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

type WalletRepositoryContracts interface {
	GenerateWalletRepository(wallet entities.Wallet) error
	ChargeWalletRepository(walletID, userID uuid.UUID , amount int)  error
	GetWalletAmountRepository(walletID  uuid.UUID) (*entities.Wallet, error)
	HasEnoughCreditRepository(walletID uuid.UUID) (*entities.Wallet, error)
}
