package utils

func FindCogsAccount(unitPrice float64) string {
	if unitPrice == 0 {
		return "32232"
	} else {
		return "21600"
	}
}
