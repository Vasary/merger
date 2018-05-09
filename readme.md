# Persons merger
[![Build Status](https://travis-ci.org/Vasary/merger.svg?branch=master)](https://travis-ci.org/Vasary/merger)
[![Maintainability](https://api.codeclimate.com/v1/badges/8744e09d7127e1a6ac45/maintainability)](https://codeclimate.com/github/Vasary/merger/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/8744e09d7127e1a6ac45/test_coverage)](https://codeclimate.com/github/Vasary/merger/test_coverage)

> The little helper for UserData project which merges duplicates into one entity

#### Prepare
```
# cp .env.dist .env
```

#### Environment parameters

- DATABASE_USER=
- DATABASE_PASSWORD=
- DATABASE_NAME=
- DATABASE_HOST=
- DATABASE_PORT=

#### Run tests
```
# go test -v ./...
```

#### Build application
```
# cd src
# go build -o ../dist/app
```

#### Run application
```
# docker-compose up -d
```

#### View logs
```
# tail -fn 100 /app/logs/logs.log
```