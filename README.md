# GoMicroSearch

##### 1. Workspace:
#
```sh
cd $GOPATH/src/github.com/hectorgool/gomicrosearch 
```
##### 2. Get dependencies:
#
```sh
go get ./...
```

##### 3. Compile:
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

### Create Docker Image
```sh
docker build -t gomicrosearch:0 .
```

### App Docker Container
```sh
docker run -d -p 8080:8080 --name gomicrosearch -e ELASTICSEARCH_HOSTS=elasticsearch:9200 --restart=always --link elasticsearch gomicrosearch:0 
```

### Get ip Docker Container
```sh
docker inspect --format '{{ .NetworkSettings.IPAddress }}' gomicrosearch
```

### Test Docker Container
```sh
curl -H "Content-Type: application/json" -X POST -d '{"term":"villa"}' http://`docker inspect --format '{{ .NetworkSettings.IPAddress }}' gomicrosearch`/search
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