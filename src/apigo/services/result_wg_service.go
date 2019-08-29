package services

import (
	"../domains"
	"../utils"
	"sync"
)

func GetResultWG(userId int) (*domains.Result, *utils.Apierror) {

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

	err := utils.Apierror{}
	var wg sync.WaitGroup
	wg.Add(2)
	go country.GetWG(&wg, &err)
	go site.GetWG(&wg, &err)
	wg.Wait()

	if err.Status != 0 {
		return nil, &err
	}

	result := domains.Result{
		User:    &user,
		Site:    &site,
		Country: &country,
	}

	return &result, nil

}
