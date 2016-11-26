export ELASTICSEARCH_HOSTS=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch`:9200
echo "ELASTICSEARCH_HOSTS: $ELASTICSEARCH_HOSTS"