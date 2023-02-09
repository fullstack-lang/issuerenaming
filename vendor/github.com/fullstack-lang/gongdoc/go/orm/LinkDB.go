// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Link_sql sql.NullBool
var dummy_Link_time time.Duration
var dummy_Link_sort sort.Float64Slice

// LinkAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model linkAPI
type LinkAPI struct {
	gorm.Model

	models.Link

	// encoding of pointers
	LinkPointersEnconding
}

// LinkPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LinkPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field Middlevertice is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	MiddleverticeID sql.NullInt64

	// Implementation of a reverse ID for field GongStructShape{}.Links []*Link
	GongStructShape_LinksDBID sql.NullInt64

	// implementation of the index of the withing the slice
	GongStructShape_LinksDBID_Index sql.NullInt64
}

// LinkDB describes a link in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model linkDB
type LinkDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field linkDB.Name
	Name_Data sql.NullString

	// Declation for basic field linkDB.Structname
	Structname_Data sql.NullString

	// Declation for basic field linkDB.Identifier
	Identifier_Data sql.NullString

	// Declation for basic field linkDB.Fieldtypename
	Fieldtypename_Data sql.NullString

	// Declation for basic field linkDB.TargetMultiplicity
	TargetMultiplicity_Data sql.NullString

	// Declation for basic field linkDB.SourceMultiplicity
	SourceMultiplicity_Data sql.NullString
	// encoding of pointers
	LinkPointersEnconding
}

// LinkDBs arrays linkDBs
// swagger:response linkDBsResponse
type LinkDBs []LinkDB

// LinkDBResponse provides response
// swagger:response linkDBResponse
type LinkDBResponse struct {
	LinkDB
}

// LinkWOP is a Link without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LinkWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Structname string `xlsx:"2"`

	Identifier string `xlsx:"3"`

	Fieldtypename string `xlsx:"4"`

	TargetMultiplicity models.MultiplicityType `xlsx:"5"`

	SourceMultiplicity models.MultiplicityType `xlsx:"6"`
	// insertion for WOP pointer fields
}

var Link_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Structname",
	"Identifier",
	"Fieldtypename",
	"TargetMultiplicity",
	"SourceMultiplicity",
}

type BackRepoLinkStruct struct {
	// stores LinkDB according to their gorm ID
	Map_LinkDBID_LinkDB *map[uint]*LinkDB

	// stores LinkDB ID according to Link address
	Map_LinkPtr_LinkDBID *map[*models.Link]uint

	// stores Link according to their gorm ID
	Map_LinkDBID_LinkPtr *map[uint]*models.Link

	db *gorm.DB
}

func (backRepoLink *BackRepoLinkStruct) GetDB() *gorm.DB {
	return backRepoLink.db
}

// GetLinkDBFromLinkPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLink *BackRepoLinkStruct) GetLinkDBFromLinkPtr(link *models.Link) (linkDB *LinkDB) {
	id := (*backRepoLink.Map_LinkPtr_LinkDBID)[link]
	linkDB = (*backRepoLink.Map_LinkDBID_LinkDB)[id]
	return
}

