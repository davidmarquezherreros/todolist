

# Backend
Golang
## Retrieve all tasks
GET /task
{
  "HttpCode": 200,
  "Message": "OK",
  "Response": [
    {
      "ID": 1,
      "Name": "TaskName",
      "Status": 1,
      "StartDate": "20/07/2018",
      "EndDate": ""
    },
    {
      "ID": 2,
      "Name": "TaskName",
      "Status": 0,
      "StartDate": "27/07/2018",
      "EndDate": ""
    },
    {
      "ID": 3,
      "Name": "TaskName",
      "Status": 2,
      "StartDate": "20/07/2018",
      "EndDate": "21/07/2018"
    }
  ]
}

## Retrieve a task
GET /task/{id}

Result:
{
  "HttpCode": 200,
  "Message": "OK",
  "Response": {
    "ID": 3,
    "Name": "Name of the task",
    "Status": 1,
    "StartDate": "20/07/2018",
    "EndDate": "21/07/2018"
  }
}
## Create a task
POST /task
Body:
{
  "Name": "TaskName",
  "Status": 2,
  "StartDate": "20/07/2018",
  "EndDate": "21/07/2018"
}
Result:
{
  "HttpCode": 201,
  "Message": "Created",
  "Response": {
    "ID": 4,
    "Name": "TaskName",
    "Status": 2,
    "StartDate": "20/07/2018",
    "EndDate": "21/07/2018"
  }
}
## Update a task
PUT /task/{id}
Body:
{
  "Name": "TaskName",
  "Status": 2,
  "StartDate": "20/07/2018",
  "EndDate": "21/07/2018"
}
Result:
{
  "HttpCode": 200,
  "Message": "Created",
  "Response": {
    "ID": 4,
    "Name": "TaskName",
    "Status": 2,
    "StartDate": "20/07/2018",
    "EndDate": "21/07/2018"
  }
}
## Delete a task
DELETE /task/{id}
Result:
{
  "HttpCode": 200,
  "Message": "OK",
  "Response": "Task deleted"
}


# Run program
go build
./backend
