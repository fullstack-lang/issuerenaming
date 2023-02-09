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

	"github.com/fullstack-lang/issuerenaming/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Waldo_sql sql.NullBool
var dummy_Waldo_time time.Duration
var dummy_Waldo_sort sort.Float64Slice

// WaldoAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model waldoAPI
type WaldoAPI struct {
	gorm.Model

	models.Waldo

	// encoding of pointers
	WaldoPointersEnconding
}

// WaldoPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type WaldoPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Foo{}.Waldos []*Waldo
	Foo_WaldosDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Foo_WaldosDBID_Index sql.NullInt64
}

// WaldoDB describes a waldo in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model waldoDB
type WaldoDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field waldoDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	WaldoPointersEnconding
}

// WaldoDBs arrays waldoDBs
// swagger:response waldoDBsResponse
type WaldoDBs []WaldoDB

// WaldoDBResponse provides response
// swagger:response waldoDBResponse
type WaldoDBResponse struct {
	WaldoDB
}

// WaldoWOP is a Waldo without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type WaldoWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Waldo_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoWaldoStruct struct {
	// stores WaldoDB according to their gorm ID
	Map_WaldoDBID_WaldoDB *map[uint]*WaldoDB

	// stores WaldoDB ID according to Waldo address
	Map_WaldoPtr_WaldoDBID *map[*models.Waldo]uint

	// stores Waldo according to their gorm ID
	Map_WaldoDBID_WaldoPtr *map[uint]*models.Waldo

	db *gorm.DB
}

func (backRepoWaldo *BackRepoWaldoStruct) GetDB() *gorm.DB {
	return backRepoWaldo.db
}

// GetWaldoDBFromWaldoPtr is a handy function to access the back repo instance from the stage instance
func (backRepoWaldo *BackRepoWaldoStruct) GetWaldoDBFromWaldoPtr(waldo *models.Waldo) (waldoDB *WaldoDB) {
	id := (*backRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo]
	waldoDB = (*backRepoWaldo.Map_WaldoDBID_WaldoDB)[id]
	return
}

// BackRepoWaldo.Init set up the BackRepo of the Waldo
func (backRepoWaldo *BackRepoWaldoStruct) Init(db *gorm.DB) (Error error) {

	if backRepoWaldo.Map_WaldoDBID_WaldoPtr != nil {
		err := errors.New("In Init, backRepoWaldo.Map_WaldoDBID_WaldoPtr should be nil")
		return err
	}

	if backRepoWaldo.Map_WaldoDBID_WaldoDB != nil {
		err := errors.New("In Init, backRepoWaldo.Map_WaldoDBID_WaldoDB should be nil")
		return err
	}

	if backRepoWaldo.Map_WaldoPtr_WaldoDBID != nil {
		err := errors.New("In Init, backRepoWaldo.Map_WaldoPtr_WaldoDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Waldo, 0)
	backRepoWaldo.Map_WaldoDBID_WaldoPtr = &tmp

	tmpDB := make(map[uint]*WaldoDB, 0)
	backRepoWaldo.Map_WaldoDBID_WaldoDB = &tmpDB

	tmpID := make(map[*models.Waldo]uint, 0)
	backRepoWaldo.Map_WaldoPtr_WaldoDBID = &tmpID

	backRepoWaldo.db = db
	return
}

// BackRepoWaldo.CommitPhaseOne commits all staged instances of Waldo to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoWaldo *BackRepoWaldoStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for waldo := range stage.Waldos {
		backRepoWaldo.CommitPhaseOneInstance(waldo)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, waldo := range *backRepoWaldo.Map_WaldoDBID_WaldoPtr {
		if _, ok := stage.Waldos[waldo]; !ok {
			backRepoWaldo.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoWaldo.CommitDeleteInstance commits deletion of Waldo to the BackRepo
func (backRepoWaldo *BackRepoWaldoStruct) CommitDeleteInstance(id uint) (Error error) {

	waldo := (*backRepoWaldo.Map_WaldoDBID_WaldoPtr)[id]

	// waldo is not staged anymore, remove waldoDB
	waldoDB := (*backRepoWaldo.Map_WaldoDBID_WaldoDB)[id]
	query := backRepoWaldo.db.Unscoped().Delete(&waldoDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoWaldo.Map_WaldoPtr_WaldoDBID), waldo)
	delete((*backRepoWaldo.Map_WaldoDBID_WaldoPtr), id)
	delete((*backRepoWaldo.Map_WaldoDBID_WaldoDB), id)

	return
}

// BackRepoWaldo.CommitPhaseOneInstance commits waldo staged instances of Waldo to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoWaldo *BackRepoWaldoStruct) CommitPhaseOneInstance(waldo *models.Waldo) (Error error) {

	// check if the waldo is not commited yet
	if _, ok := (*backRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo]; ok {
		return
	}

	// initiate waldo
	var waldoDB WaldoDB
	waldoDB.CopyBasicFieldsFromWaldo(waldo)

	query := backRepoWaldo.db.Create(&waldoDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo] = waldoDB.ID
	(*backRepoWaldo.Map_WaldoDBID_WaldoPtr)[waldoDB.ID] = waldo
	(*backRepoWaldo.Map_WaldoDBID_WaldoDB)[waldoDB.ID] = &waldoDB

	return
}

// BackRepoWaldo.CommitPhaseTwo commits all staged instances of Waldo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWaldo *BackRepoWaldoStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, waldo := range *backRepoWaldo.Map_WaldoDBID_WaldoPtr {
		backRepoWaldo.CommitPhaseTwoInstance(backRepo, idx, waldo)
	}

	return
}

