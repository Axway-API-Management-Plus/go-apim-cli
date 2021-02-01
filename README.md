## Axway API Manager CLI
apimctl is a command line (CLI) utility which allows you to manage Axway's API Manager resources. This utility also allows you to automate the deployment of API Manager resources. It also allows you to view and edit the existing resources. This CLI is develop using [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) (golang packages).

# Installation

```
curl -LJO https://github.com/Axway-API-Management-Plus/go-apim-cli/releases/download/2.0.0/apimctl_darwin_amd64
chmod +x apimctl_darwin_amd64
mv apimctl_darwin_amd64 /usr/local/bin/apimctl 
```

# Supported Command

This CLI supports all the CRUD operations on the API Manager resources. 
 ```
$ apimctl 
apimctl controls and manages API Manager resources.

Usage:
  apimctl [command]

Available Commands:
  create      Create an API Manager resource from a file
  delete      delete an API Manager resource
  describe    describe an API Manager resource
  edit        edit an API Manager resource
  help        Help about any command
  list        list all API Manager resource
  login       Stores the login info of API Manager
  publish     publish a proxy
  unpublish   unpublish a proxy
  update      Update an API Manager resource

Flags:
      --config string   config file to store connection details (default ".apimanager")
  -h, --help            help for apimctl
  -v, --version         version for apimctl

Use "apimctl [command] --help" for more information about a command.
 ```

# Examples
```
apimctl login
```
## Create API Manager resources

* apimctl create org -n Marvel -ed 
* apimctl create user -n 'AntMan' -l antman -o Marvel -r user -p antman
* apimctl create user -n 'Thor' -l thor -o Marvel -r admin
* apimctl create user -n 'Captain America' -l captain -o Marvel -r admin
* apimctl create user -n 'Iron Man' -l ironman -o Marvel -r oadmin
* apimctl create app -n Avengers -o Marvel
* apimctl create key -a Avengers
* apimctl create oauth -a Avengers -c resources/cert.pem 
* apimctl create api -n 'Captain America' -o 'Marvel' -f resources/swagger.json 
* apimctl create proxy -n 'The First Avenger' -b 'Captain America' -c resources/cert.pem -o Marvel -s passthrough
* apimctl create proxy -n 'The Winter Soldier' -b 'Captain America' -c resources/cert.pem -o Marvel -s apikey -a Avengers
* apimctl create proxy -n 'Civil War' -b 'Captain America' -c resources/cert.pem -o Marvel -s oauth -a Avengers
* apimctl create access -a Avengers -n 'The Winter Soldier'
* apimctl create access -a Avengers -n 'Civil War'
* apimctl create app -n X-Men -o Marvel
* apimctl create key -a X-Men
* apimctl create oauth -a X-Men -c resources/cert.pem 
* apimctl create api -n 'X-Men' -o 'Marvel' -f resources/swagger.json 
* apimctl create proxy -n 'First Class' -b 'X-Men' -c resources/cert.pem -o Marvel 
* apimctl create proxy -n 'Days of the Future Past' -b 'X-Men' -c resources/cert.pem -o Marvel -s apikey -a Avengers
* apimctl create proxy -n 'Apocalypse' -c resources/cert.pem -o Marvel -s oauth -a Avengers
* apimctl create access -a X-Men -n 'Days of the Future Past'
* apimctl create access -a X-Men -n 'Apocalypse'

## Listing API Manager resources

* apimctl list orgs
* apimctl list users
* apimctl list apps
* apimctl list keys -a 'Avengers'
* apimctl list oauths -a 'Avengers'
* apimctl list apis
* apimctl list proxies
* apimctl list access -a 'Avengers'
* apimctl list access -a 'X-Men'

## Describe API Manager resources

* apimctl describe org -n 'Marvel'
* apimctl describe user -n 'Iron Man'
* apimctl describe app -n Avengers
* apimctl describe api -n 'Captain America'
* apimctl describe proxy -n 'The First Avenger'

## Edit API Manager resources

* apimctl edit org -n 'Marvel'
* apimctl edit user -n 'Iron Man'
* apimctl edit app -n Asgard

## Publish or unpublish api proxy

* apimctl unpublish proxy -n 'The First Avenger'
* apimctl unpublish proxy -n 'The Winter Soldier'
* apimctl unpublish proxy -n 'Civil War'
* apimctl unpublish proxy -n 'First Class'
* apimctl unpublish proxy -n 'Days of the Future Past'
* apimctl unpublish proxy -n 'Apocalypse'
* 
## Delete API Manager resources

* apimctl delete proxy -n 'The First Avenger'
* apimctl delete api -n 'Captain America'
* apimctl delete key -k 90b349a0-482c-44d4-b79a-fec70c71f809 -a Avengers
* apimctl delete oauth -k 6126028c-5230-4b3b-b27b-f28f26dae28b -a Avengers
* apimctl delete app -n Avengers
* apimctl delete user -n 'AntMan'
* apimctl delete user -n Thor
* apimctl delete user -n 'Iron Man'
* apimctl delete org -n Marvel
