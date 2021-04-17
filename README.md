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
```

## Option

```
-p : Tha path to go.mod (default is "./go.mod")
```
