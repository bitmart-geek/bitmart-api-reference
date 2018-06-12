# Symbol Details

Get the symbol detail.

### Sample Request \(Python\)

```py
import requests

url = "https://api.bitmart.com/v2/symbols/BMX_ETH"

response = requests.request("GET", url)

print(response.text)
```

### Sample Response

```js
{  
   "min_precision":4,
   "max_precision":6,
   "from":"BMX",
   "to":"ETH",
   "expiration":"NA"
}
```

#### Request Parameter

| Parameter | Type | Description |
| :--- | :--- | :--- |
| symbol | path | Trading pair symbol |

#### Response Details

| Key | Type | Description |
| :--- | :--- | :--- |
| min\_precision | number | Minimum precision of this pair |
| max\_precision | number | Maximum precision of this pair |
| from | string | Trade from |
| to | string | Trade to |
| expiration | string | Expiration date for limited contracts/pairs |



