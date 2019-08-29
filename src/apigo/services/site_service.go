package services

import (
	"../domains"
	"../utils"
)

func GetSite(siteId string) (*domains.Site, *utils.Apierror) {
	site := domains.Site{
		ID: siteId,
	}
	if err := site.Get(); err != nil {
		return nil, err

	}
	return &site, nil

}
