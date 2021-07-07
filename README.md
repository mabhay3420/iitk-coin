# IITK-COIN

Updated after each task so that if you get to a commit it shows what were the things that were working at that time.


## Testing Doc Specifications

### Client log
```bash
# Student 1 signup
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190195,"name":"Ashok kumar Saini","password":"6789"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 07 Jul 2021 15:18:12 GMT
Content-Length: 54

{"rollno":190195,"name":"Ashok kumar Saini","coin":0}

# student 2 signup
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"name":"Abhay Mishrai","password":"12345"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 07 Jul 2021 15:18:26 GMT
Content-Length: 50

{"rollno":190017,"name":"Abhay Mishrai","coin":0}

# admin 1 signup
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190058,"name":"Gaurav Sharma","password":"abcd"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 07 Jul 2021 15:19:05 GMT
Content-Length: 50

{"rollno":190058,"name":"Gaurav Sharma","coin":0}

# Council core member signup
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190345,"name":"Aman Agrawal","password":"efgh"}' http://localhost:8080/signup
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 07 Jul 2021 15:19:25 GMT
Content-Length: 49

{"rollno":190345,"name":"Aman Agrawal","coin":0}

# Council core member login
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190345,"password":"efgh"}' http://localhost:8080/loginHTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDM0NSwiZXhwIjoxNjI1NjcxNTAxfQ.J17IGyG6Wb6Wb4xpcr_a2gALh-YTSOgv89ShFKQZ9fk; Expires=Wed, 07 Jul 2021 15:25:01 GMT
Date: Wed, 07 Jul 2021 15:20:01 GMT
Content-Length: 49

{"rollno":190345,"name":"Aman Agrawal","coin":0}

# Only admin allowed to reward others
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"award":100}' --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDM0NSwiZXhwIjoxNjI1NjcxNTAxfQ.J17IGyG6Wb6Wb4xpcr_a2gALh-YTSOgv89ShFKQZ9fk" http://localhost:8080/award
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:20:59 GMT
Content-Length: 36

You need admin role to award others

# Admin login
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190058,"password":"abcd"}' http://localhost:8080/loginHTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxNjAzfQ.v5aRfEuTHHp0bTEWFbkD7MV2CCITugisE8AfF1rLGUY; Expires=Wed, 07 Jul 2021 15:26:43 GMT
Date: Wed, 07 Jul 2021 15:21:43 GMT
Content-Length: 50

{"rollno":190058,"name":"Gaurav Sharma","coin":0}

# Admin rewards a student succesfully
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"award":100}' --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxNjAzfQ.v5aRfEuTHHp0bTEWFbkD7MV2CCITugisE8AfF1rLGUY" http://localhost:8080/award
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 07 Jul 2021 15:22:10 GMT
Content-Length: 30

{"rollno":190017,"award":100}

# Admin reward Council Core member : perfectly valid
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190345,"award":200}' --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxNjAzfQ.v5aRfEuTHHp0bTEWFbkD7MV2CCITugisE8AfF1rLGUY" http://localhost:8080/award
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 07 Jul 2021 15:22:47 GMT
Content-Length: 30

{"rollno":190345,"award":200}

# Admin cannot be involved in transfers and awards
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190058,"award":300}' --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxNjAzfQ.v5aRfEuTHHp0bTEWFbkD7MV2CCITugisE8AfF1rLGUY" http://localhost:8080/award
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:23:11 GMT
Content-Length: 43

No of coins in admin account cannot change

# Student login : Awarded succesfully
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190017,"password":"12345"}' http://localhost:8080/login
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywiZXhwIjoxNjI1NjcxNzQ3fQ.wTuwBuO_i9SP1DepQs66sYl44hop7C7k5wICZw7kvoM; Expires=Wed, 07 Jul 2021 15:29:07 GMT
Date: Wed, 07 Jul 2021 15:24:07 GMT
Content-Length: 52

{"rollno":190017,"name":"Abhay Mishrai","coin":100}

# JSON data not in proper format
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190195,"amount"300}' http://localhost:8080/award --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywiZXhwIjoxNjI1NjcxNzQ3fQ.wTuwBuO_i9SP1DepQs66sYl44hop7C7k5wICZw7kvoM"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:25:25 GMT
Content-Length: 26

Invalid Input to the form

# No award field : assumes 0 : Non - positive award
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190195,"amount":300}' http://localhost:8080/award --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywiZXhwIjoxNjI1NjcxNzQ3fQ.wTuwBuO_i9SP1DepQs66sYl44hop7C7k5wICZw7kvoM"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:25:43 GMT
Content-Length: 52

No of coins to be awarded must be a positive number

# Only Admin can award others
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190195,"award":300}' http://localhost:8080/award --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDAxNywiZXhwIjoxNjI1NjcxNzQ3fQ.wTuwBuO_i9SP1DepQs66sYl44hop7C7k5wICZw7kvoM"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:26:19 GMT
Content-Length: 36

You need admin role to award others

#Admin login
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190058,"password":"abcd"}' http://localhost:8080/login
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxOTAzfQ.WI6dQPhFBKx-nlDYMMI_upsXgM7uxNgX-bx8mSatsDg; Expires=Wed, 07 Jul 2021 15:31:43 GMT
Date: Wed, 07 Jul 2021 15:26:43 GMT
Content-Length: 50

{"rollno":190058,"name":"Gaurav Sharma","coin":0}

# A universal cap on no of coins user can have : 10000 here
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"rollno":190195,"award":1000000}' http://localhost:8080/award --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxOTAzfQ.WI6dQPhFBKx-nlDYMMI_upsXgM7uxNgX-bx8mSatsDg"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:27:58 GMT
Content-Length: 34

Balance limit reached for awardee

# No award field
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"from":190345,"to":190195,"amount":100}' http://localhost:8080/award --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxOTAzfQ.WI6dQPhFBKx-nlDYMMI_upsXgM7uxNgX-bx8mSatsDg"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:29:05 GMT
Content-Length: 52

No of coins to be awarded must be a positive number

# Only user can transfer from there account
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"from":190345,"to":190195,"amount":100}' http://localhost:8080/transfer --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxOTAzfQ.WI6dQPhFBKx-nlDYMMI_upsXgM7uxNgX-bx8mSatsDg"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:29:24 GMT
Content-Length: 40

You can transfer from your account only

# Must participate in pre-decided minimum no of activities in order to exchange coins
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -X POST -H "Content-Type: application/json" -i -d '{"from":190058,"to":190195,"amount":100}' http://localhost:8080/transfer --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOjE5MDA1OCwiZXhwIjoxNjI1NjcxOTAzfQ.WI6dQPhFBKx-nlDYMMI_upsXgM7uxNgX-bx8mSatsDg"
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 07 Jul 2021 15:29:38 GMT
Content-Length: 57

You are not eligible to transfer coins, Participate More

# Need to automate for further testing
:~/go/src/github.com/mabhay3420/iitk-coin$
```

