package retrier

import "time"

func Retry(backoffincrement time.Duration, retries int, fn func() error) error {
	var err error
	for i := 0; i < retries; i++ {
		time.Sleep(time.Duration(i) * backoffincrement)
		err = fn()
		if err == nil {
			return nil
		}
	}
	return err
}
