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

// CreateSpotMarketRequest - Create a spot market request
func CreateSpotMarketRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteSpotMarketRequest - Delete the spot market request
func DeleteSpotMarketRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindMetroSpotMarketPrices - Get current spot market prices for metros
func FindMetroSpotMarketPrices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindSpotMarketPrices - Get current spot market prices
func FindSpotMarketPrices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindSpotMarketPricesHistory - Get spot market prices for a given period of time
func FindSpotMarketPricesHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindSpotMarketRequestById - Retrieve a spot market request
func FindSpotMarketRequestById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ListSpotMarketRequests - List spot market requests
func ListSpotMarketRequests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
