package txpool

import (
	"math/big"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

type TxSigner struct {
	signer      common.Address
	balance     *big.Int
	feeIncoming *big.Int
}
