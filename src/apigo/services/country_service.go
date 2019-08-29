package services

import (
	"../domains"
	"../utils"
)

func GetCountry(countryId string) (*domains.Country, *utils.Apierror) {
	country := domains.Country{
		ID: countryId,
	}
	if err := country.Get(); err != nil {
		return nil, err

	}
	return &country, nil

}
