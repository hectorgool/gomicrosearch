

curl -X DELETE `docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch`':9200/mx?pretty'


curl -X PUT `docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch`':9200/mx?pretty' -d '
{
    "settings": {
        "number_of_shards": 1, 
        "analysis": {
            "filter": {
                "autocomplete_filter": { 
                    "type":     "edge_ngram",
                    "min_gram": 1,
                    "max_gram": 20
                }
            },
            "analyzer": {
                "autocomplete": {
                    "type":      "custom",
                    "tokenizer": "standard",
                    "filter": [
                        "lowercase",
                        "autocomplete_filter" 
                    ]
                }
            }
        }
    },
    "mappings": {
        "postal_code": {
            "properties": {
                "cp": {
                    "type": "multi_field",
                    "fields": {
                        "cp": {
                            "type" : "float", 
                            "store" : "yes", 
                            "precision_step" : "4"
                        }
                    }
                },
                "colonia": {
                    "type": "multi_field",
                    "fields": {
                        "colonia": {
                            "type": "string",
                            "index": "not_analyzed",
                            "store": "yes"
                        }
                    }
                },                 
                "ciudad": {
                    "type": "multi_field",
                    "fields": {
                        "ciudad": {
                            "type": "string",
                            "index": "not_analyzed",
                            "store": "yes"
                        }
                    }
                },                                
                "delegacion": {
                    "type": "multi_field",
                    "fields": {
                        "delegacion": {
                            "type": "string",
                            "index": "not_analyzed",
                            "store": "yes"
                        }
                    }
                },    
                "location": {
                    "type": "geo_point"
                }
            }
        }
    }
}
'
