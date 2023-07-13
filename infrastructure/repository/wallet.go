package repositories

import (
	"errors"
	"fmt"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *PGRepository) GenerateWallet(wallet entities.Wallet) error {
	result := r.DB.Create(&wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *PGRepository) ChargeWallet(walletID, userID uuid.UUID, amount int) error {

	// 1. Check if there is a wallet with the given walletID
	var wallet entities.Wallet
	err := r.DB.First(&wallet, "id = ?", walletID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("wallet not found")
		}
		return err
	}

	// 2. Check if the userID of the found wallet matches the provided userID
	if wallet.UserID != userID {
		return fmt.Errorf("userID is not match with walletID")
	}

	// 3. Update the wallet's credit by adding the amount
	wallet.Credit += amount
	err = r.DB.Save(&wallet).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *PGRepository) GetWallet(walletID uuid.UUID) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.DB.First(&wallet, "id = ?", walletID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("wallet not found")
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *PGRepository) WithdrawFromWallet(userid uuid.UUID, amount int) error {
	var wallet entities.Wallet
	err := r.DB.Model(&entities.Wallet{}).Where("user_id = ?", userid).First(&wallet).Error
	if err != nil {
		return err
	}
	if wallet.Credit-amount < 0 {
		return errors.New("Not enough credit")
	}
	wallet.Credit = wallet.Credit - amount

	err = r.DB.Save(&wallet).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisDB) SetPaymentRequest(orderID uuid.UUID, payerID, refID, amount string) error {
	key := fmt.Sprintf("%s-%v-%s", payerID, orderID.ID(), refID)
	r.Client.Set(key, amount, -1)
	return nil
}
func (r *RedisDB) VerifyPaymentRequest(payerID, refID, orderID string) (string, bool, error) {
	keys, _ := r.Client.Keys(fmt.Sprintf("%s-%s-%s", payerID, orderID, refID)).Result()
	if len(keys) > 0 {
		amount := r.Client.Get(keys[0])
		return amount.String(), true, nil
	}
	return "", false, errors.New("invalid payment")
}
