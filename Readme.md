# Problem statement -> Stock Server
​
Based on https://www.worldtradingdata.com/

Create an end point /stock/{symbol} or a URL with query params localhost:3000/?stock_exchange=XXX

This should return the stock prices of the given stock symbol from the exchanges provided in the query parameter.

Symbol must be a valid stock symbol Example: /stock/NSE,NASDAQ

## Run with docker

`docker compose up`

Then access the app on port 8080. Ex:

`http://localhost:8080/stock/AAPL`