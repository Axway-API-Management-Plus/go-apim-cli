package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// ConfigureHAProxy : Cofigure HA Proxy for Master nodes
func (getProps *Properties) ConfigureHAProxy(sshConfig *ssh.ClientConfig) {
	client := &SSHClient{
		Config: sshConfig,
		Host:   getProps.LoadBalancer[0],
		Port:   22,
	}

	client.copyLocalFiletoRemote("resources/"+getProps.HAProxyCfg, getProps.HAProxyCfg)

	var hostIP string
	for i, master := range getProps.MasterNodes {
		if hostIP != "" {
			masterNode := "export MASTER_NODE" + strconv.Itoa(i+1) + "=" + master
			hostIP = hostIP + " && " + masterNode
		} else {
			hostIP = "export MASTER_NODE" + strconv.Itoa(i+1) + "=" + master
		}
	}

	haProxyCfg := hostIP + " && chmod +x " + getProps.HAProxyCfg + " && " + "./" + getProps.HAProxyCfg + " > haproxycfg.cfg" + " && " + appendHAProxyCfgFile

	haProxyCmds := []string{
		haProxyCfg,
		startHAProxy,
	}

	err := client.execCommand(haProxyCmds)
	if err != nil {
		PrettyPrintErr("Failed to configure and start HA Proxy on %s : %v", client.Host, err)
		os.Exit(1)
	}
}

func (getProps *Properties) createCerts(sshConfig *ssh.ClientConfig) {

	client := &SSHClient{
		Config: sshConfig,
		Host:   getProps.LoadBalancer[0],
		Port:   22,
	}

	cacsr := getProps.CaCsr
	caConfig := getProps.CaConfig
	KubernetesCsr := getProps.KubernetesCsr

	client.copyLocalFiletoRemote("resources/"+cacsr, cacsr)
	client.copyLocalFiletoRemote("resources/"+caConfig, caConfig)
	client.copyLocalFiletoRemote("resources/"+KubernetesCsr, KubernetesCsr)

	nodes := ""

	for _, master := range getProps.MasterNodes {
		if nodes != "" {
			nodes = nodes + "," + master
		} else {
			nodes = master
		}
	}

	for _, lb := range getProps.LoadBalancer {
		nodes = nodes + "," + lb
	}
	cmdCertAuth := "cfssl gencert -initca " + cacsr + "  | cfssljson -bare ca"
	etcdCertsCreate := "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=" + caConfig + " -hostname=" + nodes +
		",127.0.0.1,kubernetes.default -profile=kubernetes " + KubernetesCsr + " |  cfssljson -bare kubernetes"

	cmdCreateCert := []string{
		cmdCertAuth,
		etcdCertsCreate,
	}
	PrettyPrintOk("Creating CA and Kube certificates ")
	err := client.execCommand(cmdCreateCert)
	if err != nil {
		PrettyPrintErr("Certificate Creation failed on %s : %v", client.Host, err)
		os.Exit(1)
	}

	PrettyPrintOk("Copying Certificates to Client ")
	client.copyRemoteFiletoLocal("ca.pem", "ca.pem")
	client.copyRemoteFiletoLocal("kubernetes-key.pem", "kubernetes-key.pem")
	client.copyRemoteFiletoLocal("kubernetes.pem", "kubernetes.pem")

}