## Server Log

```bash

# Initial Log
:~/go/src/github.com/mabhay3420/iitk-coin$ go run .
2021/07/07 20:48:08 tables ready.
===-------User List-------===
===-------Award History------===
===-------Transfer History------===
===----------Here you go--------===
Starting Server at http://localhost:8080
2021/07/07 20:48:11 add New User....
2021/07/07 20:48:12 succesfully Added New User.
2021/07/07 20:48:25 add New User....
2021/07/07 20:48:26 succesfully Added New User.
2021/07/07 20:49:04 add New User....
2021/07/07 20:49:05 succesfully Added New User.
2021/07/07 20:49:24 add New User....
2021/07/07 20:49:25 succesfully Added New User.
Login of Aman Agrawal Succesful
190345 tried to award 190017
Login of Gaurav Sharma Succesful
100 Coins awarded to 190017
200 Coins awarded to 190345
190058 who has admin access tried to award himself.Aborted!
Login of Abhay Mishrai Succesful
invalid character '3' after object key
Non-positive Award Requested. Aborted the process
190017 tried to award 190195
Login of Gaurav Sharma Succesful
Ashok kumar Saini is full of coin
Non-positive Award Requested. Aborted the process
190058 tried to send from 190345 Account. Aborted!
190058 Has participated in only 0 Activities. Unable to transfer!
^Csignal: interrupt

# Updated tables
:~/go/src/github.com/mabhay3420/iitk-coin$ go run .
2021/07/07 21:00:03 tables ready.
===-------User List-------===
# Rollno Name Coins Password(hashed form) Role Activities(+1 when awarded)
{190017 Abhay Mishrai 100 $2a$14$2SUfwKsURP7pjYkqgq8V4u0jvFNqHF8Fs3b61ll7/j3hMGg3PfRZ6 STUDENT 1}
{190345 Aman Agrawal 200 $2a$14$Ql.NLt4Ikx73OCO7z/V7D.Q2a9BL8z9IbYGeNpunKXiPnvhDM5nd6 COUNCIL_CORE 1}
{190195 Ashok kumar Saini 0 $2a$14$bL4LEpRf345zL5Y8DtQOtO5tunGKBWLkOxc92F6WvgneM.kE4JyFa STUDENT 0}
{190058 Gaurav Sharma 0 $2a$14$2CBAqzSDytenKaoMkoQyJOeeL9BGvhbk8eTikiZmjfwmJWVKHVpb. ADMIN 0}
===-------Award History------===
# timeStamp awardeeRollno award
2021-07-07 20:52:10.231986432 +0530 +0530 {190017 100}
2021-07-07 20:52:47.61959543 +0530 +0530 {190345 200}
===-------Transfer History------===
# timeStamp SenderRollno RecieverRollno Amount
===----------Here you go--------===
Starting Server at http://localhost:8080
^Csignal: interrupt
```
