package storage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Fetching/structs"
)

//function show all storages + currency in USD
func ShowStorage(c *gin.Context) {

	type Storages []structs.Storage

	type IDRUSD struct {
		IDR_USD float32
	}

	var (
		result   gin.H
		storages Storages
		idr_usd  IDRUSD
	)

	forex, _ := http.Get("https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=d6c36c654fa9cf38b219")

	df, _ := ioutil.ReadAll(forex.Body)

	json.Unmarshal(df, &idr_usd)

	response, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		json.Unmarshal(data, &storages)

		// log.Println(body)
		for i, storage := range storages {
			valusd, _ := strconv.Atoi(storage.Price)
			storages[i].Usd = ConvertIDRtoUSD(valusd) * idr_usd.IDR_USD
		}

		result = gin.H{
			"data": storages,
		}
	}

	c.JSON(http.StatusOK, result)
}

// convert int to float32
func ConvertIDRtoUSD(theIDR int) float32 {

	var y float32 = float32(theIDR)

	return y

}
