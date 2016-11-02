# GoMicroSearch

##### 1. Workspace:
#
```sh
cd $GOPATH/src/github.com/hectorgool/gomicrosearch 
```
##### 2. Compile:
#
```sh
go build
```

### commit y push
```sh
git remote add origin https://github.com/hectorgool/gomicrosearch.git
git push -u origin master
```

### ElastiSearch Docker Container
```sh
docker pull elasticsearch:2.3
docker run -d --name elasticsearch elasticsearch:2.3
docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch
```

### Referencias:

##### go-kit 
http://gokit.io/

##### get docker ip container 
https://gist.github.com/hectorgool/c01ed842f75297c011591b1af92dec55

##### ElasticSearch Client go example
https://gist.github.com/hectorgool/78986b1684a35dc8e8ecbe744f524553

##### aws ElasticSearch support version
https://aws.amazon.com/es/elasticsearch-service/faqs/