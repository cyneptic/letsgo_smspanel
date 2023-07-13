package ports

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

type WalletRepositoryContracts interface {
	GenerateWallet(wallet entities.Wallet) error
	ChargeWallet(walletID, userID uuid.UUID, amount int) error
	GetWallet(walletID uuid.UUID) (*entities.Wallet, error)
	WithdrawFromWallet(userid uuid.UUID, amount int) error
}
