package cmd

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/tabwriter"

	"git.ecd.axway.int/apigov/kubecrt-vms/utils"
	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/spf13/viper"
)

var (
	file         string // blob file location
	config       string // config file location
	orgName      string
	appName      string
	keyID        string
	oauthID      string
	name         string
	password     string
	apiName      string
	userName     string
	userRole     string
	loginName    string
	security     string
	resourcePath string
	certPath     string
	proxyState   string
	proxyVersion string
	image        string
	enabled      bool
	development  bool
)

type configAPI struct {
	APIManagerHost string `yaml:"apiManagerHost"`
	APIManagerPort string `yaml:"apiManagerPort"`
	Authorization  string `yaml:"authorization"`
}

func getConfig() *apimgr.Configuration {
	if viper.GetString("apimanagerhost") == "" || viper.GetString("apimanagerport") == "" {
		utils.PrettyPrintErr("Please login to API Manager, use 'login' command")
		os.Exit(0)
	}

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	cfg := apimgr.NewConfiguration()
	cfg.Host = viper.GetString("apimanagerhost") + ":" + viper.GetString("apimanagerport")
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Basic "+viper.GetString("authorization"))
	cfg.HTTPClient = &http.Client{Transport: transCfg}
	return cfg
}

func fmtDisplay() *tabwriter.Writer {
	writeTab := new(tabwriter.Writer)
	writeTab.Init(os.Stdout, 0, 8, 0, '\t', 0)
	return writeTab
}

func getUniqueID(nbofBytes int) string {
	b := make([]byte, nbofBytes)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("Error generating uniqueID: %v \n", err)
	}
	uuid := fmt.Sprintf("%x", b[0:nbofBytes])
	return uuid
}

func createTempFile(fileName string, content []byte) error {
	err := ioutil.WriteFile("/tmp/"+fileName, content, 0644)
	if err != nil {
		fmt.Println("Failed to write temp file", err)
		return err
	}
	cm := exec.Command("vim", "/tmp/"+fileName)
	cm.Stdin = os.Stdin
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	err = cm.Run()
	if err != nil {
		fmt.Println("Failed to open editor", err)
		return err
	}
	return nil
}
func deleteTempFile(fileName string) {
	cmd := exec.Command("rm", "/tmp/"+fileName)
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
		return
	}
}
