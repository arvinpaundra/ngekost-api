package contract

import "gorm.io/gorm"

type TxBeginner interface {
	Begin() *gorm.DB
}
