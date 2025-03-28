package utils

import (
	"errors"

	"gorm.io/gorm"
)

func (user *User) ExecWithCoins(
	db *gorm.DB,
	costs uint,
	permOnly bool,
	logicFunc func(tx *gorm.DB) error,
) error {
	return db.Transaction(func(tx *gorm.DB) error {

		balance := user.DailyCoin
		if !permOnly {
			balance += user.HonorCoin
		}

		if balance < costs {
			return errors.New("金币不足")
		}

		if permOnly {
			user.HonorCoin -= costs
		} else if user.DailyCoin < costs {
			user.HonorCoin -= (costs - user.DailyCoin)
			user.DailyCoin = 0
		} else {
			user.DailyCoin -= costs
		}

		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return logicFunc(tx)
	})
}
