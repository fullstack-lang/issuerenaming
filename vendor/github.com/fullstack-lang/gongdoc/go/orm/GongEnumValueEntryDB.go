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
var dummy_GongEnumValueEntry_sql sql.NullBool
var dummy_GongEnumValueEntry_time time.Duration
var dummy_GongEnumValueEntry_sort sort.Float64Slice

// GongEnumValueEntryAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongenumvalueentryAPI
type GongEnumValueEntryAPI struct {
	gorm.Model

	models.GongEnumValueEntry

	// encoding of pointers
	GongEnumValueEntryPointersEnconding
}

// GongEnumValueEntryPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongEnumValueEntryPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field GongEnumShape{}.GongEnumValueEntrys []*GongEnumValueEntry
	GongEnumShape_GongEnumValueEntrysDBID sql.NullInt64

	// implementation of the index of the withing the slice
	GongEnumShape_GongEnumValueEntrysDBID_Index sql.NullInt64
}

// GongEnumValueEntryDB describes a gongenumvalueentry in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongenumvalueentryDB
type GongEnumValueEntryDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongenumvalueentryDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongenumvalueentryDB.Identifier
	Identifier_Data sql.NullString
	// encoding of pointers
	GongEnumValueEntryPointersEnconding
}

// GongEnumValueEntryDBs arrays gongenumvalueentryDBs
// swagger:response gongenumvalueentryDBsResponse
type GongEnumValueEntryDBs []GongEnumValueEntryDB

// GongEnumValueEntryDBResponse provides response
// swagger:response gongenumvalueentryDBResponse
type GongEnumValueEntryDBResponse struct {
	GongEnumValueEntryDB
}

// GongEnumValueEntryWOP is a GongEnumValueEntry without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongEnumValueEntryWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Identifier string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var GongEnumValueEntry_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Identifier",
}

type BackRepoGongEnumValueEntryStruct struct {
	// stores GongEnumValueEntryDB according to their gorm ID
	Map_GongEnumValueEntryDBID_GongEnumValueEntryDB *map[uint]*GongEnumValueEntryDB

	// stores GongEnumValueEntryDB ID according to GongEnumValueEntry address
	Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID *map[*models.GongEnumValueEntry]uint

	// stores GongEnumValueEntry according to their gorm ID
	Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr *map[uint]*models.GongEnumValueEntry

	db *gorm.DB
}

func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) GetDB() *gorm.DB {
	return backRepoGongEnumValueEntry.db
}

// GetGongEnumValueEntryDBFromGongEnumValueEntryPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) GetGongEnumValueEntryDBFromGongEnumValueEntryPtr(gongenumvalueentry *models.GongEnumValueEntry) (gongenumvalueentryDB *GongEnumValueEntryDB) {
	id := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry]
	gongenumvalueentryDB = (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[id]
	return
}

