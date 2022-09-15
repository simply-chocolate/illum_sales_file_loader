package ftp_handler

import (
	"fmt"

	"github.com/jlaffaye/ftp"
)

func GetFtpFileList() ([]*ftp.Entry, error) {
	ftpClient, err := getFtpClient()
	if err != nil {
		return []*ftp.Entry{}, err
	}
	defer ftpClient.Quit()
	// Use list to see all the files in the FTP folder

	fileList, err := ftpClient.List("/illum/IL15147/Sales")
	if err != nil {
		return []*ftp.Entry{}, fmt.Errorf("error error getting the list of files. error: %v", err)
	}

	return fileList, nil
}
