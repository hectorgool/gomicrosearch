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
docker pull elasticsearch
docker run -d --name elasticsearch elasticsearch
docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch
```

### Referencias:

##### go-kit 
http://gokit.io/
