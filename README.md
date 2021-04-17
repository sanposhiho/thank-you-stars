# thank-you-stars

Give stars to your dependencies of Go repositories and say thank you to developers!!


inspired by
- https://github.com/y-taka-23/thank-you-stars
- https://github.com/teppeis/thank-you-stars

## Set up

### Go version < 1.16

```
GO111MODULE=on go get github.com/sanposhiho/thank-you-stars
```

### Go version 1.16+

```
go install github.com/sanposhiho/thank-you-stars@latest
```

## How to use

You have to generate GitHub Auth Token.

Go https://github.com/settings/tokens and generate token.
This repository only needs the scope of public_repo.

```
$ export GITHUB_AUTH_TOKEN=your_github_auth_token
$ thank-you-stars
2021/04/17 15:09:49 Starred! github.com/go-gorp/gorp/v3 
2021/04/17 15:09:50 Starred! github.com/go-sql-driver/mysql 
2021/04/17 15:09:50 Starred! github.com/golang/mock 
2021/04/17 15:09:50 Starred! github.com/golang/protobuf 
2021/04/17 15:09:51 Starred! github.com/google/go-cmp 
2021/04/17 15:09:51 Starred! github.com/google/uuid 
2021/04/17 15:09:51 Starred! github.com/gorilla/websocket 
2021/04/17 15:09:52 Starred! github.com/labstack/echo/v4 
2021/04/17 15:09:52 Starred! github.com/labstack/gommon 
2021/04/17 15:09:52 Starred! github.com/patrickmn/go-cache 
2021/04/17 15:09:52 Starred! github.com/zmb3/spotify 
```

## Option

```
-p : Tha path to go.mod (default is "./go.mod")
```
