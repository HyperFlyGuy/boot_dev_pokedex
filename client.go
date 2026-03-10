package main 

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type PokeResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`

}

func PokeRequest(url string) PokeResponse {
	res,err := http.Get(url)
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
	data := PokeResponse{}
	err = json.Unmarshal (body, &data)
	if err != nil {
		fmt.Println(err)
		return PokeResponse{}
	}
	return data
}
