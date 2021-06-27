# IITK-COIN

Updated after each task so that if you get to a commit it shows what were the things that were working at that time.

## Testing Work Flow 

### Client log
```bash
# Signup
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"name":"Abhay Mishra","password":"12345"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 27 Jun 2021 06:45:27 GMT
Content-Length: 49

{"rollno":190017,"name":"Abhay Mishra","coin":0}

# Login
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"password":"12345"}' http://localhost:8080/login
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjM5Nn0.mtDlUIopqy5sbEfmXCKDh-la7dq68y5a-XjLA1bRkok; Expires=Sun, 27 Jun 2021 06:46:36 GMT
Date: Sun, 27 Jun 2021 06:45:36 GMT
Content-Length: 49

{"rollno":190017,"name":"Abhay Mishra","coin":0}

# Wrong input to login form
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjM5Nn0.mtDlUIopqy5sbEfmXCKDh-la7dq68y5a-XjLA1bRkok" http://localhost:8080/login
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 27 Jun 2021 06:46:03 GMT
Content-Length: 26

invalid input to the form

# Request Secret page with cookie
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjM5Nn0.mtDlUIopqy5sbEfmXCKDh-la7dq68y5a-XjLA1bRkok" http://localhost:8080/secret
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 27 Jun 2021 06:46:10 GMT
Content-Length: 84

{"message":"We are so excited to meet you","user":"Abhay Mishra","from":"invictus"}

# Request Again
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjM5Nn0.mtDlUIopqy5sbEfmXCKDh-la7dq68y5a-XjLA1bRkok" http://localhost:8080/secret
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 27 Jun 2021 06:46:13 GMT
Content-Length: 84

{"message":"We are so excited to meet you","user":"Abhay Mishra","from":"invictus"}

# Token Expired
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjM5Nn0.mtDlUIopqy5sbEfmXCKDh-la7dq68y5a-XjLA1bRkok" http://localhost:8080/secret
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 27 Jun 2021 06:46:45 GMT
Content-Length: 14

Unknown Error

# Login Again
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"password":"12345"}' http://localhost:8080/login
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjQ4MH0.gkmrg3SbQsfAt5xBlUW-3xR6QnfyN3BbB5qOvE9VbR8; Expires=Sun, 27 Jun 2021 06:48:00 GMT
Date: Sun, 27 Jun 2021 06:47:00 GMT
Content-Length: 49

{"rollno":190017,"name":"Abhay Mishra","coin":0}
# Request Secret Page again
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywibmFtZSI6IkFiaGF5IE1pc2hyYSIsImV4cCI6MTYyNDc3NjQ4MH0.gkmrg3SbQsfAt5xBlUW-3xR6QnfyN3BbB5qOvE9VbR8" http://localhost:8080/secret
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 27 Jun 2021 06:47:35 GMT
Content-Length: 84

{"message":"We are so excited to meet you","user":"Abhay Mishra","from":"invictus"}

# Done
```

## Server Log

```bash
:~/go/src/github.com/mabhay3420/iitk-coin$ go run .
2021/06/27 12:15:22 tables ready.
Starting Server at http://localhost:8080
2021/06/27 12:15:26 add New User....
2021/06/27 12:15:27 succesfully Added New User.
2021/06/27 12:15:27 User: Rollno: 190017 Name: Abhay Mishra
Login of Abhay Mishra Succesful
EOF
Secret content Delievered to  Abhay Mishra
Secret content Delievered to  Abhay Mishra
token is expired by 9s
Login of Abhay Mishra Succesful
Secret content Delievered to  Abhay Mishra
^Csignal: interrupt
```
