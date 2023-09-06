package ftp_handler

import (
	"fmt"
	"io"
)

// TODO: We need to either delete the file from the server as the last thing. So when we get a positive feedback value from SAP, we send the delete request to the FTP.
func GetFtpFile(fileName string) (string, error) {
	ftpClient, err := getFtpClient()
	if err != nil {
		return "", err
	}
	defer ftpClient.Quit()

	csvFile, err := ftpClient.Retr(fmt.Sprintf("/illum/IL15147/Sales/%s", fileName))
	if err != nil {
		return "", fmt.Errorf("error retrieving file %s from the ftp server. error: %v ", fileName, err)
	}
	defer csvFile.Close()

	csvAsBuffer, err := io.ReadAll(csvFile)

	if err != nil {
		return "", fmt.Errorf("error reading the file %s. error: %v", fileName, err)
	}

	return string(csvAsBuffer), nil
}
