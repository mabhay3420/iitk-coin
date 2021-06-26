# IITK-COIN

Updated after each task so that if you get to a commit it shows what were the things that were working at that time.

## Testing Work Flow 

### Client log
```bash
# Signup
:$ curl -i -X POST -H "Content-Type: application/json" -d '{"rollno":190017,"name":"Abhay Mishra","password":"12345"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 26 Jun 2021 06:32:00 GMT
Content-Length: 49

{"rollno":190017,"name":"Abhay Mishra","coin":0}
# Succesful Login
:$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"password":"12345"}' http://localhost:8080/login
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywiZXhwIjoxNjI0Njg5NDY0fQ.nmErKRh0Ux62gVWJ8znuCT2unAkpXXcxHTg6IJ4LSu4; Expires=Sat, 26 Jun 2021 06:37:44 GMT
Date: Sat, 26 Jun 2021 06:32:44 GMT
Content-Length: 49

{"rollno":190017,"name":"Abhay Mishra","coin":0}

# Signup with same name and different roll no.
:$ curl -i -X POST -H "Content-Type: application/json" -d '{"rollno":190195,"name":"Abhay Mishra","password":"123a5"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 26 Jun 2021 06:34:33 GMT
Content-Length: 49

{"rollno":190195,"name":"Abhay Mishra","coin":0}

# Roll no is not an integer.
:$ curl -i -X POST -H "Content-Type: application/json" -d '{"rollno":190a95,"name":"Abhay Mishra","password":"123a5"}' http://localhost:8080/signup
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 26 Jun 2021 06:34:44 GMT
Content-Length: 22

invalid input to form
```

## Server Log

```bash
:~/go/src/github.com/mabhay3420/iitk-coin$ go run .
2021/06/26 12:00:52 tables ready.
Starting Server at http://localhost:8080
2021/06/26 12:01:59 add New User....
2021/06/26 12:02:00 succesfully Added New User.
2021/06/26 12:02:00 User: Rollno: 190017 Name: Abhay Mishra
Login of Abhay Mishra Succesful
2021/06/26 12:04:32 add New User....
2021/06/26 12:04:33 succesfully Added New User.
2021/06/26 12:04:33 User: Rollno: 190017 Name: Abhay Mishra
2021/06/26 12:04:33 User: Rollno: 190195 Name: Abhay Mishra
invalid character 'a' after object key:value pair
^Csignal: interrupt
```
