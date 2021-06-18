# IITK-COIN


## Signup Error Handling

### Client log
```bash
# Invalid Roll no
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -i -d "rollno=190a95&name=Ashok%20kumar%20%Saini" http://localhost:8080/signup
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Fri, 18 Jun 2021 15:25:19 GMT
Content-Length: 22

Invalid input to Form

# Invalid name
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -i -d "rollno=190195&name=Ashok%20kumar%20%Saini" http://localhost:8080/signup
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Fri, 18 Jun 2021 15:25:41 GMT
Content-Length: 22

Invalid input to Form

# Proper Post request
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -i -d "rollno=190195&name=Ashok%20kumar%20Saini" http://localhost:8080/signup
HTTP/1.1 200 OK
Date: Fri, 18 Jun 2021 15:25:56 GMT
Content-Length: 94
Content-Type: text/plain; charset=utf-8

POST Request Succesful.
User Succesfully Added
Your Rollno:190195 Your Name:Ashok kumar Saini

# Proper request
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -i -d "rollno=190017&name=Abhay%20Mishra" http://localhost:8080/signup
HTTP/1.1 200 OK
Date: Fri, 18 Jun 2021 15:26:50 GMT
Content-Length: 89
Content-Type: text/plain; charset=utf-8

POST Request Succesful.
User Succesfully Added
Your Rollno:190017 Your Name:Abhay Mishra

# User Already Exists
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -i -d "rollno=190017&name=Abhay%20Mishra" http://localhost:8080/signup
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Fri, 18 Jun 2021 15:26:56 GMT
Content-Length: 50

Invalid User Name/Password or User Already Exists

# Checks for Roll number only
:~/go/src/github.com/mabhay3420/iitk-coin$ curl -i -d "rollno=190017&name=Abhay" http://localhost:8080/signup
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Fri, 18 Jun 2021 15:27:27 GMT
Content-Length: 50

Invalid User Name/Password or User Already Exists
```

## Server Log

```bash
:~/go/src/github.com/mabhay3420/iitk-coin$ go run .
Create Student table....
Student table Created Succesfully.
Starting Server at http://localhost:8080
invalid URL escape "%Sa"
invalid URL escape "%Sa"
Add New User....
Succesfully Added New User.
RollNo: 190195 Name: Ashok kumar Saini
Add New User....
Succesfully Added New User.
RollNo: 190017 Name: Abhay Mishra
RollNo: 190195 Name: Ashok kumar Saini
Add New User....
Unable to Add user
UNIQUE constraint failed: students.rollno
Add New User....
Unable to Add user
UNIQUE constraint failed: students.rollno
^Csignal: interrupt
```
