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

// CreateLicense - Create a License
func CreateLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteLicense - Delete the license
func DeleteLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindLicenseById - Retrieve a license
func FindLicenseById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindProjectLicenses - Retrieve all licenses
func FindProjectLicenses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateLicense - Update the license
func UpdateLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}