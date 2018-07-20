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
   "base_currency":"BMX",
   "quote_currency":"ETH",
   "quote_increment":"0.000001"
   "base_min_size":"0.01"
   "base_max_size":"100000000"
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
| base\_currency | string | Base currency |
| quote\_currency | string | Quote currency |
| quote\_increment | number | Minimum order price as well as the price increment |
| base\_min\_size | number | Minimum trade volume |
| base\_max\_size | number | Maximum trade volume |
| expiration | string | Expiration date for limited contracts/pairs |



