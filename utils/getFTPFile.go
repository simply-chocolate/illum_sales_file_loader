package utils

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

func ReadFileFtp() error {
	ftpClient, err := ftp.Dial(os.Getenv("FTP_HOST") + ":" + os.Getenv("FTP_PORT"))
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error connecting to the ftp server at file")
	}
	defer ftpClient.Quit()

	if err = ftpClient.Login(os.Getenv("FTP_USER"), os.Getenv("FTP_PASS")); err != nil {
		fmt.Println(err)
		return fmt.Errorf("error getting authenticaed at the ftp server at file")
	}

	// Use list to see all the files in the FTP folder

	fileList, err := ftpClient.List("/illum/IL15147/Sales")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error uploading the file server at file ")
	}
	fmt.Printf("%v", fileList[0])

	return nil
}
