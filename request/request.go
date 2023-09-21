package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	customError "github.com/Inalegwu/babel/error"
)

type ColorApiResponse struct {
	Hex  map[string]string `json:"hex"`
	Name map[string]any    `json:"name"`
}

func GetColorCodeName(colorCode string) ColorApiResponse {
	resp, err := http.Get("https://www.thecolorapi.com/id?hex=" + colorCode)

	customError.HandleError(err)

	var response ColorApiResponse

	body, err := ioutil.ReadAll(resp.Body)

	customError.HandleError(err)

	err = json.Unmarshal(body, &response)

	return response
}