// BackRepoGongEnumValueEntry.Init set up the BackRepo of the GongEnumValueEntry
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) Init(db *gorm.DB) (Error error) {

	if backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr != nil {
		err := errors.New("In Init, backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr should be nil")
		return err
	}

	if backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB != nil {
		err := errors.New("In Init, backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB should be nil")
		return err
	}

	if backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID != nil {
		err := errors.New("In Init, backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.GongEnumValueEntry, 0)
	backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr = &tmp

	tmpDB := make(map[uint]*GongEnumValueEntryDB, 0)
	backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB = &tmpDB

	tmpID := make(map[*models.GongEnumValueEntry]uint, 0)
	backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID = &tmpID

	backRepoGongEnumValueEntry.db = db
	return
}

// BackRepoGongEnumValueEntry.CommitPhaseOne commits all staged instances of GongEnumValueEntry to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongenumvalueentry := range stage.GongEnumValueEntrys {
		backRepoGongEnumValueEntry.CommitPhaseOneInstance(gongenumvalueentry)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongenumvalueentry := range *backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr {
		if _, ok := stage.GongEnumValueEntrys[gongenumvalueentry]; !ok {
			backRepoGongEnumValueEntry.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongEnumValueEntry.CommitDeleteInstance commits deletion of GongEnumValueEntry to the BackRepo
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CommitDeleteInstance(id uint) (Error error) {

	gongenumvalueentry := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[id]

	// gongenumvalueentry is not staged anymore, remove gongenumvalueentryDB
	gongenumvalueentryDB := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[id]
	query := backRepoGongEnumValueEntry.db.Unscoped().Delete(&gongenumvalueentryDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID), gongenumvalueentry)
	delete((*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr), id)
	delete((*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB), id)

	return
}

// BackRepoGongEnumValueEntry.CommitPhaseOneInstance commits gongenumvalueentry staged instances of GongEnumValueEntry to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CommitPhaseOneInstance(gongenumvalueentry *models.GongEnumValueEntry) (Error error) {

	// check if the gongenumvalueentry is not commited yet
	if _, ok := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry]; ok {
		return
	}

	// initiate gongenumvalueentry
	var gongenumvalueentryDB GongEnumValueEntryDB
	gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntry(gongenumvalueentry)

	query := backRepoGongEnumValueEntry.db.Create(&gongenumvalueentryDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry] = gongenumvalueentryDB.ID
	(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID] = gongenumvalueentry
	(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[gongenumvalueentryDB.ID] = &gongenumvalueentryDB

	return
}

// BackRepoGongEnumValueEntry.CommitPhaseTwo commits all staged instances of GongEnumValueEntry to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongenumvalueentry := range *backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr {
		backRepoGongEnumValueEntry.CommitPhaseTwoInstance(backRepo, idx, gongenumvalueentry)
	}

	return
}

// BackRepoGongEnumValueEntry.CommitPhaseTwoInstance commits {{structname }} of models.GongEnumValueEntry to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongenumvalueentry *models.GongEnumValueEntry) (Error error) {

	// fetch matching gongenumvalueentryDB
	if gongenumvalueentryDB, ok := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[idx]; ok {

		gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntry(gongenumvalueentry)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoGongEnumValueEntry.db.Save(&gongenumvalueentryDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongEnumValueEntry intance %s", gongenumvalueentry.Name))
		return err
	}

	return
}

