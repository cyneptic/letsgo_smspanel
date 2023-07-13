package service

import (
	"errors"
	"github.com/cyneptic/letsgo-smspanel/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type WalletService struct {
	db       ports.WalletRepositoryContracts
	memoryDB ports.PaymentDbContract
	gateway  ports.PaymentGateWayContract
}

func NewWalletService() *WalletService {
	db := repositories.NewGormDatabase()
	redisDB := repositories.RedisInit()
	gateway := provider.NewMellatGateway()
	return &WalletService{
		db:       db,
		memoryDB: redisDB,
		gateway:  gateway,
	}
}
func (r *WalletService) GenerateWalletService(userID uuid.UUID) error {
	var generatedWallet entities.Wallet
	generatedWallet.UserID = userID
	generatedWallet.CreatedAt = time.Now()
	generatedWallet.Credit = 0
	err := r.db.GenerateWallet(generatedWallet)
	if err != nil {
		return err
	}
	return nil
}
func (r *WalletService) ChargeWalletService(walletID, userID uuid.UUID, amount int) (redirectLink string, err error) {
	redirectLink, refId, err := r.gateway.CreatePayment(string(amount), walletID, userID.String())
	go r.memoryDB.SetPaymentRequest(walletID, userID.String(), refId, string(amount))
	return redirectLink, err
}

func (r *WalletService) ChargeWalletServiceVerify(refId string, walletID, userID uuid.UUID, SaleReferenceId string) (status bool, err error) {
	amount, ok, err := r.memoryDB.VerifyPaymentRequest(userID.String(), refId, walletID.String())
	if !ok {
		return false, errors.New("there is no pending payment")
	}
	result, err := r.gateway.VerifyPayment(userID.String(), refId, walletID.String(), SaleReferenceId)
	if !result {
		return false, errors.New("error in verify payment request")
	}
	IntegerAmount, _ := strconv.ParseInt(amount, 10, 64)
	err = r.db.ChargeWallet(walletID, userID, int(IntegerAmount))
	return false, err
}

func (r *WalletService) GetWalletAmountService(walletID, userID uuid.UUID) (int, error) {
	wallet, err := r.db.GetWallet(walletID)
	if err != nil {
		return 0, err
	}
	if wallet.UserID != userID {
		err := errors.New("userId is not match with this walletId")
		return 0, err
	}
	return wallet.Credit, nil
}
func (r *WalletService) HasEnoughCreditService(walletID uuid.UUID, amount int) (bool, error) {
	wallet, err := r.db.GetWallet(walletID)
	if err != nil {
		return false, err
	}
	if wallet.Credit >= amount {
		return true, nil
	}
	return false, nil
}
