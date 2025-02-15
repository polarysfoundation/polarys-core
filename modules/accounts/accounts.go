package accounts

import (
	"github.com/polarysfoundation/polarys-core/modules/accounts/keystore"
	"github.com/polarysfoundation/polarys-core/modules/common"
)


type Account struct {
	Address    common.Address
	keypair    *keystore.Keypair
}



type Accounts struct {
	Accounts []*Account
}
