package retrier

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRetry_Failed(t *testing.T) {
	var result string
	err := Retry(time.Second*1, 5, func() error {
		temp, err2 := failedFunc()
		result = temp
		return err2
	})
	require.Error(t, err)
	require.Equal(t,
		"something is wrong",
		err.Error())
	require.Equal(t, "I'm a failure", result)
}

func failedFunc() (string, error) {
	return "I'm a failure", errors.New("something is wrong")
}

func TestRetry_Success(t *testing.T) {
	var result string
	err := Retry(time.Second*1, 5, func() error {
		temp, err2 := successFunc()
		result = temp
		return err2
	})
	require.NoError(t, err)
	require.Equal(t, "I Succeed", result)
}

func successFunc() (string, error) {
	return "I Succeed", nil
}
