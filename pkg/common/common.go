package common

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// translate map to string
func MapToString(mapData map[string]interface{}) string {
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

// Base64 Encoding
func Base64Encode(inputString string) string {
	inputByte := []byte(inputString)
	encodedString := base64.StdEncoding.EncodeToString(inputByte)

	return encodedString
}

// HTTP 메시지를 문자 형식으로 읽어서 HTTP 요청을 전송하고, Go HTTP Response 객체를 반환합니다.
func HttpMessageToGoResp(message string) *http.Response {
	reader := bufio.NewReader(strings.NewReader(message))
	request, err := http.ReadRequest(reader)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	client := &http.Client{}
	resp, respErr := client.Do(request)

	if respErr != nil {
		log.Fatal(respErr)
		return nil
	}

	return resp
}
