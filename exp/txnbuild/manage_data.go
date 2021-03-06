package txnbuild

import (
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

// ManageData represents the Stellar manage data operation. See
// https://www.stellar.org/developers/guides/concepts/list-of-operations.html
type ManageData struct {
	Name  string
	Value []byte
}

// BuildXDR for ManageData returns a fully configured XDR Operation.
func (md *ManageData) BuildXDR() (xdr.Operation, error) {
	xdrOp := xdr.ManageDataOp{DataName: xdr.String64(md.Name)}

	// No data value clears the named data entry on the account
	if md.Value == nil {
		xdrOp.DataValue = nil
	} else {
		xdrDV := xdr.DataValue(md.Value)
		xdrOp.DataValue = &xdrDV
	}

	opType := xdr.OperationTypeManageData
	body, err := xdr.NewOperationBody(opType, xdrOp)

	return xdr.Operation{Body: body}, errors.Wrap(err, "Failed to build XDR OperationBody")
}