// BackRepoWaldo.CommitPhaseTwoInstance commits {{structname }} of models.Waldo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWaldo *BackRepoWaldoStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, waldo *models.Waldo) (Error error) {

	// fetch matching waldoDB
	if waldoDB, ok := (*backRepoWaldo.Map_WaldoDBID_WaldoDB)[idx]; ok {

		waldoDB.CopyBasicFieldsFromWaldo(waldo)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoWaldo.db.Save(&waldoDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Waldo intance %s", waldo.Name))
		return err
	}

	return
}

// BackRepoWaldo.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoWaldo *BackRepoWaldoStruct) CheckoutPhaseOne() (Error error) {

	waldoDBArray := make([]WaldoDB, 0)
	query := backRepoWaldo.db.Find(&waldoDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	waldoInstancesToBeRemovedFromTheStage := make(map[*models.Waldo]any)
	for key, value := range models.Stage.Waldos {
		waldoInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, waldoDB := range waldoDBArray {
		backRepoWaldo.CheckoutPhaseOneInstance(&waldoDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		waldo, ok := (*backRepoWaldo.Map_WaldoDBID_WaldoPtr)[waldoDB.ID]
		if ok {
			delete(waldoInstancesToBeRemovedFromTheStage, waldo)
		}
	}

	// remove from stage and back repo's 3 maps all waldos that are not in the checkout
	for waldo := range waldoInstancesToBeRemovedFromTheStage {
		waldo.Unstage()

		// remove instance from the back repo 3 maps
		waldoID := (*backRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo]
		delete((*backRepoWaldo.Map_WaldoPtr_WaldoDBID), waldo)
		delete((*backRepoWaldo.Map_WaldoDBID_WaldoDB), waldoID)
		delete((*backRepoWaldo.Map_WaldoDBID_WaldoPtr), waldoID)
	}

	return
}

// CheckoutPhaseOneInstance takes a waldoDB that has been found in the DB, updates the backRepo and stages the
// models version of the waldoDB
func (backRepoWaldo *BackRepoWaldoStruct) CheckoutPhaseOneInstance(waldoDB *WaldoDB) (Error error) {

	waldo, ok := (*backRepoWaldo.Map_WaldoDBID_WaldoPtr)[waldoDB.ID]
	if !ok {
		waldo = new(models.Waldo)

		(*backRepoWaldo.Map_WaldoDBID_WaldoPtr)[waldoDB.ID] = waldo
		(*backRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo] = waldoDB.ID

		// append model store with the new element
		waldo.Name = waldoDB.Name_Data.String
		waldo.Stage()
	}
	waldoDB.CopyBasicFieldsToWaldo(waldo)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	waldo.Stage()

	// preserve pointer to waldoDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_WaldoDBID_WaldoDB)[waldoDB hold variable pointers
	waldoDB_Data := *waldoDB
	preservedPtrToWaldo := &waldoDB_Data
	(*backRepoWaldo.Map_WaldoDBID_WaldoDB)[waldoDB.ID] = preservedPtrToWaldo

	return
}

// BackRepoWaldo.CheckoutPhaseTwo Checkouts all staged instances of Waldo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWaldo *BackRepoWaldoStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, waldoDB := range *backRepoWaldo.Map_WaldoDBID_WaldoDB {
		backRepoWaldo.CheckoutPhaseTwoInstance(backRepo, waldoDB)
	}
	return
}

// BackRepoWaldo.CheckoutPhaseTwoInstance Checkouts staged instances of Waldo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWaldo *BackRepoWaldoStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, waldoDB *WaldoDB) (Error error) {

	waldo := (*backRepoWaldo.Map_WaldoDBID_WaldoPtr)[waldoDB.ID]
	_ = waldo // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitWaldo allows commit of a single waldo (if already staged)
