package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

// Properties ...
type Properties struct {
	SSHInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"sshInfo"`
	MasterNodes   []string `json:"masterNodes"`
	MinionNodes   []string `json:"minionNodes"`
	LoadBalancer  []string `json:"loadBalancer"`
	CaConfig      string   `json:"ca-config"`
	CaCsr         string   `json:"ca-csr"`
	KubernetesCsr string   `json:"kubernetes-csr"`
	EtcdService   string   `json:"etcdService"`
	ConfigFile    string   `json:"configFile"`
	HAProxyCfg    string   `json:"haProxyCfg"`
	CNINetwork    string   `json:"cniNetwork"`
}

// VerifyNodes : Pre-flight check for Master and Worker Nodes
func (props *Properties) VerifyNodes(sshConf *ssh.ClientConfig) {
	preFlightChecks("HANode", props.LoadBalancer, sshConf, verifyHANodes)
	verifykubeMasters := append(verifykubeNodes, verifyMaster...)
	preFlightChecks("master", props.MasterNodes, sshConf, verifykubeMasters)
	preFlightChecks("minion", props.MinionNodes, sshConf, verifykubeNodes)
}

func preFlightChecks(ntype string, nodes []string, sshConf *ssh.ClientConfig, cmdlist []string) {
	for _, node := range nodes {
		client := &SSHClient{
			Config: sshConf,
			Host:   node,
			Port:   22,
		}
		PrintHeader("Pre-flight check on "+ntype+": "+node, '=')
		err := client.execCommand(cmdlist)
		if err != nil {
			PrettyPrintErr(err.Error())
			PrettyPrintErr("Failed Execution on host %s:  ", client.Host)
			os.Exit(1)
		}
	}
}

//ReadConfFile ..
func ReadConfFile(filename string) *Properties {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		PrettyPrintErr(err.Error())
		os.Exit(1)
	}
	var props Properties
	e := json.Unmarshal(bs, &props)
	if e != nil {
		PrettyPrintErr(e.Error())
	}
	return &props
}

// RemoveTempFiles .
func RemoveTempFiles() {
	PrettyPrintOk("Cleaning up Temp Files")
	os.Remove("ca.pem")
	os.Remove("kubernetes-key.pem")
	os.Remove("kubernetes.pem")
	os.RemoveAll("./pki")
}
