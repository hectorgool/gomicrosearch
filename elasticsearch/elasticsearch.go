/*
If you want to export variables, constants, and functions to be used with other
programs, the name of the identifier must start with an uppercase letter.
*/

package elasticsearch

import (
	"encoding/json"
	j "github.com/ricardolonga/jsongo"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"os"
)

var client *elastic.Client

func init() {

	var err error
	client, err = elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_HOSTS")))
	if err != nil {
		panic(err)
	}

}

func SearchTerm(term string) (string, error) {

	searchJson := j.Object().
		Put("size", 10).
		Put("query", j.Object().
			Put("match", j.Object().
				Put("_all", j.Object().
					Put("query", term).
					Put("operator", "and")))).
		Put("sort", j.Array().
			Put(j.Object().
				Put("colonia", j.Object().
					Put("order", "asc").
					Put("mode", "avg"))))

	searchResult, err := client.Search().
		Index("mx").
		Type("postal_code").
		Source(searchJson).
		Do()
	if err != nil {
		panic(err)
	}

	var documents []Document
	for _, hit := range searchResult.Hits.Hits {
		var d Document
		err := json.Unmarshal(*hit.Source, &d)
		if err != nil {
			log.Fatal(err)
		}
		documents = append(documents, d)
	}

	rawJsonDocuments, err := json.Marshal(documents)
	if err != nil {
		log.Fatal(err)
	}

	return string(rawJsonDocuments), nil

}

type Document struct {
	Ciudad     string `json:"ciudad"`
	Colonia    string `json:"colonia"`
	Cp         string `json:"cp"`
	Delegacion string `json:"delegacion"`
	Location   struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"location"`
}
