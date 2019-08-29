package services

import (
	"../domains"
	"../utils"
)

func GetResult(userId int) (*domains.Result, *utils.Apierror) {

	user := domains.User{
		ID: userId,
	}
	err := user.Get()
	if err != nil {
		return nil, err
	}

	country := domains.Country{
		ID: user.CountryID,
	}
	err = country.Get()
	if err != nil {
		return nil, err
	}

	site := domains.Site{
		ID: user.SiteID,
	}
	err = site.Get()
	if err != nil {
		return nil, err
	}

	result := domains.Result{
		User:    &user,
		Site:    &site,
		Country: &country,
	}

	return &result, nil

}
