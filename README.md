## Axway API Manager CLI?
apimanager is a command line (CLI) utility which allows you to manage Axway's API Manager resources. This utility also allows you to automate the deployment of API Manager resources. It also allows you to view and edit the existing resources. This CLI is develop using cobra and viper (golang packages).

# Installation



# Supported Command

This CLI supports all the CRUD operations on the API Manager resources. 
 ```
$ apimanager 
apimanager controls and manages API Manager resources.

Usage:
  apimanager [command]

Available Commands:
  create      Create an API Manager resource
  delete      delete an API Manager resource
  describe    describe an API Manager resource
  edit        edit an API Manager resource
  help        Help about any command
  list        list all API Manager resource
  login       Stores the login info of API Manager
  publish     publish a proxy
  unpublish   unpublish a proxy

Flags:
      --config string   config file (default is $HOME/.apimanager.yaml)
  -h, --help            help for apimanager
  -v, --version         version for apimanager

Use "apimanager [command] --help" for more information about a command.
 ```

# Examples

apimanager login

## Create apimanager resources

* apimanager create org -n Marvel -ed 
* apimanager create user -n 'AntMan' -l antman -o Marvel -r user -p antman
* apimanager create user -n 'Thor' -l thor -o Marvel -r admin
* apimanager create user -n 'Captain America' -l captain -o Marvel -r admin
* apimanager create user -n 'Iron Man' -l ironman -o Marvel -r oadmin
* apimanager create app -n Avengers -o Marvel
* apimanager create key -a Avengers
* apimanager create oauth -a Avengers -c resources/cert.pem 
* apimanager create api -n 'Captain America' -o 'Marvel' -f resources/swagger.json 
* apimanager create proxy -n 'The First Avenger' -b 'Captain America' -c resources/cert.pem -o Marvel -s passthrough
* apimanager create proxy -n 'The Winter Soldier' -b 'Captain America' -c resources/cert.pem -o Marvel -s apikey -a Avengers
* apimanager create proxy -n 'Civil War' -b 'Captain America' -c resources/cert.pem -o Marvel -s oauth -a Avengers
* apimanager create api -n 'Iron Man' -o 'Marvel' -f resources/swagger.json 
* apimanager create proxy -n 'Iron Man' -b 'Iron Man' -c resources/cert.pem -o Marvel 
* apimanager create proxy -n 'Iron Man 2' -b 'Iron Man' -c resources/cert.pem -o Marvel -s apikey -a Avengers
* apimanager create proxy -n 'Iron Man 3' -b 'Iron Man' -c resources/cert.pem -o Marvel -s oauth -a Avengers

## Listing apimanager resources

* apimanager list orgs
* apimanager list users
* apimanager list apps
* apimanager list keys -a 'Avengers'
* apimanager list oauths -a 'Avengers'
* apimanager list apis
* apimanager list proxies

## Describe apimanager resources

* apimanager describe org -n 'Marvel'
* apimanager describe user -n 'Iron Man'
* apimanager describe app -n Asgard
* apimanager describe api -n 'Captain America'
* apimanager describe proxy -n 'The First Avenger'

## Edit apimanager resources

* apimanager edit org -n 'Marvel'
* apimanager edit user -n 'Iron Man'
* apimanager edit app -n Asgard

## Publish or unpublish api proxy

* apimanager unpublish proxy -n 'The First Avenger'
* apimanager unpublish proxy -n 'The Winter Soldier'
* apimanager unpublish proxy -n 'Civil War'

## Delete apimanager resources

* apimanager delete proxy -n 'The First Avenger'
* apimanager delete api -n 'Captain America'
* apimanager delete key -k 90b349a0-482c-44d4-b79a-fec70c71f809 -a Asgard
* apimanager delete oauth -k 6126028c-5230-4b3b-b27b-f28f26dae28b -a Asgard
* apimanager delete app -n Asgard
* apimanager delete user -n 'AntMan'
* apimanager delete user -n Thor
* apimanager delete user -n 'Iron Man'
* apimanager delete org -n Marvel
