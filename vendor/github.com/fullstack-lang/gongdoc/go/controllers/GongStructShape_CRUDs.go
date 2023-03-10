// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongStructShape__dummysDeclaration__ models.GongStructShape
var __GongStructShape_time__dummyDeclaration time.Duration

// An GongStructShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongStructShape updateGongStructShape deleteGongStructShape
type GongStructShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongStructShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongStructShape updateGongStructShape
type GongStructShapeInput struct {
	// The GongStructShape to submit or modify
	// in: body
	GongStructShape *orm.GongStructShapeAPI
}

// GetGongStructShapes
//
// swagger:route GET /gongstructshapes gongstructshapes getGongStructShapes
//
// # Get all gongstructshapes
//
// Responses:
// default: genericError
//
//	200: gongstructshapeDBResponse
func GetGongStructShapes(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStructShape.GetDB()

	// source slice
	var gongstructshapeDBs []orm.GongStructShapeDB
	query := db.Find(&gongstructshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongstructshapeAPIs := make([]orm.GongStructShapeAPI, 0)

	// for each gongstructshape, update fields from the database nullable fields
	for idx := range gongstructshapeDBs {
		gongstructshapeDB := &gongstructshapeDBs[idx]
		_ = gongstructshapeDB
		var gongstructshapeAPI orm.GongStructShapeAPI

		// insertion point for updating fields
		gongstructshapeAPI.ID = gongstructshapeDB.ID
		gongstructshapeDB.CopyBasicFieldsToGongStructShape(&gongstructshapeAPI.GongStructShape)
		gongstructshapeAPI.GongStructShapePointersEnconding = gongstructshapeDB.GongStructShapePointersEnconding
		gongstructshapeAPIs = append(gongstructshapeAPIs, gongstructshapeAPI)
	}

	c.JSON(http.StatusOK, gongstructshapeAPIs)
}

// PostGongStructShape
//
// swagger:route POST /gongstructshapes gongstructshapes postGongStructShape
//
// Creates a gongstructshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostGongStructShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStructShape.GetDB()

	// Validate input
	var input orm.GongStructShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongstructshape
	gongstructshapeDB := orm.GongStructShapeDB{}
	gongstructshapeDB.GongStructShapePointersEnconding = input.GongStructShapePointersEnconding
	gongstructshapeDB.CopyBasicFieldsFromGongStructShape(&input.GongStructShape)

	query := db.Create(&gongstructshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoGongStructShape.CheckoutPhaseOneInstance(&gongstructshapeDB)
	gongstructshape := (*orm.BackRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr)[gongstructshapeDB.ID]

	if gongstructshape != nil {
		models.AfterCreateFromFront(&models.Stage, gongstructshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongstructshapeDB)
}

// GetGongStructShape
//
// swagger:route GET /gongstructshapes/{ID} gongstructshapes getGongStructShape
//
// Gets the details for a gongstructshape.
//
// Responses:
// default: genericError
//
//	200: gongstructshapeDBResponse
func GetGongStructShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStructShape.GetDB()

	// Get gongstructshapeDB in DB
	var gongstructshapeDB orm.GongStructShapeDB
	if err := db.First(&gongstructshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongstructshapeAPI orm.GongStructShapeAPI
	gongstructshapeAPI.ID = gongstructshapeDB.ID
	gongstructshapeAPI.GongStructShapePointersEnconding = gongstructshapeDB.GongStructShapePointersEnconding
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(&gongstructshapeAPI.GongStructShape)

	c.JSON(http.StatusOK, gongstructshapeAPI)
}

// UpdateGongStructShape
//
// swagger:route PATCH /gongstructshapes/{ID} gongstructshapes updateGongStructShape
//
// # Update a gongstructshape
//
// Responses:
// default: genericError
//
//	200: gongstructshapeDBResponse
func UpdateGongStructShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStructShape.GetDB()

	// Get model if exist
	var gongstructshapeDB orm.GongStructShapeDB

	// fetch the gongstructshape
	query := db.First(&gongstructshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.GongStructShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	gongstructshapeDB.CopyBasicFieldsFromGongStructShape(&input.GongStructShape)
	gongstructshapeDB.GongStructShapePointersEnconding = input.GongStructShapePointersEnconding

	query = db.Model(&gongstructshapeDB).Updates(gongstructshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongstructshapeNew := new(models.GongStructShape)
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(gongstructshapeNew)

	// get stage instance from DB instance, and call callback function
	gongstructshapeOld := (*orm.BackRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr)[gongstructshapeDB.ID]
	if gongstructshapeOld != nil {
		models.AfterUpdateFromFront(&models.Stage, gongstructshapeOld, gongstructshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongstructshapeDB
	c.JSON(http.StatusOK, gongstructshapeDB)
}

// DeleteGongStructShape
//
// swagger:route DELETE /gongstructshapes/{ID} gongstructshapes deleteGongStructShape
//
// # Delete a gongstructshape
//
// default: genericError
//
//	200: gongstructshapeDBResponse
func DeleteGongStructShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongStructShape.GetDB()

	// Get model if exist
	var gongstructshapeDB orm.GongStructShapeDB
	if err := db.First(&gongstructshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongstructshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongstructshapeDeleted := new(models.GongStructShape)
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(gongstructshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongstructshapeStaged := (*orm.BackRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr)[gongstructshapeDB.ID]
	if gongstructshapeStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, gongstructshapeStaged, gongstructshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
