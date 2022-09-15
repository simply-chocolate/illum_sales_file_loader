package ftp_handler

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

func getFtpClient() (ftp.ServerConn, error) {
	ftpClient, err := ftp.Dial(os.Getenv("FTP_HOST") + ":" + os.Getenv("FTP_PORT"))
	if err != nil {
		return ftp.ServerConn{}, fmt.Errorf("error connecting to the ftp server at file. error: %v", err)
	}

	if err = ftpClient.Login(os.Getenv("FTP_USER"), os.Getenv("FTP_PASS")); err != nil {
		return ftp.ServerConn{}, fmt.Errorf("error getting authenticaed at the ftp server at file. error: %v", err)
	}

	return *ftpClient, nil
}
