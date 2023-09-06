package utils

import (
	"fmt"
	"illum_sales_file_loader/ftp_handler"
)

func CreateOrdersSap() error {
	data, err := GetAllCSVFilesFromFTP()
	if err != nil {
		return err
	}

	ItemBarCodeCollection, err := GetItemsFromSap()
	if err != nil {
		return err
	}

	for _, salesDay := range data {
		err := formatCSVLinesAndPostOrder(salesDay.salesDataString, ItemBarCodeCollection)
		if err != nil {
			SendUnknownErrorToTeams(err)
			continue
		}

		err = ftp_handler.DeleteFtpFile(salesDay.fileName)
		if err != nil {
			SendUnknownErrorToTeams(fmt.Errorf("error deleting file %s from FTP. Error: %v", salesDay.fileName, err))
		}
	}

	return nil
}
