package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/getsubtitle"
)

type ApiHandler struct {
	response *http.Response
}

func NewApiHandler() getsubtitle.ApiHandler {

	apiHandler := new(ApiHandler)
	// rakutenApiHandler.request, _ = goquery.NewDocumentFromReader(res.Body)
	return apiHandler
}

func (handler *ApiHandler) Post(body interface{}, url string) ([]byte, error) {
	//後で引数として指定
	// requestBody := &domain.GetSubReq{
	// 	VideoId: "Z2Y0GMCFWq0",
	// }
	requestBody := body
	endpoint := url
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}
	fmt.Printf("json: %v", jsonString)

	//後で引数として指定
	//endpoint := "http://get-subtitle-service:5001/get_subtittle"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	handler.response = resp
	defer resp.Body.Close()
	handler.response.Body = resp.Body

	resBody, err := io.ReadAll(handler.response.Body)

	if err != nil {
		return nil, err
	}
	return resBody, nil

}
