package service

import (
	"context"
	"log"
	"strconv"
	"strings"

	"druva.com/godevkit/logger/svclog"
	"git.druva.org/druva.com/DI-Master/config"
	"git.druva.org/druva.com/DI-Master/rerror"
	"git.druva.org/druva.com/DI-Master/restlib/authjwt"
	"git.druva.org/druva.com/DI-Master/rtypes"
	licensingClient "git.druva.org/druva.com/DI-Master/sdk/remote/license"
	"github.com/gin-gonic/gin"
)

var Pkg_name = "service"
var logger = svclog.NewStdoutCtxLogger("Realize")
var RemoteLicensingClient licensingClient.LicenseServiceJob

//handler for GetTotalBackupDataForDays
func GetTotalBackupDataForDayshandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var reqbody rtypes.TotalbackedupdataReq
	var err error
	//Reading filter from Request Body
	err = c.BindJSON(&reqbody)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	if reqbody.StartDate < 0 || reqbody.EndDate < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.InValidDate))
		return
	}
	productsCustomers, err = GetCustomerAndProductIds(c)

	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := GetTotalBackupDataForDays(productsCustomers, reqbody.StartDate, reqbody.EndDate, reqbody.Filter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.IndentedJSON(STATUS200, response)
	return
}

//handler for Get File Types api
func GetFileTypesHandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.GetFileTypesReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}
	//TODO
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}

	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := GetFileTypes(productsCustomers, reqbody.Date, reqbody.Limit, reqbody.Filter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetFileTypesHandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.IndentedJSON(STATUS200, response)
	return

}

//handler for Get File Extensions api
func GetFileExtensionsHandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.GetFileExtensionsReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}
	//TODO
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := GetFileExtension(productsCustomers, reqbody.Date, reqbody.Limit, reqbody.Filter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetFileExtensionsHandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.IndentedJSON(STATUS200, response)
	return

}

//handler for GetDataByResourceType api
func GetDataByResourceTypehandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.GetDataByResourceTypeReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}
	//TODO
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := GetDataByResourceType(productsCustomers, reqbody.Date, reqbody.Filter)
	if err != nil && false == strings.Contains(err.Error(), rerror.ErrorMessage(rerror.NoContent)) {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataByResourceTypehandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	// if err != nil && strings.Contains(err.Error(), rerror.ErrorMessage(rerror.NoContent)) {
	// 	c.IndentedJSON(STATUS200, gin.H{})
	// } else {
	// 	c.IndentedJSON(STATUS200, response)
	// }
	c.IndentedJSON(STATUS200, response)
	return
}

//handler for GetDataSourceGroups api
func GetDataSourceGroupshandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productId []int
	var err error
	//TODO
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productId = append(productId, config.ProductIdPhoenix)

	var epochDate int64
	if queryParam, ok := c.GetQuery("date"); ok {
		epochDate, _ = strconv.ParseInt(queryParam, 10, 64)
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "queryParam:", queryParam, "date:", epochDate)
	}

	response, err := GetDataSourceGroups(productsCustomers, epochDate)
	if err != nil && false == strings.Contains(err.Error(), rerror.ErrorMessage(rerror.NoContent)) {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	if err != nil && strings.Contains(err.Error(), rerror.ErrorMessage(rerror.NoContent)) {
		c.IndentedJSON(STATUS200, gin.H{})
	} else {
		c.IndentedJSON(STATUS200, response)
	}
	return

}

//hadler for Get FileTypes And Extension API
func GetFileTypesAndExtensionHandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productId []int
	var err error
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productId = append(productId, config.ProductIdPhoenix)
	response, err := GetFileTypesAndExtension(productsCustomers)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.IndentedJSON(STATUS200, response)
	return

}

//handler for get Source List api
func GetSourceListhandler(c *gin.Context) {

	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.GetSourceListReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	jwtwebtoken := c.Request.Header.Get("Authontication")
	response, err := GetSourceList(productsCustomers, reqbody.Date, reqbody.Limit, reqbody.Filter, jwtwebtoken)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.IndentedJSON(STATUS200, response)
	return
}

//handler for Get Distribution By Size api
func GetDistributionBySizeHandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.GetDistributionBySizeReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}

	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDistributionBySizeHandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}

	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := GetDistributionBySize(productsCustomers, reqbody.Date, reqbody.SizeRange, reqbody.Filter)

	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.IndentedJSON(STATUS200, response)

	return

}

//handler for distribution by modified month
func DistributionByModifiedMonthhandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.DistributionByModifiedMonthReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}
	//Read customer id and product id from jwt
	//
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := DistributionByModifiedMonth(productsCustomers, reqbody.Date, reqbody.Filter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "DistributionByModifiedMonthhandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	c.SecureJSON(STATUS200, response)
	return
}

