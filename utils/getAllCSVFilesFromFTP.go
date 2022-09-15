package utils

import (
	"illum_sales_file_loader/ftp_handler"
)

func GetAllCSVFilesFromFTP() ([]string, error) {
	fileList, err := ftp_handler.GetFtpFileList()
	if err != nil {
		return []string{}, err
	}

	salesDataString := []string{}

	for _, file := range fileList {
		salesData, err := ftp_handler.GetFtpFile(file.Name)
		if err != nil {
			return salesDataString, err
		}

		salesDataString = append(salesDataString, salesData)
	}

	return salesDataString, nil
}