// BackRepoGongEnumValueEntry.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CheckoutPhaseOne() (Error error) {

	gongenumvalueentryDBArray := make([]GongEnumValueEntryDB, 0)
	query := backRepoGongEnumValueEntry.db.Find(&gongenumvalueentryDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongenumvalueentryInstancesToBeRemovedFromTheStage := make(map[*models.GongEnumValueEntry]any)
	for key, value := range models.Stage.GongEnumValueEntrys {
		gongenumvalueentryInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongenumvalueentryDB := range gongenumvalueentryDBArray {
		backRepoGongEnumValueEntry.CheckoutPhaseOneInstance(&gongenumvalueentryDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongenumvalueentry, ok := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID]
		if ok {
			delete(gongenumvalueentryInstancesToBeRemovedFromTheStage, gongenumvalueentry)
		}
	}

	// remove from stage and back repo's 3 maps all gongenumvalueentrys that are not in the checkout
	for gongenumvalueentry := range gongenumvalueentryInstancesToBeRemovedFromTheStage {
		gongenumvalueentry.Unstage()

		// remove instance from the back repo 3 maps
		gongenumvalueentryID := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry]
		delete((*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID), gongenumvalueentry)
		delete((*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB), gongenumvalueentryID)
		delete((*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr), gongenumvalueentryID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongenumvalueentryDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongenumvalueentryDB
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CheckoutPhaseOneInstance(gongenumvalueentryDB *GongEnumValueEntryDB) (Error error) {

	gongenumvalueentry, ok := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID]
	if !ok {
		gongenumvalueentry = new(models.GongEnumValueEntry)

		(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID] = gongenumvalueentry
		(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry] = gongenumvalueentryDB.ID

		// append model store with the new element
		gongenumvalueentry.Name = gongenumvalueentryDB.Name_Data.String
		gongenumvalueentry.Stage()
	}
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(gongenumvalueentry)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongenumvalueentry.Stage()

	// preserve pointer to gongenumvalueentryDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[gongenumvalueentryDB hold variable pointers
	gongenumvalueentryDB_Data := *gongenumvalueentryDB
	preservedPtrToGongEnumValueEntry := &gongenumvalueentryDB_Data
	(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[gongenumvalueentryDB.ID] = preservedPtrToGongEnumValueEntry

	return
}

// BackRepoGongEnumValueEntry.CheckoutPhaseTwo Checkouts all staged instances of GongEnumValueEntry to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongenumvalueentryDB := range *backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB {
		backRepoGongEnumValueEntry.CheckoutPhaseTwoInstance(backRepo, gongenumvalueentryDB)
	}
	return
}

// BackRepoGongEnumValueEntry.CheckoutPhaseTwoInstance Checkouts staged instances of GongEnumValueEntry to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongenumvalueentryDB *GongEnumValueEntryDB) (Error error) {

	gongenumvalueentry := (*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID]
	_ = gongenumvalueentry // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitGongEnumValueEntry allows commit of a single gongenumvalueentry (if already staged)
func (backRepo *BackRepoStruct) CommitGongEnumValueEntry(gongenumvalueentry *models.GongEnumValueEntry) {
	backRepo.BackRepoGongEnumValueEntry.CommitPhaseOneInstance(gongenumvalueentry)
	if id, ok := (*backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry]; ok {
		backRepo.BackRepoGongEnumValueEntry.CommitPhaseTwoInstance(backRepo, id, gongenumvalueentry)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongEnumValueEntry allows checkout of a single gongenumvalueentry (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongEnumValueEntry(gongenumvalueentry *models.GongEnumValueEntry) {
	// check if the gongenumvalueentry is staged
	if _, ok := (*backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry]; ok {

		if id, ok := (*backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID)[gongenumvalueentry]; ok {
			var gongenumvalueentryDB GongEnumValueEntryDB
			gongenumvalueentryDB.ID = id

			if err := backRepo.BackRepoGongEnumValueEntry.db.First(&gongenumvalueentryDB, id).Error; err != nil {
				log.Panicln("CheckoutGongEnumValueEntry : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongEnumValueEntry.CheckoutPhaseOneInstance(&gongenumvalueentryDB)
			backRepo.BackRepoGongEnumValueEntry.CheckoutPhaseTwoInstance(backRepo, &gongenumvalueentryDB)
		}
	}
}

// CopyBasicFieldsFromGongEnumValueEntry
func (gongenumvalueentryDB *GongEnumValueEntryDB) CopyBasicFieldsFromGongEnumValueEntry(gongenumvalueentry *models.GongEnumValueEntry) {
	// insertion point for fields commit

	gongenumvalueentryDB.Name_Data.String = gongenumvalueentry.Name
	gongenumvalueentryDB.Name_Data.Valid = true

	gongenumvalueentryDB.Identifier_Data.String = gongenumvalueentry.Identifier
	gongenumvalueentryDB.Identifier_Data.Valid = true
}

// CopyBasicFieldsFromGongEnumValueEntryWOP
func (gongenumvalueentryDB *GongEnumValueEntryDB) CopyBasicFieldsFromGongEnumValueEntryWOP(gongenumvalueentry *GongEnumValueEntryWOP) {
	// insertion point for fields commit

	gongenumvalueentryDB.Name_Data.String = gongenumvalueentry.Name
	gongenumvalueentryDB.Name_Data.Valid = true

	gongenumvalueentryDB.Identifier_Data.String = gongenumvalueentry.Identifier
	gongenumvalueentryDB.Identifier_Data.Valid = true
}

// CopyBasicFieldsToGongEnumValueEntry
func (gongenumvalueentryDB *GongEnumValueEntryDB) CopyBasicFieldsToGongEnumValueEntry(gongenumvalueentry *models.GongEnumValueEntry) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumvalueentry.Name = gongenumvalueentryDB.Name_Data.String
	gongenumvalueentry.Identifier = gongenumvalueentryDB.Identifier_Data.String
}

// CopyBasicFieldsToGongEnumValueEntryWOP
func (gongenumvalueentryDB *GongEnumValueEntryDB) CopyBasicFieldsToGongEnumValueEntryWOP(gongenumvalueentry *GongEnumValueEntryWOP) {
	gongenumvalueentry.ID = int(gongenumvalueentryDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumvalueentry.Name = gongenumvalueentryDB.Name_Data.String
	gongenumvalueentry.Identifier = gongenumvalueentryDB.Identifier_Data.String
}

// Backup generates a json file from a slice of all GongEnumValueEntryDB instances in the backrepo
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongEnumValueEntryDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongEnumValueEntryDB, 0)
	for _, gongenumvalueentryDB := range *backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB {
		forBackup = append(forBackup, gongenumvalueentryDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongEnumValueEntry ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongEnumValueEntry file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongEnumValueEntryDB instances in the backrepo
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongEnumValueEntryDB, 0)
	for _, gongenumvalueentryDB := range *backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB {
		forBackup = append(forBackup, gongenumvalueentryDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongEnumValueEntry")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongEnumValueEntry_Fields, -1)
	for _, gongenumvalueentryDB := range forBackup {

		var gongenumvalueentryWOP GongEnumValueEntryWOP
		gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntryWOP(&gongenumvalueentryWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongenumvalueentryWOP, -1)
	}
}

// RestoreXL from the "GongEnumValueEntry" sheet all GongEnumValueEntryDB instances
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongEnumValueEntryid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongEnumValueEntry"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongEnumValueEntry.rowVisitorGongEnumValueEntry)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) rowVisitorGongEnumValueEntry(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongenumvalueentryWOP GongEnumValueEntryWOP
		row.ReadStruct(&gongenumvalueentryWOP)

		// add the unmarshalled struct to the stage
		gongenumvalueentryDB := new(GongEnumValueEntryDB)
		gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntryWOP(&gongenumvalueentryWOP)

		gongenumvalueentryDB_ID_atBackupTime := gongenumvalueentryDB.ID
		gongenumvalueentryDB.ID = 0
		query := backRepoGongEnumValueEntry.db.Create(gongenumvalueentryDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[gongenumvalueentryDB.ID] = gongenumvalueentryDB
		BackRepoGongEnumValueEntryid_atBckpTime_newID[gongenumvalueentryDB_ID_atBackupTime] = gongenumvalueentryDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongEnumValueEntryDB.json" in dirPath that stores an array
// of GongEnumValueEntryDB and stores it in the database
// the map BackRepoGongEnumValueEntryid_atBckpTime_newID is updated accordingly
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongEnumValueEntryid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongEnumValueEntryDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongEnumValueEntry file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongEnumValueEntryDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongEnumValueEntryDBID_GongEnumValueEntryDB
	for _, gongenumvalueentryDB := range forRestore {

		gongenumvalueentryDB_ID_atBackupTime := gongenumvalueentryDB.ID
		gongenumvalueentryDB.ID = 0
		query := backRepoGongEnumValueEntry.db.Create(gongenumvalueentryDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB)[gongenumvalueentryDB.ID] = gongenumvalueentryDB
		BackRepoGongEnumValueEntryid_atBckpTime_newID[gongenumvalueentryDB_ID_atBackupTime] = gongenumvalueentryDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongEnumValueEntry file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongEnumValueEntry>id_atBckpTime_newID
// to compute new index
func (backRepoGongEnumValueEntry *BackRepoGongEnumValueEntryStruct) RestorePhaseTwo() {

	for _, gongenumvalueentryDB := range *backRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB {

		// next line of code is to avert unused variable compilation error
		_ = gongenumvalueentryDB

		// insertion point for reindexing pointers encoding
		// This reindex gongenumvalueentry.GongEnumValueEntrys
		if gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64 != 0 {
			gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64 =
				int64(BackRepoGongEnumShapeid_atBckpTime_newID[uint(gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoGongEnumValueEntry.db.Model(gongenumvalueentryDB).Updates(*gongenumvalueentryDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongEnumValueEntryid_atBckpTime_newID map[uint]uint