// handler for Get Organizations api
func GetOrganizationshandler(c *gin.Context) {
	var productsCustomers map[int]int
	var productid []int
	var err error
	var reqbody rtypes.GetOrganizationsReq
	err = c.BindJSON(&reqbody)
	if err != nil {
		c.IndentedJSON(rerror.BadRequest, rerror.Error(rerror.BadRequest))
		return
	}
	/*date validation*/
	if reqbody.Date < 0 {
		c.IndentedJSON(rerror.BadRequest, rerror.ErrorMessage(rerror.InValidDate))
		return
	}
	productsCustomers, err = GetCustomerAndProductIds(c)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroupshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	//productsCustomers[ProductIdPhoenix] = 2
	productid = append(productid, config.ProductIdPhoenix)
	response, err := GetOrganizations(productsCustomers, reqbody.Date, reqbody.Limit, reqbody.Filter, c.Request.Header.Get("Authontication"))
	if err != nil && false == strings.Contains(err.Error(), rerror.ErrorMessage(rerror.NoContent)) {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetOrganizationshandler", "Error:", err.Error())
		respErr := ErrorResponse(err)
		c.IndentedJSON(respErr.(rerror.RError).Code, respErr)
		return
	}
	// if err != nil && strings.Contains(err.Error(), rerror.ErrorMessage(rerror.NoContent)) {
	// 	c.IndentedJSON(STATUS200, response)
	// } else {
	// 	c.IndentedJSON(STATUS200, response)
	// }
	c.IndentedJSON(STATUS200, response)
	return
}

//funtion to convert all error to rerror
//input pram : error / rerror
//output pram : rerror
//if provided error is of type error function will return the rerrorr with error code of 500(internal server Error)
//if provided error is rerror and it is Api Exposed error code then return same error
//if provided error is rerror and it is not Api Exposed then return rerror with error code of 500(internal server Error)
func ErrorResponse(err error) error {
	var errRes rerror.RError
	data := make(map[string]string)
	switch err.(type) {
	case rerror.RError:
		if err.(rerror.RError).Code >= rerror.ApiExposedError {
			errRes = err.(rerror.RError)
		} else {

			errRes = rerror.Error(rerror.InternalError)
		}
	case error:
		data["error"] = err.Error()
		errRes = rerror.Error(rerror.InternalError, nil, data)

	default:
		errRes = rerror.Error(rerror.InternalError)
	}

	return errRes
}

//function to read customerID and Product ID from jWT token
//input param : jwt token from authjwt.ReadJWT function
//output: Customer ID(string) , list of product id(int), error
func GetCustomerAndProductIds(c *gin.Context) (map[int]int, error) {
	//var customerId int
	var err error
	productsCustomers := make(map[int]int)
	jwtwebtoken := c.Request.Header.Get("Authontication")
	jwtdata, err := authjwt.ReadJWT(jwtwebtoken, 0, "")

	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetCustomerAndProductIds", "Error:", err.Error())
		return productsCustomers, err
	} else {
		if jwtdata != nil {
			if data, ok := jwtdata["admin_data"]; ok && len(data.(map[string]interface{})) > 0 {
				role := int(data.(map[string]interface{})["role"].(float64))
				if role != authjwt.DruvOneAdminRole {
					logger.Debug(context.Background(), "package", Pkg_name, "Method", "getCustomerAndProductID", data.(map[string]interface{})["role"])
					return productsCustomers, rerror.Error(rerror.MethodNotAllowed)
				}
				licenseCustomers, err := RemoteLicensingClient.GetAllSAlicenseCustomer(context.Background(), jwtwebtoken)
				if err != nil {
					logger.Debug(context.Background(), "package", Pkg_name, "Method", "getCustomerAndProductID", "Error for GetAllSAlicenseCustomer :", err.Error())
					return productsCustomers, rerror.Error(rerror.InternalError)
				} else {
					if len(licenseCustomers.LicensedCustomers) != 0 {
						if len(licenseCustomers.LicensedCustomers[0].Licenses) == 0 {
							logger.Debug(context.Background(), "package", Pkg_name, "Method", "getCustomerAndProductID", "Error:", "Empty LicensesInfo")
							return productsCustomers, rerror.Error(rerror.MethodNotAllowed)
						}
					} else {
						logger.Debug(context.Background(), "package", Pkg_name, "Method", "getCustomerAndProductID", "Error:", "Empty LicensedCustomersInfo")
						return productsCustomers, rerror.Error(rerror.MethodNotAllowed)
					}
				}
			}
			/*} else if data, ok := jwtdata["client_data"]; ok && len(data.(map[string]interface{})) > 0 {
				customerId = data.(map[string]interface{})["global_customer_id"].(string)
			} else {
				rootError := make(map[string]string)
				rootError["error"] = rerror.EmptyClientData
				err = rerror.Error(rerror.InValidateJWT, nil, rootError)
			}*/

			if productdata, ok := jwtdata["product_data"]; ok {
				for _, prodata := range productdata.([]interface{}) {
					productdata = prodata.(map[string]interface{})["product_ids"]
					product_id := productdata.(map[string]interface{})["known_product_id"]
					customer_id := productdata.(map[string]interface{})["customer_id"]
					// As of now we are not support the insync products
					if product_id == config.ProductIdPhoenix /*|| product_id == config.ProductIdInsync*/ {
						productsCustomers[product_id.(int)] = customer_id.(int)
					}
				}
			} else {
				rootError := make(map[string]string)
				rootError["error"] = rerror.EmptyProductData
				err = rerror.Error(rerror.InValidateJWT, nil, rootError)
			}

		} else {
			logger.Debug(context.Background(), "package", Pkg_name, "Method", "getCustomerAndProductID", "Error:", "Invalid Jwt Token")
			err = rerror.Error(rerror.InValidateJWT)
		}

	}
	if err == nil {
		if _, ok := productsCustomers[config.ProductIdPhoenix]; false == ok {
			logger.Debug(context.Background(), "package", Pkg_name, "Method", "getCustomerAndProductID", "Error:", "JWT does not have phoenix product")
			err = rerror.Error(rerror.InValidateJWT)
		}
	}
	return productsCustomers, err
}
