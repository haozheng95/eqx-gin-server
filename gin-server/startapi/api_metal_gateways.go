/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package startapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMetalGateway - Create a metal gateway
func CreateMetalGateway(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteMetalGateway - Deletes the metal gateway
func DeleteMetalGateway(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindMetalGatewayById - Returns the metal gateway
func FindMetalGatewayById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindMetalGatewaysByProject - Returns all metal gateways for a project
func FindMetalGatewaysByProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
