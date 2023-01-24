package main

import (
	"fmt"
	"io/ioutil"

	"github.com/povsister/scp"
)

var (
	localPath  = "local paht here"
	remotePath = "rempte path here"
	remoteIP   = "type remote path here"
	ipAddres   = "remote ip addres here" // prefere to read it from sepatrate file
	password   = ""                      //remote pass must be in serparate file . use Getpass function
)

func main() {

	err := transfare(ipAddres, password)
	fmt.Println(err)
}

func transfare(remoteIP, password string) error {
	// conf
	sshConf := scp.NewSSHConfigFromPassword("root", password)

	// new client
	scpClient, err := scp.NewClient(remoteIP, sshConf, &scp.ClientOption{})
	if err != nil {
		return err
	}

	return upload(scpClient, localPath, remotePath)
}

func upload(client *scp.Client, localPath, remotePath string) error {
	// upload file
	return client.CopyFileToRemote(localPath, remotePath, &scp.FileTransferOption{})

}

// to scure app read pass form seprite file
func Getpass() string {
	data, err := ioutil.ReadFile(".mypass")
	if err != nil {
		return err.Error()
	}
	psw := string(data)
	return psw[:len(psw)-1]
}
