package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SSHCommand ..
type SSHCommand struct {
	Path   string
	Env    []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// SSHClient ..
type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

// RunCommand   to run ssh command
func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
	var (
		session *ssh.Session
		err     error
	)
	if session, err = client.newSession(); err != nil {
		return err
	}
	defer session.Close()

	if cmd.Path == "sudo kubeadm init --config=config.yaml" {
		fmt.Println("this might take a minute or longer if the control plane images have to be pulled")
		output, err := session.CombinedOutput(cmd.Path)
		if err != nil {
			PrettyPrintErr("Error Executing the cmd", err)
		}
		kubeadmOut := string(output)
		kubeJoinIndex := strings.Index(kubeadmOut, "kubeadm join")
		if kubeJoinIndex == -1 {
			PrettyPrintErr("Failed to execute kubeadm init: \n %v", kubeadmOut)
			return errors.New("Failed to execute kubeadm")
		}
		fmt.Println(kubeadmOut)
		var _, errfile = os.Stat("kubejoin.cmd")
		if os.IsNotExist(errfile) {
			fo, err := os.Create("kubejoin.cmd")
			if err != nil {
				fmt.Println("Failed to create file")
			}
			defer fo.Close()
		}
		file, err := os.OpenFile("kubejoin.cmd", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			PrettyPrintErr("Failed to open file kubejoin.cmd: %v", err.Error())
		}
		_, err = file.WriteString(kubeadmOut[kubeJoinIndex : len(kubeadmOut)-1])
		if err != nil {
			fmt.Println("Failed to write to the file")
		}
	} else if cmd.Path == "curl -o /dev/null -s -w \"%{http_code}\" http://www.google.com" {
		output, err := session.CombinedOutput(cmd.Path)
		if err != nil {
			PrettyPrintErr("Error Executing the cmd", err)
		}
		statusCode := string(output)
		fmt.Println("Internet access Status: ", statusCode)
		if statusCode != "200" {
			fmt.Println("Please verify internet access on Node")
			err = errors.New("No Internet Access")
			return err
		}
	} else {
		if err = client.prepareCommand(session, cmd); err != nil {
			return err
		}
		err = session.Run(cmd.Path)
		return err
	}
	return nil
}

func (client *SSHClient) prepareCommand(session *ssh.Session, cmd *SSHCommand) error {

	for i, env := range cmd.Env {
		fmt.Println(i)
		variable := strings.Split(env, "=")
		if len(variable) != 2 {
			continue
		}
		if err := session.Setenv(variable[0], variable[1]); err != nil {
			return err
		}
	}

	if cmd.Stdin != nil {
		stdin, err := session.StdinPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stdin for session: %v", err)
		}
		go io.Copy(stdin, cmd.Stdin)
	}
	if cmd.Stdout != nil {
		stdout, err := session.StdoutPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stdout for session: %v", err)
		}
		go io.Copy(cmd.Stdout, stdout)
	}

	if cmd.Stderr != nil {
		stderr, err := session.StderrPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stderr for session: %v", err)
		}
		go io.Copy(cmd.Stderr, stderr)
	}

	return nil
}

func (client *SSHClient) newSession() (*ssh.Session, error) {
	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}
	return session, nil
}

func (client *SSHClient) newFtpClient() (*sftp.Client, error) {
	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	sftpclient, err := sftp.NewClient(connection)
	if err != nil {
		PrettyPrintErr("Failed to create connections :%s -- %v", client.Host, err.Error())
	}
	return sftpclient, nil
}

func (client *SSHClient) copyLocalFiletoRemote(srcFile string, destFile string) {

	ftpClient, err := client.newFtpClient()
	if err != nil {
		PrettyPrintErr("Failed to create new ftp client:%s -- %v", client.Host, err.Error())
	}

	dFile, err := ftpClient.Create("./" + destFile)
	if err != nil {
		PrettyPrintErr("Failed to create dest file:%s -- %v", destFile, err.Error())
	}
	defer dFile.Close()
	sFile, err := os.Open("./" + srcFile)
	if err != nil {
		PrettyPrintErr("Failed to open source file:%s -- %v", srcFile, err.Error())
	}

	_, err = io.Copy(dFile, sFile)
	if err != nil {
		PrettyPrintErr("Failed to copy source file:%s -- %v", srcFile, err.Error())
	}
}

