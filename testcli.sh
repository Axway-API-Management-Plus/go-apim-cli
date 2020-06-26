
apimctl login -H <> -P 8075 -u apiadmin -p <>

apimctl create org -n Marvel -ed 
apimctl create user -n 'AntMan' -l antman -o Marvel -r user -p antman
apimctl create user -n 'Thor' -l thor -o Marvel -r admin
apimctl create user -n 'Captain America' -l captain -o Marvel -r admin
apimctl create user -n 'Iron Man' -l ironman -o Marvel -r oadmin
apimctl create app -n Avengers -o Marvel
apimctl create key -a Avengers
apimctl create oauth -a Avengers -c resources/cert.pem 
apimctl create api -n 'Captain America' -o 'Marvel' -f resources/swagger.json 
apimctl create proxy -n 'The First Avenger' -b 'Captain America' -c resources/cert.pem -o Marvel -s passthrough
apimctl create proxy -n 'The Winter Soldier' -b 'Captain America' -c resources/cert.pem -o Marvel -s apikey -a Avengers
apimctl create proxy -n 'Civil War' -b 'Captain America' -c resources/cert.pem -o Marvel -s oauth -a Avengers
apimctl create access -a Avengers -n 'The Winter Soldier'
apimctl create access -a Avengers -n 'Civil War'

apimctl create app -n X-Men -o Marvel
apimctl create key -a X-Men
apimctl create oauth -a X-Men -c resources/cert.pem 
apimctl create api -n 'X-Men' -o 'Marvel' -f resources/swagger.json 
apimctl create proxy -n 'First Class' -b 'X-Men' -c resources/cert.pem -o Marvel 
apimctl create proxy -n 'Days of the Future Past' -b 'X-Men' -c resources/cert.pem -o Marvel -s apikey -a 'X-Men'
apimctl create proxy -n 'Apocalypse' -b 'X-Men' -c resources/cert.pem -o Marvel -s oauth -a 'X-Men'
apimctl create access -a X-Men -n 'Days of the Future Past'
apimctl create access -a X-Men -n 'Apocalypse'

apimctl list orgs
apimctl list users
apimctl list apps
apimctl list keys -a 'Avengers'
apimctl list oauths -a 'Avengers'
apimctl list apis
apimctl list proxies
apimctl list access -a 'Avengers'
apimctl list access -a 'X-Men'

apimctl describe org -n 'Marvel'
apimctl describe user -n 'Iron Man'
apimctl describe app -n Avengers
apimctl describe api -n 'Captain America'
apimctl describe proxy -n 'The First Avenger'

apimctl unpublish proxy -n 'The First Avenger'
apimctl unpublish proxy -n 'The Winter Soldier'
apimctl unpublish proxy -n 'Civil War'
apimctl unpublish proxy -n 'First Class'
apimctl unpublish proxy -n 'Days of the Future Past'
apimctl unpublish proxy -n 'Apocalypse'

apimctl delete proxy -n 'The First Avenger'
apimctl delete api -n 'Captain America'
apimctl delete app -n Avengers
apimctl delete user -n 'AntMan'
apimctl delete user -n Thor
apimctl delete user -n 'Iron Man'
apimctl delete org -n Marvel

