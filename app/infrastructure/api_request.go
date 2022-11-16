package infrastructure

import (
	"bytes"
	"encoding/json"

	"io"
	"net/http"

	//"golang.org/x/net/context"

	//"google.golang.org/grpc"

	//"github.com/Kantaro0829/clean-architecture-in-go/domain"
	//"github.com/Kantaro0829/clean-architecture-in-go/infrastructure/chat"
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

func (handler *ApiHandler) Post(reqBody interface{}, url string) ([]byte, error) {
	//後で引数として指定
	// requestBody := &domain.GetSubReq{
	// 	VideoId: "Z2Y0GMCFWq0",
	// }

	requestBody := reqBody
	endpoint := url
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}

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

	res, err := io.ReadAll(handler.response.Body)
	//json.Unmarshal(res, &responseBody)
	if err != nil {
		return nil, err
	}
	return res, nil

	// if err != nil {
	// 	return nil, err
	// }
	// return responseBody, nil

}

// type ApiHandler struct {
// 	conn *grpc.ClientConn
// }

// func (handler *ApiHandler) Post(body interface{}, url string) ([]byte, error) {

// 	resBody := domain.GetSubResp{}

// 	jsonString, err := json.Marshal(body)
// 	json.Unmarshal(jsonString, &resBody)
// 	conn, err := grpc.Dial("grpc-subtitle-service:5002", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %s", err)
// 	}
// 	defer conn.Close()

// 	c := chat.NewChatServiceClient(conn)
// 	response, err := c.SayHello(context.Background(), &chat.Message{Body: resBody.Body})

// 	if err != nil {
// 		log.Fatalf("Error when calling SayHello: %s", err)
// 	}

// 	getSubResp := domain.GetSubResp{}
// 	getSubResp.Body = response.Body
// 	obj, err := json.Marshal(getSubResp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return obj, err

// }
