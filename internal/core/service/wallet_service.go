package service

import (
	"errors"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/google/uuid"
	"time"
)

type WalletService struct {
	db ports.WalletRepositoryContracts
}
func NewWalletService() *WalletService {
	db := repositories.NewGormDatabase()
	return &WalletService{
		db: db,
	}
}
func (r *WalletService) GenerateWalletService(userID uuid.UUID) error  {
	var generatedWallet entities.Wallet
	generatedWallet.UserID = userID
	generatedWallet.CreatedAt = time.Now()
	generatedWallet.Credit = 0
	err := r.db.GenerateWalletRepository(generatedWallet)
	if err != nil {
		return err
	}
	return nil
}
func (r *WalletService) ChargeWalletService(walletID, userID uuid.UUID , amount int) error  {

	 err := r.db.ChargeWalletRepository(walletID , userID , amount)
	return err
}
func (r *WalletService) GetWalletAmountService(walletID, userID uuid.UUID) (int, error) {
	wallet , err := r.db.GetWalletAmountRepository(walletID)
	if err != nil {
		return 0 , err
	}
	if wallet.UserID != userID {
		err := errors.New("userId is not match with this walletId")
		return 0, err
	}
	return wallet.Credit , nil
}
func (r *WalletService ) HasEnoughCreditService(walletID uuid.UUID, amount int) (bool, error)  {
	wallet , err := r.db.HasEnoughCreditRepository(walletID)
	if err != nil {
		return false, err
	}
	if wallet.Credit >= amount {
		return true , nil
	}
	return false , nil
}