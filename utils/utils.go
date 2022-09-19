package utils

func FindCogsAccount(unitPrice float64) string {
	if unitPrice == 0 {
		return "32740"
	} else {
		return "22200"
	}
}
