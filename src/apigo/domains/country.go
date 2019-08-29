package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) Get() *utils.Apierror {

	if country.ID == "" {
		return &utils.Apierror{
			Message: "Country ID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)



	response, err :=utils.Client.Get(url)

	if err != nil {

		return &utils.Apierror{
			Message: err.Error(),
			Status:  533,
		}

	}

	if response.StatusCode==500 {

		return &utils.Apierror{
			Message: "Server Error",
			Status:  533,
		}

	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return &utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

	}

	if err := json.Unmarshal(data, &country); err != nil {

		return &utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

	}

	return nil

}

func (country *Country) GetWG(wg *sync.WaitGroup, apierror *utils.Apierror) {
	defer wg.Done()

	if country.ID == "" {
		apierror.Message = "Country ID is empty"
		apierror.Status = http.StatusBadRequest

		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)

	response, err :=utils.Client.Get(url)


	if err != nil {

		apierror.Message = err.Error()
		apierror.Status = 533
		return

	}
	if response.StatusCode==500 {

		apierror.Message = "Server Error"
		apierror.Status = 533
		return

	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {

		apierror.Message = err.Error()
		apierror.Status = http.StatusInternalServerError
		return

	}

	if err := json.Unmarshal(data, &country); err != nil {

		apierror.Message = err.Error()
		apierror.Status = http.StatusInternalServerError
		return

	}

	return

}

func GetChanCountry(country Country, result chan Result, erc chan *utils.Apierror)  {

	if country.ID == "" {

		apierror :=  &utils.Apierror{
			Message: "Site ID is empty",
			Status:  http.StatusBadRequest,
		}

		erc <- apierror

		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)

	response, err :=utils.Client.Get(url)

	if err != nil {

		apierror :=  &utils.Apierror{
			Message: err.Error(),
			Status:  533,
		}

		erc <- apierror
		return

	}
	if response.StatusCode==500 {

		apierror :=  &utils.Apierror{
			Message: "Server Error",
			Status:  533,
		}

		erc <- apierror
		return

	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {

		apierror :=  &utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

		erc <- apierror

		return

	}

	if err := json.Unmarshal(data, &country); err != nil {

		apierror :=  &utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

		erc <- apierror

		return

	}

	result <- Result{
		Country: &country,
	}

	return

}