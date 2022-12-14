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

// CreateVrf - Create a new VRF in the specified project
func CreateVrf(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteVrf - Delete the VRF
func DeleteVrf(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindVrfById - Retrieve a VRF
func FindVrfById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindVrfIpReservations - Retrieve all VRF IP Reservations in the VRF
func FindVrfIpReservations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindVrfs - Retrieve all VRFs in the project
func FindVrfs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateVrf - Update the VRF
func UpdateVrf(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
