# Order Book

Get the full order book.

### Sample Request \(Python\)

```py
import requests

url = "https://api.bitmart.com/v2/symbols/BMX_ETH/orders?precision=6"

response = requests.request("GET", url)

print(response.text)
```

### Sample Response

```js
{  
   "buys":[
      {
         "amount":"99999596.79",
         "total":"99999999.000000",
         "price":"0.000201",
         "count":"1"
      },
      {
         "amount":"68.00",
         "total":"68.000000",
         "price":"0.000102",
         "count":"1"
      },
      {
         "amount":"165.00",
         "total":"55.000000",
         "price":"0.000053",
         "count":"3"
      },
      {
         "amount":"55.00",
         "total":"55.000000",
         "price":"0.000028",
         "count":"1"
      },
      {
         "amount":"50.00",
         "total":"50.000000",
         "price":"0.000020",
         "count":"1"
      },
      {
         "amount":"0.05",
         "total":"0.050000",
         "price":"0.000019",
         "count":"1"
      },
      {
         "amount":"5019.97",
         "total":"20.000000",
         "price":"0.000010",
         "count":"2"
      },
      {
         "amount":"166.00",
         "total":"10.000000",
         "price":"0.000005",
         "count":"17"
      },
      {
         "amount":"1.00",
         "total":"1.000000",
         "price":"0.000001",
         "count":"1"
      }
   ],
   "sells":[
      {
         "amount":"25.61",
         "total":"26.610000",
         "price":"0.000500",
         "count":"1"
      },
      {
         "amount":"7000.01",
         "total":"0.010000",
         "price":"1.000000",
         "count":"2"
      }
   ]
}
```

#### Request Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| symbol | path | Trading pair symbol |
| precision | query | Price precision whose range is defined in symbol details |

#### Response Details

| Key | Type | Description |
| :--- | :--- | :--- |
| amount | string | This is the number of coins offered at a specific price point. |
| total | string | This is the cumulation number of coins that offered to a specific price point. |
| price | string | This is the offer price by a trader or traders at a specific price point. |
| count | string | This is the total number of orders at that level. |
| buys | array | Array of 'buy' type orders |
| sells | array | Array of 'sell' type orders |



