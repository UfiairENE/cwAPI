# cwAPI
Cryptowatch API package- written in Go


go get -u -v github.com/ireneufia/cwAPI

Get single coin info

getBTC, err := cwAPI.GetCoinInfo("bitcoin")
Get Allcoins info (all/limit)

getCoins, err := cwAPI.GetAllCoinInfo(0)
If you want to limit results to top 10, use cwAPI.GetAllCoinInfo(10)

Get Global Data

getMarket, err := cwAPI.GetMarketData()

See the main.go for more details.
