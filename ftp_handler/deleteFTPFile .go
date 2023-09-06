package ftp_handler

import (
	"fmt"
)

// TODO: We need to either delete the file from the server as the last thing. So when we get a positive feedback value from SAP, we send the delete request to the FTP.
func DeleteFtpFile(fileName string) error {
	ftpClient, err := getFtpClient()
	if err != nil {
		return err
	}
	defer ftpClient.Quit()

	err = ftpClient.Delete(fmt.Sprintf("/illum/IL15147/Sales/%s", fileName))
	if err != nil {
		return fmt.Errorf("error deleting file %s from the ftp server. error: %v ", fileName, err)
	}

	return nil
}
