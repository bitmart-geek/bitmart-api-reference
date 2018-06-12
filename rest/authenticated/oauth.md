# Bearer Token

Get bearer token issued by the authorization server.

### Sample Request \(Python\)

```py
import time
import requests
import json
from M2Crypto import *
import base64
import random

def privateKeyEncrypt(content, privateKey):
	# Encrypt content and remember to use your privateKey to encrypt
	encryptContent = privateKey.private_encrypt(content, RSA.pkcs1_padding)
	encodeContent = base64.b64encode(encryptContent)
	return encodeContent

# Get accessToken
def getAccessToken(apiKey, clientSecret):
	url = "https://api.bitmart.com/v2/token"
    data = {"grant_type": "client_credentials","client_id": apiKey, "client_secret": clientSecret}
	response = requests.post(url, data = payload)
	print(response.content)
	accessToken = response.json()['data']['access_token']
	return accessToken

############################ Main Function ##########################

    # RSA private key string which starts with "MI"
    privateKey = xxxxxx

    apiKey = xxxx
    apiSecret = xxxx

    timestamp = str(long(time.time() * 1000))

    signContent = apiKey + ":" + apiSecret + ":" + timestamp

    clientSecret = privateKeyEncrypt(signContent, privateKey)
    accessToken = getAccessToken(apiKey, clientSecret)
    print(accessToken)

```


### Sample Response
```js
{
   "access_token":"m261aeb5bfa471c67c6ac41243959ae0dd408838cdc1a47e945305dd558e2fa78",
   "expires_in":900
}
```

#### Response Details

| Key | Type | Description |
| :--- | :--- | :--- |
| access_token | string | Bearer token |
| expires_in | numeric | Token expiration time (in seconds) |





