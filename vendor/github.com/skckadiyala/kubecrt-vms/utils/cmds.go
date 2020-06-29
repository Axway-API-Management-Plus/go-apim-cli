package utils

const (
	restartKubelet     = "sudo systemctl restart kubelet"
	checkInternet      = "curl -o /dev/null -s -w \"%{http_code}\" http://www.google.com"
	startHAProxy       = "sudo systemctl restart haproxy"
	cmdkubeTokens      = "sudo kubeadm token list"
	cmdkubeInit        = "sudo kubeadm init --config=config.yaml"
	cmdrmPKIFiles      = "sudo rm -r ~/pki && sudo rm kubejoin.cmd"
	cmdCNIWeaveNwk     = "kubectl apply -f \"https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')\""
	cmdCNIrmWeaveNwk   = "kubectl delete -f \"https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')\""
	cmdCNICalicorbac   = "kubectl apply -f https://docs.projectcalico.org/v3.1/getting-started/kubernetes/installation/hosted/rbac-kdd.yaml"
	cmdCNIrmCalicorbac = "kubectl delete -f https://docs.projectcalico.org/v3.1/getting-started/kubernetes/installation/hosted/rbac-kdd.yaml"
	cmdCNICalicoNwk    = "kubectl apply -f https://docs.projectcalico.org/v3.1/getting-started/kubernetes/installation/hosted/kubernetes-datastore/calico-networking/1.7/calico.yaml"
	cmdCNIrmCalicoNwk  = "kubectl delete -f https://docs.projectcalico.org/v3.1/getting-started/kubernetes/installation/hosted/kubernetes-datastore/calico-networking/1.7/calico.yaml"
)

var verifykubeNodes = []string{
	checkInternet,
	"sudo docker version -f 'Docker Client: {{.Client.Version}}'",
	"kubeadm version -o short",
	"kubelet --version",
	"kubectl version --client --short",
}

var verifyMaster = []string{
	"etcd --version",
	cmdkubeTokens,
}

var verifyHANodes = []string{
	checkInternet,
	"type cfssl && type cfssljson",
	"haproxy -v",
}

var etcdConfig = []string{
	"sudo mkdir -p /etc/etcd /var/lib/etcd  && sudo mv ~/ca.pem ~/kubernetes.pem ~/kubernetes-key.pem /etc/etcd",
	"sudo mv etcd.service /etc/systemd/system/etcd.service",
	"sudo systemctl daemon-reload && sudo systemctl enable etcd && sudo systemctl stop etcd  &&  sudo systemctl start etcd",
}

var primaryMaster = []string{
	cmdkubeTokens,
	cmdkubeInit,
	"mkdir -p $HOME/pki && sudo cp -r /etc/kubernetes/pki/* $HOME/pki/ && sudo chown -R $(id -u):$(id -g) $HOME/pki",
	"mkdir -p $HOME/.kube && sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config",
}

var seconderyMasters = []string{
	cmdkubeTokens,
	"rm $HOME/pki/apiserver* && sudo mv $HOME/pki /etc/kubernetes/",
	cmdkubeInit,
}

var cmdkubeReset = []string{
	"sudo kubeadm reset -f",
}

var cmdResetMasters = []string{
	"sudo systemctl stop etcd && sudo rm /etc/etcd/*.pem && sudo rm /etc/systemd/system/etcd.service && sudo rm -r /var/lib/etcd/",
	"sudo kubeadm reset -f",
	"sudo rm $HOME/config.yaml",
}

var cmdRestartKubelet = []string{
	restartKubelet,
}

const appendHAProxyCfgFile = "if [ \"$(grep  kubernetes-master-nodes /etc/haproxy/haproxy.cfg | wc -l)\"  -eq 0 ]; then " +
	"echo 'INFO: Updating HA Proxy Configfile with Kuberneties Master Nodes' ;" +
	"grep -Fq 'backend kubernetes-master-nodes'  /etc/haproxy/haproxy.cfg  || echo 'backend kubernetes-master-nodes' | sudo tee --append /etc/haproxy/haproxy.cfg < haproxycfg.cfg" +
	"; else " +
	"echo 'Kubeternies Master hosts IPs exits on ha proxy configuration file.Verify the IPs'            [WARN] ;" +
	"fi"

var removeHAProxyCfg = []string{
	"sudo sed -i '/# Begin k8s HAConfig/,/# End k8s HAConfig/d' /etc/haproxy/haproxy.cfg",
}
