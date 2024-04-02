package helper

import (
	"gorm.io/gorm"
)

func PanifIfError(err error) {
	if err != nil {
		panic(err)
	} else {

	}
}
func Rollbackiferror(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errTx := tx.Rollback()
		PanifIfError(errTx.Error)
		panic(err)
	} else {
		errTx := tx.Commit()
		PanifIfError(errTx.Error)
	}
}
