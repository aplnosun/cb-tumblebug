package mcir

import (
	"fmt"
	"net/http"

	"github.com/cloud-barista/cb-tumblebug/src/core/common"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
	"github.com/labstack/echo/v4"
)

// RestPostSpec godoc
// @Summary Register spec
// @Description Register spec
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param registeringMethod query string true "registerWithInfo or else"
// @Param nsId path string true "Namespace ID"
// @Param specInfo body mcir.TbSpecInfo false "Details for an spec object"
// @Param specName body mcir.TbSpecReq false "name, connectionName and cspSpecName"
// @Success 200 {object} mcir.TbSpecInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/spec [post]
func RestPostSpec(c echo.Context) error {

	nsId := c.Param("nsId")

	action := c.QueryParam("action")
	fmt.Println("[POST Spec requested action: " + action)

	if action == "registerWithInfo" { // `RegisterSpecWithInfo` will be deprecated in Cappuccino.
		fmt.Println("[Registering Spec with info]")
		u := &mcir.TbSpecInfo{}
		if err := c.Bind(u); err != nil {
			return err
		}
		content, err := mcir.RegisterSpecWithInfo(nsId, u)
		if err != nil {
			common.CBLog.Error(err)
			mapA := map[string]string{
				"message": err.Error()}
			return c.JSON(http.StatusFailedDependency, &mapA)
		}
		return c.JSON(http.StatusCreated, content)

	} else { // if action == "registerWithCspSpecName" { // The default mode.
		fmt.Println("[Registering Spec with CspSpecName]")
		u := &mcir.TbSpecReq{}
		if err := c.Bind(u); err != nil {
			return err
		}
		content, err := mcir.RegisterSpecWithCspSpecName(nsId, u)
		if err != nil {
			common.CBLog.Error(err)
			mapA := map[string]string{
				"message": err.Error()}
			return c.JSON(http.StatusFailedDependency, &mapA)
		}
		return c.JSON(http.StatusCreated, content)

	} /* else {
		mapA := map[string]string{"message": "LookupSpec(specRequest) failed."}
		return c.JSON(http.StatusFailedDependency, &mapA)
	} */

}

/* function RestPutSpec not yet implemented
// RestPutSpec godoc
// @Summary Update spec
// @Description Update spec
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param specInfo body mcir.TbSpecInfo true "Details for an spec object"
// @Success 200 {object} mcir.TbSpecInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/spec/{specId} [put]
*/
func RestPutSpec(c echo.Context) error {
	//nsId := c.Param("nsId")

	return nil
}

// Request structure for RestLookupSpec
type RestLookupSpecRequest struct {
	ConnectionName string `json:"connectionName"`
}

// RestLookupSpec godoc
// @Summary Lookup spec
// @Description Lookup spec
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param connectionName body RestLookupSpecRequest true "Specify connectionName"
// @Param specName path string true "Spec name"
// @Success 200 {object} mcir.SpiderSpecInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /lookupSpec/{specName} [get]
func RestLookupSpec(c echo.Context) error {

	u := &RestLookupSpecRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}

	specName := c.Param("specName")
	fmt.Println("[Lookup spec]" + specName)
	content, err := mcir.LookupSpec(u.ConnectionName, specName)
	if err != nil {
		common.CBLog.Error(err)
		return c.JSONBlob(http.StatusFailedDependency, []byte(err.Error()))
	}

	return c.JSON(http.StatusOK, &content)

}

// RestLookupSpecList godoc
// @Summary Lookup spec list
// @Description Lookup spec list
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param connectionName body RestLookupSpecRequest true "Specify connectionName"
// @Success 200 {object} mcir.SpiderSpecList
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /lookupSpec [get]
func RestLookupSpecList(c echo.Context) error {

	//type JsonTemplate struct {
	//	ConnectionName string
	//}

	u := &RestLookupSpecRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println("[Get Region List]")
	content, err := mcir.LookupSpecList(u.ConnectionName)
	if err != nil {
		common.CBLog.Error(err)
		return c.JSONBlob(http.StatusFailedDependency, []byte(err.Error()))
	}

	return c.JSON(http.StatusOK, &content)

}

// RestFetchSpecs godoc
// @Summary Fetch specs
// @Description Fetch specs
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/fetchSpecs [post]
func RestFetchSpecs(c echo.Context) error {

	nsId := c.Param("nsId")

	connConfigCount, specCount, err := mcir.FetchSpecs(nsId)
	if err != nil {
		common.CBLog.Error(err)
		mapA := map[string]string{
			"message": err.Error()}
		return c.JSON(http.StatusFailedDependency, &mapA)
	}

	mapA := map[string]string{
		"message": "Fetched " + fmt.Sprint(specCount) + " specs (from " + fmt.Sprint(connConfigCount) + " connConfigs)"}
	return c.JSON(http.StatusCreated, &mapA) //content)
}

// RestGetSpec godoc
// @Summary Get spec
// @Description Get spec
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param specId path string true "Spec ID"
// @Success 200 {object} mcir.TbSpecInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/spec/{specId} [get]
func RestGetSpec(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}

// Response structure for RestGetAllSpec
type RestGetAllSpecResponse struct {
	Spec []mcir.TbSpecInfo `json:"spec"`
}

// RestGetAllSpec godoc
// @Summary List all specs
// @Description List all specs
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} RestGetAllSpecResponse
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/spec [get]
func RestGetAllSpec(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}

// RestDelSpec godoc
// @Summary Delete spec
// @Description Delete spec
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param specId path string true "Spec ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/spec/{specId} [delete]
func RestDelSpec(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}

// RestDelAllSpec godoc
// @Summary Delete all specs
// @Description Delete all specs
// @Tags Spec
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/spec [delete]
func RestDelAllSpec(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}
