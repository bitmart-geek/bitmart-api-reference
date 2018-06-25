# Bearer Token

Get bearer token issued by the authorization server.

### Sample Request \(Python\)

```py
from M2Crypto import *
import time
import base64
import requests
import json

api_url = "http://api.bitmart.com/v2/access"
api_key = xxx # please replace with your api key here
api_secret = xxx # please replace with your api secret here
private_key = xxx # please replace with your private key here

def encrypt_rsa_pri_key(message):
    '''
    Encrypt with private key.
    '''

    pem_prefix = '-----BEGIN RSA PRIVATE KEY-----\n'
    pem_suffix = '\n-----END RSA PRIVATE KEY-----'
    private_key_str = '{}{}{}'.format(pem_prefix, private_key, pem_suffix)

    pri_key = RSA.load_key_string(private_key_str)
    # pri_key = RSA.load_key('_pri.pem') # You can also load key from .pem file.
    encrypted = pri_key.private_encrypt(message, RSA.pkcs1_padding)
    ciphertext = base64.b64encode(encrypted)

    return ciphertext

def get_signed_content(api_key, api_secret):
    timestamp = time.time() * 1000
    signed_content = api_key + ":" + api_secret + ":" + str(long(timestamp))
    return signed_content


def get_access_token(apiKey, client_secret):
    url = "https://api.bitmart.com/v2/token"
    data = {"grant_type": "client_credentials","client_id": apiKey, "client_secret": client_secret}
    response = requests.post(url, data = data)
    print(response.content)
    accessToken = response.json()['access_token']
    return accessToken

if __name__ == '__main__':
    signed_content = get_signed_content(api_key, api_secret)
    client_secret = encrypt_rsa_pri_key(signed_content)

    print "client_secret:"
    print client_secret

    access_token = get_access_token(api_key, client_secret)
    print access_token

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





