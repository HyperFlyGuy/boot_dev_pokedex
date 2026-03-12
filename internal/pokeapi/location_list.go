package pokeapi

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func (c *Client) LocationAreaRequest(url string) PokeResponse {
	// Check the cache
	if val, ok := c.cache.Get(url); ok {
		data := PokeResponse{}
		err := json.Unmarshal (val, &data)
		if err != nil {
			fmt.Println(err)
			return PokeResponse{}
		}
		return data
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return PokeResponse{}
	}
	res,err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return PokeResponse{}
	}
	if err != nil {
		fmt.Println(err)
		return PokeResponse{}
	}
	c.cache.Add(url,body)
	data := PokeResponse{}
	err = json.Unmarshal (body, &data)
	if err != nil {
		fmt.Println(err)
		return PokeResponse{}
	}
	return data
}