// BackRepoLink.Init set up the BackRepo of the Link
func (backRepoLink *BackRepoLinkStruct) Init(db *gorm.DB) (Error error) {

	if backRepoLink.Map_LinkDBID_LinkPtr != nil {
		err := errors.New("In Init, backRepoLink.Map_LinkDBID_LinkPtr should be nil")
		return err
	}

	if backRepoLink.Map_LinkDBID_LinkDB != nil {
		err := errors.New("In Init, backRepoLink.Map_LinkDBID_LinkDB should be nil")
		return err
	}

	if backRepoLink.Map_LinkPtr_LinkDBID != nil {
		err := errors.New("In Init, backRepoLink.Map_LinkPtr_LinkDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Link, 0)
	backRepoLink.Map_LinkDBID_LinkPtr = &tmp

	tmpDB := make(map[uint]*LinkDB, 0)
	backRepoLink.Map_LinkDBID_LinkDB = &tmpDB

	tmpID := make(map[*models.Link]uint, 0)
	backRepoLink.Map_LinkPtr_LinkDBID = &tmpID

	backRepoLink.db = db
	return
}

// BackRepoLink.CommitPhaseOne commits all staged instances of Link to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLink *BackRepoLinkStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for link := range stage.Links {
		backRepoLink.CommitPhaseOneInstance(link)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, link := range *backRepoLink.Map_LinkDBID_LinkPtr {
		if _, ok := stage.Links[link]; !ok {
			backRepoLink.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLink.CommitDeleteInstance commits deletion of Link to the BackRepo
func (backRepoLink *BackRepoLinkStruct) CommitDeleteInstance(id uint) (Error error) {

	link := (*backRepoLink.Map_LinkDBID_LinkPtr)[id]

	// link is not staged anymore, remove linkDB
	linkDB := (*backRepoLink.Map_LinkDBID_LinkDB)[id]
	query := backRepoLink.db.Unscoped().Delete(&linkDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoLink.Map_LinkPtr_LinkDBID), link)
	delete((*backRepoLink.Map_LinkDBID_LinkPtr), id)
	delete((*backRepoLink.Map_LinkDBID_LinkDB), id)

	return
}

// BackRepoLink.CommitPhaseOneInstance commits link staged instances of Link to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLink *BackRepoLinkStruct) CommitPhaseOneInstance(link *models.Link) (Error error) {

	// check if the link is not commited yet
	if _, ok := (*backRepoLink.Map_LinkPtr_LinkDBID)[link]; ok {
		return
	}

	// initiate link
	var linkDB LinkDB
	linkDB.CopyBasicFieldsFromLink(link)

	query := backRepoLink.db.Create(&linkDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoLink.Map_LinkPtr_LinkDBID)[link] = linkDB.ID
	(*backRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID] = link
	(*backRepoLink.Map_LinkDBID_LinkDB)[linkDB.ID] = &linkDB

	return
}

// BackRepoLink.CommitPhaseTwo commits all staged instances of Link to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLink *BackRepoLinkStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, link := range *backRepoLink.Map_LinkDBID_LinkPtr {
		backRepoLink.CommitPhaseTwoInstance(backRepo, idx, link)
	}

	return
}

// BackRepoLink.CommitPhaseTwoInstance commits {{structname }} of models.Link to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLink *BackRepoLinkStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, link *models.Link) (Error error) {

	// fetch matching linkDB
	if linkDB, ok := (*backRepoLink.Map_LinkDBID_LinkDB)[idx]; ok {

		linkDB.CopyBasicFieldsFromLink(link)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value link.Middlevertice translates to updating the link.MiddleverticeID
		linkDB.MiddleverticeID.Valid = true // allow for a 0 value (nil association)
		if link.Middlevertice != nil {
			if MiddleverticeId, ok := (*backRepo.BackRepoVertice.Map_VerticePtr_VerticeDBID)[link.Middlevertice]; ok {
				linkDB.MiddleverticeID.Int64 = int64(MiddleverticeId)
				linkDB.MiddleverticeID.Valid = true
			}
		}

		query := backRepoLink.db.Save(&linkDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Link intance %s", link.Name))
		return err
	}

	return
}

// BackRepoLink.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLink *BackRepoLinkStruct) CheckoutPhaseOne() (Error error) {

	linkDBArray := make([]LinkDB, 0)
	query := backRepoLink.db.Find(&linkDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	linkInstancesToBeRemovedFromTheStage := make(map[*models.Link]any)
	for key, value := range models.Stage.Links {
		linkInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, linkDB := range linkDBArray {
		backRepoLink.CheckoutPhaseOneInstance(&linkDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		link, ok := (*backRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID]
		if ok {
			delete(linkInstancesToBeRemovedFromTheStage, link)
		}
	}

	// remove from stage and back repo's 3 maps all links that are not in the checkout
	for link := range linkInstancesToBeRemovedFromTheStage {
		link.Unstage()

		// remove instance from the back repo 3 maps
		linkID := (*backRepoLink.Map_LinkPtr_LinkDBID)[link]
		delete((*backRepoLink.Map_LinkPtr_LinkDBID), link)
		delete((*backRepoLink.Map_LinkDBID_LinkDB), linkID)
		delete((*backRepoLink.Map_LinkDBID_LinkPtr), linkID)
	}

	return
}

// CheckoutPhaseOneInstance takes a linkDB that has been found in the DB, updates the backRepo and stages the
// models version of the linkDB
func (backRepoLink *BackRepoLinkStruct) CheckoutPhaseOneInstance(linkDB *LinkDB) (Error error) {

	link, ok := (*backRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID]
	if !ok {
		link = new(models.Link)

		(*backRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID] = link
		(*backRepoLink.Map_LinkPtr_LinkDBID)[link] = linkDB.ID

		// append model store with the new element
		link.Name = linkDB.Name_Data.String
		link.Stage()
	}
	linkDB.CopyBasicFieldsToLink(link)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	link.Stage()

	// preserve pointer to linkDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LinkDBID_LinkDB)[linkDB hold variable pointers
	linkDB_Data := *linkDB
	preservedPtrToLink := &linkDB_Data
	(*backRepoLink.Map_LinkDBID_LinkDB)[linkDB.ID] = preservedPtrToLink

	return
}

// BackRepoLink.CheckoutPhaseTwo Checkouts all staged instances of Link to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLink *BackRepoLinkStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, linkDB := range *backRepoLink.Map_LinkDBID_LinkDB {
		backRepoLink.CheckoutPhaseTwoInstance(backRepo, linkDB)
	}
	return
}

