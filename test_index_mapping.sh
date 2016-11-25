
curl -XGET `docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch`':9200/mx?pretty'
curl -XGET `docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch`':9200/mx/postal_code/_search?pretty'

