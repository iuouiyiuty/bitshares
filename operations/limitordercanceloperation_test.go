package operations

import (
	"github.com/iuouiyiuty/bitshares/gen/data"
	"github.com/iuouiyiuty/bitshares/types"
)

func (suite *operationsAPITest) Test_LimitOrderCancelOperation() {

	op := LimitOrderCancelOperation{
		Extensions: types.Extensions{},
	}

	sample, err := data.GetSampleByType(op.Type())
	if err != nil {
		suite.FailNow(err.Error(), "GetSampleByType")
	}

	if err := op.UnmarshalJSON([]byte(sample)); err != nil {
		suite.FailNow(err.Error(), "UnmarshalJSON")
	}

	suite.RefTx.Operations = types.Operations{
		types.Operation(&op),
	}

	suite.compareTransaction(suite.RefTx, false)
}
