// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongStruct__dummysDeclaration__ models.GongStruct
var __GongStruct_time__dummyDeclaration time.Duration

// An GongStructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongStruct updateGongStruct deleteGongStruct
type GongStructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongStructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongStruct updateGongStruct
type GongStructInput struct {
	// The GongStruct to submit or modify
	// in: body
	GongStruct *orm.GongStructAPI
}

// GetGongStructs
//
// swagger:route GET /gongstructs gongstructs getGongStructs
//
// # Get all gongstructs
//
// Responses:
// default: genericError
//
//	200: gongstructDBResponse
func GetGongStructs(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStruct.GetDB()

	// source slice
	var gongstructDBs []orm.GongStructDB
	query := db.Find(&gongstructDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongstructAPIs := make([]orm.GongStructAPI, 0)

	// for each gongstruct, update fields from the database nullable fields
	for idx := range gongstructDBs {
		gongstructDB := &gongstructDBs[idx]
		_ = gongstructDB
		var gongstructAPI orm.GongStructAPI

		// insertion point for updating fields
		gongstructAPI.ID = gongstructDB.ID
		gongstructDB.CopyBasicFieldsToGongStruct(&gongstructAPI.GongStruct)
		gongstructAPI.GongStructPointersEnconding = gongstructDB.GongStructPointersEnconding
		gongstructAPIs = append(gongstructAPIs, gongstructAPI)
	}

	c.JSON(http.StatusOK, gongstructAPIs)
}

// PostGongStruct
//
// swagger:route POST /gongstructs gongstructs postGongStruct
//
// Creates a gongstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostGongStruct(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStruct.GetDB()

	// Validate input
	var input orm.GongStructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongstruct
	gongstructDB := orm.GongStructDB{}
	gongstructDB.GongStructPointersEnconding = input.GongStructPointersEnconding
	gongstructDB.CopyBasicFieldsFromGongStruct(&input.GongStruct)

	query := db.Create(&gongstructDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoGongStruct.CheckoutPhaseOneInstance(&gongstructDB)
	gongstruct := (*orm.BackRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr)[gongstructDB.ID]

	if gongstruct != nil {
		models.AfterCreateFromFront(&models.Stage, gongstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongstructDB)
}

// GetGongStruct
//
// swagger:route GET /gongstructs/{ID} gongstructs getGongStruct
//
// Gets the details for a gongstruct.
//
// Responses:
// default: genericError
//
//	200: gongstructDBResponse
func GetGongStruct(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStruct.GetDB()

	// Get gongstructDB in DB
	var gongstructDB orm.GongStructDB
	if err := db.First(&gongstructDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongstructAPI orm.GongStructAPI
	gongstructAPI.ID = gongstructDB.ID
	gongstructAPI.GongStructPointersEnconding = gongstructDB.GongStructPointersEnconding
	gongstructDB.CopyBasicFieldsToGongStruct(&gongstructAPI.GongStruct)

	c.JSON(http.StatusOK, gongstructAPI)
}

// UpdateGongStruct
//
// swagger:route PATCH /gongstructs/{ID} gongstructs updateGongStruct
//
// # Update a gongstruct
//
// Responses:
// default: genericError
//
//	200: gongstructDBResponse
func UpdateGongStruct(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStruct.GetDB()

	// Get model if exist
	var gongstructDB orm.GongStructDB

	// fetch the gongstruct
	query := db.First(&gongstructDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.GongStructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	gongstructDB.CopyBasicFieldsFromGongStruct(&input.GongStruct)
	gongstructDB.GongStructPointersEnconding = input.GongStructPointersEnconding

	query = db.Model(&gongstructDB).Updates(gongstructDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongstructNew := new(models.GongStruct)
	gongstructDB.CopyBasicFieldsToGongStruct(gongstructNew)

	// get stage instance from DB instance, and call callback function
	gongstructOld := (*orm.BackRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr)[gongstructDB.ID]
	if gongstructOld != nil {
		models.AfterUpdateFromFront(&models.Stage, gongstructOld, gongstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongstructDB
	c.JSON(http.StatusOK, gongstructDB)
}

// DeleteGongStruct
//
// swagger:route DELETE /gongstructs/{ID} gongstructs deleteGongStruct
//
// # Delete a gongstruct
//
// default: genericError
//
//	200: gongstructDBResponse
func DeleteGongStruct(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStruct.GetDB()

	// Get model if exist
	var gongstructDB orm.GongStructDB
	if err := db.First(&gongstructDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongstructDeleted := new(models.GongStruct)
	gongstructDB.CopyBasicFieldsToGongStruct(gongstructDeleted)

	// get stage instance from DB instance, and call callback function
	gongstructStaged := (*orm.BackRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr)[gongstructDB.ID]
	if gongstructStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, gongstructStaged, gongstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}