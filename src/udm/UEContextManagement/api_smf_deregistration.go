/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UEContextManagement

import (
	"github.com/gin-gonic/gin"
	"free5gc/lib/http_wrapper"
	"free5gc/src/udm/udm_handler"
	"free5gc/src/udm/udm_handler/udm_message"
)

// DeregistrationSmfRegistrations - delete an SMF registration
func DeregistrationSmfRegistrations(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["pduSessionId"] = c.Params.ByName("pduSessionId")

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventDeregistrationSmfRegistrations, req)
	udm_handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	return
}
