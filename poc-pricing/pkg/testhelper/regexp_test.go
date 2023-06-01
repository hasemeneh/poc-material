package testhelper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeRegexp(t *testing.T) {
	resp := NormalizeRegexpQuery("Select * from table where id = ?")
	require.Equal(t, `^Select \* from table where id = \?$`, resp)
}