// BackRepoLink.CheckoutPhaseTwoInstance Checkouts staged instances of Link to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLink *BackRepoLinkStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, linkDB *LinkDB) (Error error) {

	link := (*backRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID]
	_ = link // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// Middlevertice field
	if linkDB.MiddleverticeID.Int64 != 0 {
		link.Middlevertice = (*backRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr)[uint(linkDB.MiddleverticeID.Int64)]
	}
	return
}

// CommitLink allows commit of a single link (if already staged)
func (backRepo *BackRepoStruct) CommitLink(link *models.Link) {
	backRepo.BackRepoLink.CommitPhaseOneInstance(link)
	if id, ok := (*backRepo.BackRepoLink.Map_LinkPtr_LinkDBID)[link]; ok {
		backRepo.BackRepoLink.CommitPhaseTwoInstance(backRepo, id, link)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLink allows checkout of a single link (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLink(link *models.Link) {
	// check if the link is staged
	if _, ok := (*backRepo.BackRepoLink.Map_LinkPtr_LinkDBID)[link]; ok {

		if id, ok := (*backRepo.BackRepoLink.Map_LinkPtr_LinkDBID)[link]; ok {
			var linkDB LinkDB
			linkDB.ID = id

			if err := backRepo.BackRepoLink.db.First(&linkDB, id).Error; err != nil {
				log.Panicln("CheckoutLink : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLink.CheckoutPhaseOneInstance(&linkDB)
			backRepo.BackRepoLink.CheckoutPhaseTwoInstance(backRepo, &linkDB)
		}
	}
}

// CopyBasicFieldsFromLink
func (linkDB *LinkDB) CopyBasicFieldsFromLink(link *models.Link) {
	// insertion point for fields commit

	linkDB.Name_Data.String = link.Name
	linkDB.Name_Data.Valid = true

	linkDB.Structname_Data.String = link.Structname
	linkDB.Structname_Data.Valid = true

	linkDB.Identifier_Data.String = link.Identifier
	linkDB.Identifier_Data.Valid = true

	linkDB.Fieldtypename_Data.String = link.Fieldtypename
	linkDB.Fieldtypename_Data.Valid = true

	linkDB.TargetMultiplicity_Data.String = link.TargetMultiplicity.ToString()
	linkDB.TargetMultiplicity_Data.Valid = true

	linkDB.SourceMultiplicity_Data.String = link.SourceMultiplicity.ToString()
	linkDB.SourceMultiplicity_Data.Valid = true
}

// CopyBasicFieldsFromLinkWOP
func (linkDB *LinkDB) CopyBasicFieldsFromLinkWOP(link *LinkWOP) {
	// insertion point for fields commit

	linkDB.Name_Data.String = link.Name
	linkDB.Name_Data.Valid = true

	linkDB.Structname_Data.String = link.Structname
	linkDB.Structname_Data.Valid = true

	linkDB.Identifier_Data.String = link.Identifier
	linkDB.Identifier_Data.Valid = true

	linkDB.Fieldtypename_Data.String = link.Fieldtypename
	linkDB.Fieldtypename_Data.Valid = true

	linkDB.TargetMultiplicity_Data.String = link.TargetMultiplicity.ToString()
	linkDB.TargetMultiplicity_Data.Valid = true

	linkDB.SourceMultiplicity_Data.String = link.SourceMultiplicity.ToString()
	linkDB.SourceMultiplicity_Data.Valid = true
}

// CopyBasicFieldsToLink
func (linkDB *LinkDB) CopyBasicFieldsToLink(link *models.Link) {
	// insertion point for checkout of basic fields (back repo to stage)
	link.Name = linkDB.Name_Data.String
	link.Structname = linkDB.Structname_Data.String
	link.Identifier = linkDB.Identifier_Data.String
	link.Fieldtypename = linkDB.Fieldtypename_Data.String
	link.TargetMultiplicity.FromString(linkDB.TargetMultiplicity_Data.String)
	link.SourceMultiplicity.FromString(linkDB.SourceMultiplicity_Data.String)
}

// CopyBasicFieldsToLinkWOP
func (linkDB *LinkDB) CopyBasicFieldsToLinkWOP(link *LinkWOP) {
	link.ID = int(linkDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	link.Name = linkDB.Name_Data.String
	link.Structname = linkDB.Structname_Data.String
	link.Identifier = linkDB.Identifier_Data.String
	link.Fieldtypename = linkDB.Fieldtypename_Data.String
	link.TargetMultiplicity.FromString(linkDB.TargetMultiplicity_Data.String)
	link.SourceMultiplicity.FromString(linkDB.SourceMultiplicity_Data.String)
}

// Backup generates a json file from a slice of all LinkDB instances in the backrepo
func (backRepoLink *BackRepoLinkStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LinkDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LinkDB, 0)
	for _, linkDB := range *backRepoLink.Map_LinkDBID_LinkDB {
		forBackup = append(forBackup, linkDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Link ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Link file", err.Error())
	}
}

// Backup generates a json file from a slice of all LinkDB instances in the backrepo
func (backRepoLink *BackRepoLinkStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LinkDB, 0)
	for _, linkDB := range *backRepoLink.Map_LinkDBID_LinkDB {
		forBackup = append(forBackup, linkDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Link")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Link_Fields, -1)
	for _, linkDB := range forBackup {

		var linkWOP LinkWOP
		linkDB.CopyBasicFieldsToLinkWOP(&linkWOP)

		row := sh.AddRow()
		row.WriteStruct(&linkWOP, -1)
	}
}

// RestoreXL from the "Link" sheet all LinkDB instances
func (backRepoLink *BackRepoLinkStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLinkid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Link"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLink.rowVisitorLink)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLink *BackRepoLinkStruct) rowVisitorLink(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var linkWOP LinkWOP
		row.ReadStruct(&linkWOP)

		// add the unmarshalled struct to the stage
		linkDB := new(LinkDB)
		linkDB.CopyBasicFieldsFromLinkWOP(&linkWOP)

		linkDB_ID_atBackupTime := linkDB.ID
		linkDB.ID = 0
		query := backRepoLink.db.Create(linkDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLink.Map_LinkDBID_LinkDB)[linkDB.ID] = linkDB
		BackRepoLinkid_atBckpTime_newID[linkDB_ID_atBackupTime] = linkDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LinkDB.json" in dirPath that stores an array
// of LinkDB and stores it in the database
// the map BackRepoLinkid_atBckpTime_newID is updated accordingly
func (backRepoLink *BackRepoLinkStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLinkid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LinkDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Link file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LinkDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LinkDBID_LinkDB
	for _, linkDB := range forRestore {

		linkDB_ID_atBackupTime := linkDB.ID
		linkDB.ID = 0
		query := backRepoLink.db.Create(linkDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLink.Map_LinkDBID_LinkDB)[linkDB.ID] = linkDB
		BackRepoLinkid_atBckpTime_newID[linkDB_ID_atBackupTime] = linkDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Link file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Link>id_atBckpTime_newID
// to compute new index
func (backRepoLink *BackRepoLinkStruct) RestorePhaseTwo() {

	for _, linkDB := range *backRepoLink.Map_LinkDBID_LinkDB {

		// next line of code is to avert unused variable compilation error
		_ = linkDB

		// insertion point for reindexing pointers encoding
		// reindexing Middlevertice field
		if linkDB.MiddleverticeID.Int64 != 0 {
			linkDB.MiddleverticeID.Int64 = int64(BackRepoVerticeid_atBckpTime_newID[uint(linkDB.MiddleverticeID.Int64)])
			linkDB.MiddleverticeID.Valid = true
		}

		// This reindex link.Links
		if linkDB.GongStructShape_LinksDBID.Int64 != 0 {
			linkDB.GongStructShape_LinksDBID.Int64 =
				int64(BackRepoGongStructShapeid_atBckpTime_newID[uint(linkDB.GongStructShape_LinksDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoLink.db.Model(linkDB).Updates(*linkDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLinkid_atBckpTime_newID map[uint]uint
