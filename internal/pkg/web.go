package pkg

import (
	"fmt"
	"io"
	"net/http"
)

const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.78"

// GetRequest send a get request
func GetRequest(url string, queryList [][]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	q := req.URL.Query()
	for _, queryItem := range queryList {
		if len(queryItem) != 2 {
			return nil, fmt.Errorf("%v is not a valid query", queryItem)
		}
		q.Add(queryItem[0], queryItem[1])
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
