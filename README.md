# wasedatime-backend
A Restful API for [WasedaTime](https://wasedatime.com/) written in golang

| HTTP Request  | Request Parameter  | Description                      |
| ------------- | ------------------ | -------------------------------- |
| GET /comments | `string` course_id | get all the comments of a course |

## Prerequisites

go v 1.14.4

## Build

```bash
cd app
GOOS=linux
GOARCH=amd64
go build
```

## Deploy

See https://kenyaappexperts.com/blog/how-to-deploy-golang-to-production-step-by-step/