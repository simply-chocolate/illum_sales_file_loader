package utils

func CreateOrdersSap() error {
	salesDays, err := GetAllCSVFilesFromFTP()
	if err != nil {
		return err
	}

	ItemBarCodeCollection, err := GetItemsFromSap()
	if err != nil {
		return err
	}

	for _, salesDay := range salesDays {
		err := formatCSVLinesAndPostOrder(salesDay, ItemBarCodeCollection)
		if err != nil {
			return err
		}
	}

	return nil
}
