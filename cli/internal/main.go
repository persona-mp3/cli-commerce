package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var endpoints = [5]string{"accessories", "outdoor", "jackets & coats", "sports wear", "artillery"}

func MainScreen() string {
	for i, item := range endpoints {
		fmt.Printf("%d, %s\n", i+1, item)
	}

	var choice int
	for {
		fmt.Scanf("%d", &choice)

		choice = choice - 1
		if choice < 0 || choice > len(endpoints) {
			fmt.Printf("select a valid option within range sir\n")
		} else {
			break
		}

	}
	return endpoints[choice]
}

func QueryServer(item string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	addr, err := url.Parse("http://localhost:8080/cli/products")
	if err != nil {
		return []byte{}, fmt.Errorf("error in parsing server_addr: %w ", err)
	}

	targetEndpoint, err := addr.Parse(fmt.Sprintf("/%s", item))
	if err != nil {
		return []byte{}, fmt.Errorf("error in parsing for endpoint: %w ", err)
	}

	client := &http.Client{}
	var body []byte
	reader := bytes.NewReader(body)

	req, err := http.NewRequestWithContext(ctx, "GET", targetEndpoint.String(), reader)
	if err != nil {
		return []byte{}, fmt.Errorf("error in creating new_request: %w ", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("error in making request: %w ", err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error in reading response_body: %w ", err)
	}

	log.Println("successfully queried server")
	return content, nil
}

func ParseResponse(response []byte) (string, error) {
	return "", nil
}
