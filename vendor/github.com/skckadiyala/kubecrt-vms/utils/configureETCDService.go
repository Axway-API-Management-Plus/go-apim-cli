package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// StartETCDService .
func (props *Properties) startETCDService(sshConf *ssh.ClientConfig) {

	props.createCerts(sshConf)

	etcdServiceFile := props.EtcdService
	configFile := props.ConfigFile

	for i := range props.MasterNodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   props.MasterNodes[i],
			Port:   22,
		}
		client.copyCertstoMaster()
		client.copyLocalFiletoRemote("resources/"+etcdServiceFile, etcdServiceFile)
		client.copyLocalFiletoRemote("resources/"+configFile, configFile)

		err := props.createEtcdK8sCfg(client)
		if err != nil {
			PrettyPrintErr("Failed to start etcd Service on :%s", client.Host)
			os.Exit(1)
		}
	}
}

func (props *Properties) createEtcdK8sCfg(client *SSHClient) error {

	hostIP := "export HOST_IP=" + client.Host
	for i, master := range props.MasterNodes {
		masterNode := "export MASTER_NODE" + strconv.Itoa(i+1) + "=" + master
		hostIP = hostIP + " && " + masterNode
	}
	exportNodes := hostIP + " && " + "export LB_NODE1=" + props.LoadBalancer[0]

	prepETCD := hostIP + " && chmod +x " + props.EtcdService + " && " + "./" + props.EtcdService + " > etcd.service"
	prepkubeCfg := exportNodes + " && " + "chmod +x config.sh && ./" + props.ConfigFile + " > config.yaml"
	rmTempFile := "rm " + props.EtcdService + " " + props.ConfigFile

	perpConfigs := []string{
		prepETCD,
		prepkubeCfg,
	}
	perpConfigs = append(perpConfigs, etcdConfig...)
	perpConfigs = append(perpConfigs, rmTempFile)

	err := client.execCommand(perpConfigs)
	if err != nil {
		PrettyPrintErr("Failed Execution on host %s:  ", client.Host)
		return err
	}
	return nil
}