func (client *SSHClient) copyRemoteDirToLocal(dirPath string) {

	ftpClient, err := client.newFtpClient()
	if err != nil {
		PrettyPrintErr("Failed to create new ftp client:%s -- %v", client.Host, err.Error())
	}

	list, err := ftpClient.ReadDir(dirPath)
	os.Mkdir("pki", os.ModePerm)

	fmt.Println("Copying files to local ...")
	for i := range list {
		client.copyRemoteFiletoLocal(dirPath+"/"+list[i].Name(), dirPath+"/"+list[i].Name())
	}

}

func (client *SSHClient) copyLocalDirToRemote(dirPath string) {

	ftpClient, err := client.newFtpClient()
	if err != nil {
		PrettyPrintErr("Failed to create new ftp client:%s -- %v", client.Host, err.Error())
	}
	list, err := ioutil.ReadDir(dirPath)
	sshFxFailure := uint32(4)
	err = ftpClient.Mkdir(dirPath)
	if status, ok := err.(*sftp.StatusError); ok {
		if status.Code == sshFxFailure {
			var fi os.FileInfo
			fi, err = ftpClient.Stat(dirPath)
			if err == nil {
				if !fi.IsDir() {
					PrettyPrintErr("File exists: %s", dirPath)
				}
			}
		}
	}
	fmt.Println("Copying files to remote ...")
	for i := range list {
		client.copyLocalFiletoRemote(dirPath+"/"+list[i].Name(), dirPath+"/"+list[i].Name())
	}

}

func (client *SSHClient) copyRemoteFiletoLocal(srcFile string, destFile string) {

	ftpClient, err := client.newFtpClient()
	if err != nil {
		PrettyPrintErr("Failed to create new ftp client:%s -- %v", client.Host, err.Error())
	}
	sFile, err := ftpClient.Open(srcFile)
	if err != nil {
		PrettyPrintErr("Failed to open source file:%s -- %v", srcFile, err.Error())
	}
	defer sFile.Close()

	dFile, err := os.Create(destFile)
	if err != nil {
		PrettyPrintErr("Failed to create dest file:%s -- %v", destFile, err.Error())
	}

	_, err = io.Copy(dFile, sFile)
	if err != nil {
		PrettyPrintErr("Failed to copy source file:%s -- %v", srcFile, err.Error())
	}
}

func (client *SSHClient) execCommand(cmdlist []string) error {
	for _, sshcmd := range cmdlist {
		runCommand := &SSHCommand{Path: sshcmd, Stdout: os.Stdout, Stderr: os.Stderr}
		if sshcmd == "sudo kubeadm token list" {
			PrettyPrintOk("Verifying kubernaties Installation on master node" + runCommand.Path)
			if err := client.RunCommand(runCommand); err == nil {
				PrettyPrintErr("Kuberneties cluster is already installed on : %s:  ", client.Host)
				return errors.New("use 'reset' flag to force install\n" +
					"(WARN: using this flag permanently deletes the existing k8s cluster before fresh install)")
			}
		} else {
			PrettyPrintOk("Execute Command: %s", runCommand.Path)
			if err := client.RunCommand(runCommand); err != nil {
				PrettyPrintErr("Failed Execution on host %s", client.Host)
				return err
			}
		}
	}
	return nil
}

func (client *SSHClient) copyCertstoMaster() {
	PrettyPrintOk("Copying Kubernetes Certificates to Master Node :%s", client.Host)
	client.copyLocalFiletoRemote("ca.pem", "ca.pem")
	client.copyLocalFiletoRemote("kubernetes-key.pem", "kubernetes-key.pem")
	client.copyLocalFiletoRemote("kubernetes.pem", "kubernetes.pem")
}
