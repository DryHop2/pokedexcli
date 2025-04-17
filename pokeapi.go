package main

type LocationAreaResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []LocationAreaInfo `json:"results"`
}

type LocationAreaInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
