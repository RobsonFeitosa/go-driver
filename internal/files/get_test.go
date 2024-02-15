package files

import (
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestGet() {
	setMockListFiles(ts.mock)

	_, err := Get(ts.conn, 1)
	assert.NoError(ts.T(), err)
}
