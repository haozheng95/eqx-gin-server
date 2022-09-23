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

// DeletePaymentMethod - Delete the payment method
func DeletePaymentMethod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindPaymentMethodById - Retrieve a payment method
func FindPaymentMethodById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePaymentMethod - Update the payment method
func UpdatePaymentMethod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}