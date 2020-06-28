# wasedatime-backend

Branch | Build Status
--- | --- 
master | staging
[![Build Status](https://travis-ci.com/wasedatime/wasedatime-backend.svg?branch=master)](https://travis-ci.com/wasedatime/wasedatime-backend) | [![Build Status](https://travis-ci.com/wasedatime/wasedatime-backend.svg?branch=staging)](https://travis-ci.com/wasedatime/wasedatime-backend)

A Restful API for [WasedaTime](https://wasedatime.com/) written in golang

| HTTP Request  | Request Parameter  | Description                      |
| ------------- | ------------------ | -------------------------------- |
| GET /comments | `string` course_id | get all the comments of a course |

## Prerequisites

go v 1.14.4

## Build

```bash
cd app
env GOOS=linux GOARCH=amd64 go build
```

## Deploy and daemonitization

See https://astaxie.gitbooks.io/build-web-application-with-golang/de/12.3.html
