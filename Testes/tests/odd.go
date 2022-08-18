package tests

const (
	ODD_KEYORD  = "ODD"
	EVEN_KEYORD = "EVEN"
)

func IsOdd(value int64) string {
	if value%2 == 0 {
		return EVEN_KEYORD
	}
	return ODD_KEYORD
}
