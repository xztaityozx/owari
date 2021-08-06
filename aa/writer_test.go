package aa_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/owari/aa"
)

func Test_AAWriter_New(t *testing.T) {
	aaw := aa.NewAAWriter(os.Stdout)

	assert.NotNil(t, aaw)
}
