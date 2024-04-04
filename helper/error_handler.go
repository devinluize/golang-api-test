package helper

import (
	"gorm.io/gorm"
	"io"
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
func CloseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		// Handle the error, such as logging or returning an error response
	}
}
