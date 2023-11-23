package txBeginner

import (
	"database/sql"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"gorm.io/gorm"
)

type txBeginner struct {
	db *gorm.DB
}

func NewTxBeginner(db *gorm.DB) contract.TxBeginner {
	return &txBeginner{db: db}
}

func (b *txBeginner) Begin() *gorm.DB {
	return b.db.Begin(&sql.TxOptions{})
}
