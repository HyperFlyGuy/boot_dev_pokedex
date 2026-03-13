
package pokeapi

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func (c *Client) ExploreRequest(url string) PokeExploreResponse {
	// Check the cache
	if val, ok := c.cache.Get(url); ok {
		data := PokeExploreResponse{}
		err := json.Unmarshal (val, &data)
		if err != nil {
			fmt.Println(err)
			return PokeExploreResponse{}
		}
		return data
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return PokeExploreResponse{}
	}
	res,err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return PokeExploreResponse{}
	}
	if err != nil {
		fmt.Println(err)
		return PokeExploreResponse{}
	}
	c.cache.Add(url,body)
	data := PokeExploreResponse{}
	err = json.Unmarshal (body, &data)
	if err != nil {
		fmt.Println(err)
		return PokeExploreResponse{}
	}
	return data
}
