# Ticker

The ticker is a high level overview of the state of the market. It shows you the current best bid and ask, as well as the last trade price. It also includes information such as daily volume and how much the price has moved over the last day.

### Sample Request \(Python\)

```py
import requests

url = "https://api.bitmart.com/v2/ticker?symbol=BMX_ETH"

response = requests.request("GET", url)

print(response.text)
```

### Sample Response

```js
{  
   "volume":"1.000000",
   "base_volume":"0.000500",
   "highest":"0.000500",
   "lowest":"0.000201",
   "current_price":"0.000500",
   "ask_1":"0.000500",
   "ask_1_amount": "6997.00",
   "bid_1":"0.000201",
   "bid_1_amount": "27.82"
   "fluctuation":"148.76%",
   "url":"https://www.bitmart.com/trade?symbol=BMX_ETH"
}
```

#### Request Parameter

| Parameter | Type | Description |
| :--- | :--- | :--- |
| symbol | query | Trading pair symbol (optional, if not passed in, return all products' ticker information) |

#### Response Details

| Key | Type | Description |
| :--- | :--- | :--- |
| volume | string | Trading volume |
| base\_volume | string | Target volume |
| highest | string | Highest price within 24 hr |
| lowest | string | Lowest price within 24 hr |
| current\_price | string | Current price |
| ask\_1 | string | Asked price |
| ask\_1\_amount | string | Asked amount
| bid\_1 | string | Bid price |
| bid\_1\_amount | string | Bid amount
| fluctuation | string | Price change within 24 hr |
| url | string | Trading url at bitmart.com |



