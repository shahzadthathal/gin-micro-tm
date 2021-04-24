1. ## Getting Started

2. git clone https://github.com/shahzadthathal/gin-micro-tm.git

3. cd  gin-micro-tm.git

4. Configure your credential in config/db.go


6. go build

7. go run gin-micro-tm 

8. Get All Task
```
http://localhost:8080/api/task/
Method GET
```

9. Add new task
```
http://localhost:8080/api/task/
Method POST
POST Body
{
	"title":"Second task",
	"body":"Second task body"
}
```

9.  Browse Swagger UI [http://localhost:8999/swagger/index.html](http://localhost:8999/swagger/index.html)

## Admin credentials
```
username: admin
password: admin1234
```
