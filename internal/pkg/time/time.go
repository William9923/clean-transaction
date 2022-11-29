package time

import "time"

var asiaJakartaTz, _ = time.LoadLocation("Asia/Jakarta")

func Now() time.Time {
	return time.Now().In(asiaJakartaTz)
}
