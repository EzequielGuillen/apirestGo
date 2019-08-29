package services

import (
	"../domains"
	"../utils"
)

func GetResultChan(userId int) (*domains.Result, *utils.Apierror) {

	user := domains.User{
		ID: userId,
	}
	apierr := user.Get()
	if apierr != nil {
		return nil, apierr
	}
	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}

	c := make(chan domains.Result)
	erc := make(chan *utils.Apierror)
	go domains.GetChanSite(site, c, erc)
	go domains.GetChanCountry(country, c, erc)
	result := domains.Result{User: &user}


	for i := 0; i < 2; i++ {
		select {
		case err := <-erc:
			return nil, err
		case res := <-c:
			if res.Site != nil {
				result.Site = res.Site
			}
			if res.Country != nil {
				result.Country = res.Country
			}

		}

	}

	return &result, nil

}
