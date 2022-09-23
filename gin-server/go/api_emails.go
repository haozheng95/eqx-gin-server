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

// CreateEmail - Create an email
func CreateEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteEmail - Delete the email
func DeleteEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindEmailById - Retrieve an email
func FindEmailById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateEmail - Update the email
func UpdateEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}