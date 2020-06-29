package utils

import (
	"bufio"
	"os"

	"golang.org/x/crypto/ssh"
)

// JoinWorkerstoMaster : Join worker nodes to the master node
func (props *Properties) JoinWorkerstoMaster(sshConf *ssh.ClientConfig) {

	for w := range props.MinionNodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.MinionNodes[w],
			Port:   22,
		}
		err := joinWorkerNodetoMaster(client)
		if err != nil {
			PrettyPrintErr("Worker node failed to join master", err)
			os.Exit(1)
		}

	}
}

func joinWorkerNodetoMaster(client *SSHClient) error {

	file, err := os.OpenFile("kubejoin.cmd", os.O_RDWR, 0644)
	if err != nil {
		PrettyPrintErr("Failed to Open the file to read kubeadm join cmd", err)
		PrettyPrintErr("kubeadm.cmd file should be located in the first master node", err)
		PrettyPrintErr("copy file to the client machine", err)
	}
	r := bufio.NewReader(file)
	s, _, e := r.ReadLine()
	if e != nil {
		PrettyPrintErr("Failed to Read the file kubejoin.cmd", err)
	}
	PrettyPrintOk("Using join command: \n%s", string(s))
	cmdStartk8Cluster := &SSHCommand{Path: "sudo " + string(s),
		Stdout: os.Stdout,
		Stderr: os.Stderr}
	PrettyPrintOk("Worker %s joining k8sCluster ", client.Host)
	if err := client.RunCommand(cmdStartk8Cluster); err != nil {
		PrettyPrintErr("Worker node failed to join Master:%s", client.Host)
		return err
	}
	return nil
}

//RestartWorkerNodes ..
func (props *Properties) RestartWorkerNodes(sshConf *ssh.ClientConfig) {

	for w := range props.MinionNodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.MinionNodes[w],
			Port:   22,
		}
		err := RestartWorkerNode(client)
		if err != nil {
			PrettyPrintErr("Worker node failed to join master", err)
			os.Exit(1)
		}

	}
}

// RestartWorkerNode ...
func RestartWorkerNode(client *SSHClient) error {

	if err := client.execCommand(cmdRestartKubelet); err != nil {
		PrettyPrintErr("Worker node failed to join Master:%s", client.Host)
		return err
	}
	return nil
}
