**Copyright (c) 2018 BitMart Exchange. All Rights Reserved.**

## BitMart Exchange Golang Sample API

### How to use it

Get the `bm_client.go` under your GOPATH.
Please check function `TestBMAccessToken()` in `bm_client_test.go` to see how to use it.

### Note:
- RSA private key format: PKCS8
- TimeStamp value should be in milliseconds, while normal UNIX timestamps are usually in seconds.
- Read [bitmart docs](https://github.com/bitmart-docs/bitmart-api-reference/blob/master/REST.md) if need more details
