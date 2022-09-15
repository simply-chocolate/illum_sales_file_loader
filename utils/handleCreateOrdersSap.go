package utils

func CreateOrdersSap() error {
	salesDays, err := GetAllCSVFilesFromFTP()
	if err != nil {
		return err
	}

	for _, salesDay := range salesDays {
		formatCSVLinesToSapOrder(salesDay)
	}

	return nil
}
