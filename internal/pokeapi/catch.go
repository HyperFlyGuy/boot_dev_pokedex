package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchRequest(url string) Pokemon {
	// Check the cache
	if val, ok := c.cache.Get(url); ok {
		data := Pokemon{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			fmt.Println(err)
			return Pokemon{}
		}
		return data
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return Pokemon{}
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return Pokemon{}
	}
	if err != nil {
		fmt.Println(err)
		return Pokemon{}
	}
	c.cache.Add(url, body)
	data := Pokemon{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}
	}
	return data
}