func (backRepo *BackRepoStruct) CommitWaldo(waldo *models.Waldo) {
	backRepo.BackRepoWaldo.CommitPhaseOneInstance(waldo)
	if id, ok := (*backRepo.BackRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo]; ok {
		backRepo.BackRepoWaldo.CommitPhaseTwoInstance(backRepo, id, waldo)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitWaldo allows checkout of a single waldo (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutWaldo(waldo *models.Waldo) {
	// check if the waldo is staged
	if _, ok := (*backRepo.BackRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo]; ok {

		if id, ok := (*backRepo.BackRepoWaldo.Map_WaldoPtr_WaldoDBID)[waldo]; ok {
			var waldoDB WaldoDB
			waldoDB.ID = id

			if err := backRepo.BackRepoWaldo.db.First(&waldoDB, id).Error; err != nil {
				log.Panicln("CheckoutWaldo : Problem with getting object with id:", id)
			}
			backRepo.BackRepoWaldo.CheckoutPhaseOneInstance(&waldoDB)
			backRepo.BackRepoWaldo.CheckoutPhaseTwoInstance(backRepo, &waldoDB)
		}
	}
}

// CopyBasicFieldsFromWaldo
func (waldoDB *WaldoDB) CopyBasicFieldsFromWaldo(waldo *models.Waldo) {
	// insertion point for fields commit

	waldoDB.Name_Data.String = waldo.Name
	waldoDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromWaldoWOP
func (waldoDB *WaldoDB) CopyBasicFieldsFromWaldoWOP(waldo *WaldoWOP) {
	// insertion point for fields commit

	waldoDB.Name_Data.String = waldo.Name
	waldoDB.Name_Data.Valid = true
}

// CopyBasicFieldsToWaldo
func (waldoDB *WaldoDB) CopyBasicFieldsToWaldo(waldo *models.Waldo) {
	// insertion point for checkout of basic fields (back repo to stage)
	waldo.Name = waldoDB.Name_Data.String
}

// CopyBasicFieldsToWaldoWOP
func (waldoDB *WaldoDB) CopyBasicFieldsToWaldoWOP(waldo *WaldoWOP) {
	waldo.ID = int(waldoDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	waldo.Name = waldoDB.Name_Data.String
}

// Backup generates a json file from a slice of all WaldoDB instances in the backrepo
func (backRepoWaldo *BackRepoWaldoStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "WaldoDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*WaldoDB, 0)
	for _, waldoDB := range *backRepoWaldo.Map_WaldoDBID_WaldoDB {
		forBackup = append(forBackup, waldoDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Waldo ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Waldo file", err.Error())
	}
}

// Backup generates a json file from a slice of all WaldoDB instances in the backrepo
func (backRepoWaldo *BackRepoWaldoStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*WaldoDB, 0)
	for _, waldoDB := range *backRepoWaldo.Map_WaldoDBID_WaldoDB {
		forBackup = append(forBackup, waldoDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Waldo")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Waldo_Fields, -1)
	for _, waldoDB := range forBackup {

		var waldoWOP WaldoWOP
		waldoDB.CopyBasicFieldsToWaldoWOP(&waldoWOP)

		row := sh.AddRow()
		row.WriteStruct(&waldoWOP, -1)
	}
}

// RestoreXL from the "Waldo" sheet all WaldoDB instances
func (backRepoWaldo *BackRepoWaldoStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoWaldoid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Waldo"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoWaldo.rowVisitorWaldo)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoWaldo *BackRepoWaldoStruct) rowVisitorWaldo(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var waldoWOP WaldoWOP
		row.ReadStruct(&waldoWOP)

		// add the unmarshalled struct to the stage
		waldoDB := new(WaldoDB)
		waldoDB.CopyBasicFieldsFromWaldoWOP(&waldoWOP)

		waldoDB_ID_atBackupTime := waldoDB.ID
		waldoDB.ID = 0
		query := backRepoWaldo.db.Create(waldoDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoWaldo.Map_WaldoDBID_WaldoDB)[waldoDB.ID] = waldoDB
		BackRepoWaldoid_atBckpTime_newID[waldoDB_ID_atBackupTime] = waldoDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "WaldoDB.json" in dirPath that stores an array
// of WaldoDB and stores it in the database
// the map BackRepoWaldoid_atBckpTime_newID is updated accordingly
func (backRepoWaldo *BackRepoWaldoStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoWaldoid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "WaldoDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Waldo file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*WaldoDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_WaldoDBID_WaldoDB
	for _, waldoDB := range forRestore {

		waldoDB_ID_atBackupTime := waldoDB.ID
		waldoDB.ID = 0
		query := backRepoWaldo.db.Create(waldoDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoWaldo.Map_WaldoDBID_WaldoDB)[waldoDB.ID] = waldoDB
		BackRepoWaldoid_atBckpTime_newID[waldoDB_ID_atBackupTime] = waldoDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Waldo file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Waldo>id_atBckpTime_newID
// to compute new index
func (backRepoWaldo *BackRepoWaldoStruct) RestorePhaseTwo() {

	for _, waldoDB := range *backRepoWaldo.Map_WaldoDBID_WaldoDB {

		// next line of code is to avert unused variable compilation error
		_ = waldoDB

		// insertion point for reindexing pointers encoding
		// This reindex waldo.Waldos
		if waldoDB.Foo_WaldosDBID.Int64 != 0 {
			waldoDB.Foo_WaldosDBID.Int64 =
				int64(BackRepoFooid_atBckpTime_newID[uint(waldoDB.Foo_WaldosDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoWaldo.db.Model(waldoDB).Updates(*waldoDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoWaldoid_atBckpTime_newID map[uint]uint
