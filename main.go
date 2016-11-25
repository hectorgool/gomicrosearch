/*
cd $GOPATH/src/github.com/hectorgool/gomicrosearch 
go get gopkg.in/olivere/elastic.v3
go build
*/

package main

import (
  "github.com/hectorgool/gomicrosearch/elasticsearch"
  "encoding/json"
  "errors"
  "log"
  "net/http"
  "golang.org/x/net/context"
  "github.com/go-kit/kit/endpoint"
  httptransport "github.com/go-kit/kit/transport/http"
)

// StringService provides operations on strings.
type SearchService interface {
  Term(string) (string, error)
}

type searchService struct{}

func (searchService) Term(term string) (string, error) {
  
  if term == "" {
    return "", ErrEmpty
  }
  
  result, err := elasticsearch.SearchTerm(term)
  if err != nil {
    return "", err
  }
  return result, nil

}

func main() {

  ctx := context.Background()
  svc := searchService{}

  searchHandler := httptransport.NewServer(
    ctx,
    makeSearchEndpoint(svc),
    decodeSearchRequest,
    encodeResponse,
  )

  http.Handle("/search", searchHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeSearchEndpoint(svc SearchService) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    req := request.(searchRequest)
    v, err := svc.Term(req.S)
    if err != nil {
      return searchResponse{v, err.Error()}, nil
    }
    return searchResponse{v, ""}, nil

  }
}

func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request searchRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    return nil, err
  }
  return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}

type searchRequest struct {
  S string `json:"s"`
}

type searchResponse struct {
  V   string `json:"v"`
  Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
