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
	"context"
	"github.com/gin-gonic/gin"
	metal "github.com/haozheng95/eqx-gin-server/metal/v1"
	"log"
	"net/http"
)

// CreateBgpSession - Create a BGP session
func CreateBgpSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateDevice - Create a device
func CreateDevice(c *gin.Context) {
	id := c.Param("id")
	var body *metal.CreateDeviceRequestOneOf
	if err := c.BindJSON(&body); err != nil {
		if !ErrorCheck(err, nil, c) {
			return
		}
	}
	//var privateIpv4SubnetSize float32 = 28
	//this.PrivateIpv4SubnetSize = &privateIpv4SubnetSize
	//var publicIpv4SubnetSize float32 = 31
	//this.PublicIpv4SubnetSize = &publicIpv4SubnetSize
	respDevice, r, err := apiClient.DevicesApi.FindProjectDevices(context.Background(), id).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	if len(respDevice.Devices) >= 3 {
		hostIPs := make([]string, 3)
		for i, device := range respDevice.Devices {
			log.Println("i=", i)
			for j, ip := range device.IpAddresses {
				log.Println("j=", j)
				log.Println("ip.Address=", ip.GetAddress())
				log.Println("ip.GetPublic=", ip.GetPublic())
				if ip.GetPublic() && IsIPv4(ip.GetAddress()) {
					addAuthorizedKeys(ip.GetAddress(), body)
					hostIPs = append(hostIPs, ip.GetAddress())
					continue
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{"msg": "The number of created instances reached the upper limit. Procedure"})
		return
	}

	resp, r, err := apiClient.DevicesApi.CreateDevice(context.Background(), id).CreateDeviceRequest(metal.CreateDeviceRequestOneOfAsCreateDeviceRequest(body)).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CreateIPAssignment - Create an ip assignment
func CreateIPAssignment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteDevice - Delete the device
func DeleteDevice(c *gin.Context) {
	id := c.Param("id")
	r, err := apiClient.DevicesApi.DeleteDevice(context.Background(), id).ForceDelete(true).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	c.JSON(http.StatusOK, gin.H{"resp": "delete success"})
}

// FindBgpSessions - Retrieve all BGP sessions
func FindBgpSessions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindDeviceById - Retrieve a device
func FindDeviceById(c *gin.Context) {
	id := c.Param("id")
	resp, r, err := apiClient.DevicesApi.FindDeviceById(context.Background(), id).Include(include).Exclude(exclude).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	c.JSON(http.StatusOK, resp)

}

// FindDeviceCustomdata - Retrieve the custom metadata of an instance
func FindDeviceCustomdata(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindDeviceMetadataByID - Retrieve metadata
func FindDeviceMetadataByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindDeviceUserdataByID - Retrieve userdata
func FindDeviceUserdataByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindIPAssignmentCustomdata - Retrieve the custom metadata of an IP Assignment
func FindIPAssignmentCustomdata(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindIPAssignments - Retrieve all ip assignments
func FindIPAssignments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindInstanceBandwidth - Retrieve an instance bandwidth
func FindInstanceBandwidth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindOrganizationDevices - Retrieve all devices of an organization
func FindOrganizationDevices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindProjectDevices - Retrieve all devices of a project
func FindProjectDevices(c *gin.Context) {
	id := c.Param("id")

	resp, r, err := apiClient.DevicesApi.FindProjectDevices(context.Background(), id).Execute()
	//page := int32(1)     // int32 | Page to return (optional) (default to 1)
	//perPage := int32(10) // int32 | Items returned per page (optional) (default to 10)
	//resp, r, err := apiClient.DevicesApi.FindProjectDevices(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FindTraffic - Retrieve device traffic
func FindTraffic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetBgpNeighborData - Retrieve BGP neighbor data for this device
func GetBgpNeighborData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PerformAction - Perform an action
func PerformAction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateDevice - Update the device
func UpdateDevice(c *gin.Context) {
	id := c.Param("id")                 // string | Device UUID
	var body *metal.UpdateDeviceRequest // UpdateDeviceRequest | Facility to update
	if err := c.BindJSON(&body); err != nil {
		// DO SOMETHING WITH THE ERROR
		if !ErrorCheck(err, nil, c) {
			return
		}
	}
	resp, r, err := apiClient.DevicesApi.UpdateDevice(context.Background(), id).UpdateDeviceRequest(*body).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	c.JSON(http.StatusOK, resp)
}
