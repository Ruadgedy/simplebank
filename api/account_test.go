package api

import (
	db "github.com/Ruadgedy/simplebank/db/sqlc"
	"github.com/Ruadgedy/simplebank/util"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetAccountApi(t *testing.T)  {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


}

func randomAccount() db.Account {
	return db.Account{
		ID:        util.RandomInt(1,1000),
		Owner:     util.RandomOwner(),
		Balance:   util.RandomMoney(),
		Currency:  util.RandomCurrency(),
	}
}