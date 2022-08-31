package api

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/thyra/api/swagger/server/models"
	"github.com/massalabs/thyra/api/swagger/server/restapi/operations"
	"github.com/massalabs/thyra/pkg/gui"
	"github.com/massalabs/thyra/pkg/onchain/website"
	"github.com/massalabs/thyra/pkg/wallet"
)

func PrepareForWebsiteHandler(params operations.WebsiteCreatorPrepareParams) middleware.Responder {

	wallet, err := wallet.GetWallet(*params.Body.Nickname)
	if err != nil {
		return operations.NewWebsiteCreatorPrepareInternalServerError().
			WithPayload(
				&models.Error{
					Code:    errorCodeWebCreatorPrepare,
					Message: err.Error(),
				})
	}
	password := gui.Password(*params.Body.Nickname)

	err = wallet.Unprotect(password, 0)
	if err != nil {
		return operations.NewWebsiteCreatorPrepareInternalServerError().
			WithPayload(
				&models.Error{
					Code:    errorCodeWalletWrongPassword,
					Message: err.Error(),
				})
	}

	address, err := website.PrepareForUpload(*params.Body.URL, wallet)
	if err != nil {
		return operations.NewWebsiteCreatorPrepareInternalServerError().
			WithPayload(
				&models.Error{
					Code:    errorCodeWebCreatorPrepare,
					Message: err.Error(),
				})
	}

	return operations.NewWebsiteCreatorPrepareOK().
		WithPayload(
			&models.Websites{
				Name:    *params.Body.URL,
				Address: address,
			})
}

func UploadWebsiteHandler(params operations.WebsiteCreatorUploadParams) middleware.Responder {

	wallet, err := wallet.GetWallet(params.Nickname)
	if err != nil {
		return operations.NewWebsiteCreatorPrepareInternalServerError().
			WithPayload(
				&models.Error{
					Code:    errorCodeWebCreatorPrepare,
					Message: err.Error(),
				})
	}
	password := gui.Password(params.Nickname)
	err = wallet.Unprotect(password, 0)
	if err != nil {
		return operations.NewWebsiteCreatorPrepareInternalServerError().
			WithPayload(
				&models.Error{
					Code:    errorCodeWalletWrongPassword,
					Message: err.Error(),
				})
	}
	archive, err := ioutil.ReadAll(params.Zipfile)
	if err != nil {
		return operations.NewFillWebPostInternalServerError().
			WithPayload(&models.Error{
				Code:    errorCodeWebCreatorReadArchive,
				Message: err.Error(),
			})
	}

	b64 := base64.StdEncoding.EncodeToString(archive)

	_, err = website.Upload(params.Address, b64, wallet)
	if err != nil {
		return operations.NewFillWebPostInternalServerError().
			WithPayload(&models.Error{
				Code:    errorCodeWebCreatorUpload,
				Message: err.Error(),
			})
	}

	return operations.NewFillWebPostOK().
		WithPayload(&models.Websites{
			Name:    "Name",
			Address: params.Address})
}