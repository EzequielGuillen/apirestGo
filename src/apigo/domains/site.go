package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings           struct {
		IdentificationTypes      []string `json:"identification_types"`
		TaxpayerTypes            []string `json:"taxpayer_types"`
		IdentificationTypesRules []struct {
			IdentificationType string `json:"identification_type"`
			Rules              []struct {
				EnabledTaxpayerTypes []string `json:"enabled_taxpayer_types"`
				BeginsWith           string   `json:"begins_with"`
				Type                 string   `json:"type"`
				MinLength            int      `json:"min_length"`
				MaxLength            int      `json:"max_length"`
			} `json:"rules"`
		} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
}

func (site *Site) Get() *utils.Apierror {


	if site.ID == "" {
		return &utils.Apierror{
			Message: "Site ID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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


	if err := json.Unmarshal(data, &site); err != nil {

		return &utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

	}

	return nil

}

func (site *Site) GetWG(wg *sync.WaitGroup, apierror *utils.Apierror) {

	defer wg.Done()

	if site.ID == "" {
		apierror.Message = "Site ID is empty"
		apierror.Status = http.StatusBadRequest

		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {

		apierror.Message = err.Error()
		apierror.Status = http.StatusInternalServerError
		return

	}


}

func GetChanSite(site Site, result chan Result, erc chan *utils.Apierror)  {
	
	if site.ID == "" {
		
		apierror :=  &utils.Apierror{
			Message: "Site ID is empty",
			Status:  http.StatusBadRequest,
		}
		
		erc <- apierror
		
		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {

		apierror :=  &utils.Apierror{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

		erc <- apierror

		return

	}

	result <- Result{
		Site: &site,
	}

	return

}
