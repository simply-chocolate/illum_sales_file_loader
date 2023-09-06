package utils

import (
	"illum_sales_file_loader/ftp_handler"
)

type ftpData struct {
	fileName        string
	salesDataString string
}

func GetAllCSVFilesFromFTP() ([]ftpData, error) {
	fileList, err := ftp_handler.GetFtpFileList()
	if err != nil {
		return []ftpData{}, err
	}

	ftpDataList := []ftpData{}

	for _, file := range fileList {
		salesData, err := ftp_handler.GetFtpFile(file.Name)
		if err != nil {
			return []ftpData{}, err
		}

		data := ftpData{
			fileName:        file.Name,
			salesDataString: salesData,
		}
		ftpDataList = append(ftpDataList, data)

	}

	return ftpDataList, nil
}
