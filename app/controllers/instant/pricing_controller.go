package instant

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/models/incoming_request/instant_service/dto"
	"go_base_project/app/models/instant/dto/request/pricing"
	"go_base_project/app/response"
	"go_base_project/app/services/instant"
	"net/http"
)

type PricingController struct {
}

func (c PricingController) Index(ctx *gin.Context) {

	// binding request to pricing dto request
	var pricingDTO pricing.PricingDTORequest
	errDTO := ctx.ShouldBind(&pricingDTO)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(errDTO.Error(), nil))
		return
	}

	// validate the request
	errValidate := pricingDTO.ValidationPricingDTORequest()
	if errValidate != nil {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(errValidate.Error(), nil))
		return
	}

	// validate the origin
	errValidateSecond := pricingDTO.Origin.ValidateOrigin()
	if errValidateSecond != nil {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(errValidateSecond.Error(), nil))
		return
	}

	// validate the destination
	errValidateDestination := pricingDTO.Destination.ValidationDestination()
	if errValidateDestination != nil {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(errValidateDestination.Error(), nil))
		return
	}

	var data dto.PricingDTO
	data.Origin.Lat = pricingDTO.Origin.Lat
	data.Origin.Long = pricingDTO.Origin.Long
	data.Origin.Address = pricingDTO.Origin.Address
	data.Destination.Lat = pricingDTO.Destination.Lat
	data.Destination.Long = pricingDTO.Destination.Long
	data.Destination.Address = pricingDTO.Destination.Address
	data.Service = pricingDTO.Service
	data.Weight = pricingDTO.Weight
	data.ItemPrice = pricingDTO.ItemPrice

	result := instant.PricingService{Data: data}.Do()

	ctx.JSON(http.StatusOK, response.Api().Success("OK", result.Data))
}
