package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GoogleMapSearchContextInterface interface {
	JSON(code int, obj interface{})
	GetPostForm(key string) (string, bool)
}

type GoogleMapSearchInterface interface {
	PostGoogleMapSearch(contexts GoogleMapSearchContextInterface, request *http.Request)
	GetLatLong(address string, keyInput string) (latlong string)
}

type GoogleMapSearchController struct{}

func (controller *GoogleMapSearchController) PostGoogleMapSearch(contexts GoogleMapSearchContextInterface, request *http.Request) {
	dataInput, _ := contexts.GetPostForm("address")
	keyInput, _ := contexts.GetPostForm("key")
	latlong := controller.GetLatLong(dataInput, keyInput)
	resp, err := http.Get("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latlong + "&rankby=distance&type=restaurant&key=" + keyInput)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	results := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&results)
	contexts.JSON(200, results["results"])
}

func (controller *GoogleMapSearchController) GetLatLong(address string, keyInput string) (latlong string) {
	resp, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + address + "&key=" + keyInput)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	results := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&results)
	if results["status"] == "OK" {
		for _, val := range results {
			switch valt := val.(type) {
			case []interface{}:
				for _, val2 := range valt {
					switch valt2 := val2.(type) {
					case map[string]interface{}:
						for key, val3 := range valt2 {
							if key == "geometry" {
								switch valt3 := val3.(type) {
								case map[string]interface{}:
									for key, val4 := range valt3 {
										if key == "location" {
											switch valt4 := val4.(type) {
											case map[string]interface{}:
												latByte, _ := json.Marshal(valt4["lat"])
												longByte, _ := json.Marshal(valt4["lng"])
												latlong = string(latByte) + "," + string(longByte)
											default:
												// i isn't one of the types above
											}
										}
									}
								default:
									// i isn't one of the types above
								}
							}
						}
					default:
						// i isn't one of the types above
					}
				}
			default:
				// i isn't one of the types above
			}
		}
	}
	return latlong
}
