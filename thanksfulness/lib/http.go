package lib

import (
	"log"
	"net/http"
)

func AddHeaders(current http.Header, input map[string][]string) (header http.Header) {
	header = current.Clone()
	for k, v := range input {
		for _, str := range v {
			header.Add(k, str)
		}
	}
	return header
}

// https://pkg.go.dev/net/http#NewRequest
// https://zenn.dev/not75743/scraps/9c48c22763083e
func Get(url string, headers map[string][]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req.Header = AddHeaders(req.Header, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch to %s.: %v", url, err)
		return nil, err
	}
	return resp, nil
}
