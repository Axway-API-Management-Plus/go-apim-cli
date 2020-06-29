package utils

import (
	"os"

	"golang.org/x/crypto/ssh"
)

// Createk8sCluster ..
func (props *Properties) Createk8sCluster(sshConf *ssh.ClientConfig) {
	props.startETCDService(sshConf)
	for i := range props.MasterNodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.MasterNodes[i],
			Port:   22,
		}
		if props.MasterNodes[i] == props.MasterNodes[0] {
			if props.CNINetwork == "weave" {
				primaryMaster = append(primaryMaster, cmdCNIWeaveNwk)
			} else {

				primaryMaster = append(primaryMaster, cmdCNICalicorbac)
				primaryMaster = append(primaryMaster, cmdCNICalicoNwk)
			}
			err := client.execCommand(primaryMaster)
			if err != nil {
				PrettyPrintErr("Failed to create and start kuberneties cluster, fix the errors and redeploy")
				os.Exit(1)
			}
			client.copyLocalFiletoRemote("kubejoin.cmd", "kubejoin.cmd")
			client.copyRemoteDirToLocal("pki")
		} else {
			client.copyLocalDirToRemote("pki")
			err := client.execCommand(seconderyMasters)
			if err != nil {
				PrettyPrintErr("Failed to create and start kuberneties cluster, fix the errors and redeploy")
				os.Exit(1)
			}
		}
	}
}

// Deletek8sCluster Delete k8s atrifacts on the nodes
func (props *Properties) Deletek8sCluster(sshConf *ssh.ClientConfig) {

	var cmdReset []string
	for i := range props.MinionNodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.MinionNodes[i],
			Port:   22,
		}
		PrintHeader("Deleting Worker Node from the k8s master:  "+props.MinionNodes[i], '=')
		client.execCommand(cmdkubeReset)
	}

	for i := range props.MasterNodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.MasterNodes[i],
			Port:   22,
		}
		PrintHeader("Deleting master Node from k8s cluster: "+props.MasterNodes[i], '=')
		if props.MasterNodes[i] != props.MasterNodes[0] {
			client.execCommand(cmdResetMasters)
		} else {
			if props.CNINetwork == "weave" {
				cmdReset = append(cmdReset, cmdCNIrmWeaveNwk)
				cmdReset = append(cmdReset, cmdrmPKIFiles)
				cmdReset = append(cmdReset, cmdResetMasters...)
			} else {
				cmdReset = append(cmdReset, cmdCNIrmCalicorbac)
				cmdReset = append(cmdReset, cmdCNIrmCalicoNwk)
				cmdReset = append(cmdReset, cmdrmPKIFiles)
				cmdReset = append(cmdReset, cmdResetMasters...)
			}
			err := client.execCommand(cmdReset)
			if err != nil {
				PrettyPrintErr("Error while deleting the cluster")
			}
		}
	}

	for i := range props.LoadBalancer {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.LoadBalancer[i],
			Port:   22,
		}
		PrintHeader("Removing Kubernetes HAProxy Configuration from haproxy.cfg:  "+props.LoadBalancer[i], '=')
		client.execCommand(removeHAProxyCfg)
	}

	os.Remove("kubejoin.cmd")
}
