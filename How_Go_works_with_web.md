# 3.3 How Go works with web
### Web principle
Request: POST, GET, Cookie, URL 등의 사용자에게 얻어내는 요청 데이터  
Response: 서버에서 클라이언트에 전달하는 응답 데이터  
Conn: 클라이언트 서버 사이의 연결  
Handler: Request 처리 로직과 response 를 생성해내는 것


### http package operating mechanism
1. listening socket 생성, 특정 포트에 대기하며 클라언트를 기다린다.
2. 클라이언트의 요청을 받는다.
3. Request 를 처리. HTTP 헤더를 읽는다. 만약 POST method 라면 message body 의 데이터를 읽어 handler 에 전달한다. 이후에 소켓은 response data 를 클라이언트에게 전달한다.
