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
var dummy_Umlsc_sql sql.NullBool
var dummy_Umlsc_time time.Duration
var dummy_Umlsc_sort sort.Float64Slice

// UmlscAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model umlscAPI
type UmlscAPI struct {
	gorm.Model

	models.Umlsc

	// encoding of pointers
	UmlscPointersEnconding
}

// UmlscPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type UmlscPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field DiagramPackage{}.Umlscs []*Umlsc
	DiagramPackage_UmlscsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	DiagramPackage_UmlscsDBID_Index sql.NullInt64
}

// UmlscDB describes a umlsc in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model umlscDB
type UmlscDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field umlscDB.Name
	Name_Data sql.NullString

	// Declation for basic field umlscDB.Activestate
	Activestate_Data sql.NullString

	// Declation for basic field umlscDB.IsInDrawMode
	// provide the sql storage for the boolan
	IsInDrawMode_Data sql.NullBool
	// encoding of pointers
	UmlscPointersEnconding
}

// UmlscDBs arrays umlscDBs
// swagger:response umlscDBsResponse
type UmlscDBs []UmlscDB

// UmlscDBResponse provides response
// swagger:response umlscDBResponse
type UmlscDBResponse struct {
	UmlscDB
}

// UmlscWOP is a Umlsc without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type UmlscWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Activestate string `xlsx:"2"`

	IsInDrawMode bool `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Umlsc_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Activestate",
	"IsInDrawMode",
}

type BackRepoUmlscStruct struct {
	// stores UmlscDB according to their gorm ID
	Map_UmlscDBID_UmlscDB *map[uint]*UmlscDB

	// stores UmlscDB ID according to Umlsc address
	Map_UmlscPtr_UmlscDBID *map[*models.Umlsc]uint

	// stores Umlsc according to their gorm ID
	Map_UmlscDBID_UmlscPtr *map[uint]*models.Umlsc

	db *gorm.DB
}

func (backRepoUmlsc *BackRepoUmlscStruct) GetDB() *gorm.DB {
	return backRepoUmlsc.db
}

// GetUmlscDBFromUmlscPtr is a handy function to access the back repo instance from the stage instance
func (backRepoUmlsc *BackRepoUmlscStruct) GetUmlscDBFromUmlscPtr(umlsc *models.Umlsc) (umlscDB *UmlscDB) {
	id := (*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]
	umlscDB = (*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[id]
	return
}

// BackRepoUmlsc.Init set up the BackRepo of the Umlsc
func (backRepoUmlsc *BackRepoUmlscStruct) Init(db *gorm.DB) (Error error) {

	if backRepoUmlsc.Map_UmlscDBID_UmlscPtr != nil {
		err := errors.New("In Init, backRepoUmlsc.Map_UmlscDBID_UmlscPtr should be nil")
		return err
	}

	if backRepoUmlsc.Map_UmlscDBID_UmlscDB != nil {
		err := errors.New("In Init, backRepoUmlsc.Map_UmlscDBID_UmlscDB should be nil")
		return err
	}

	if backRepoUmlsc.Map_UmlscPtr_UmlscDBID != nil {
		err := errors.New("In Init, backRepoUmlsc.Map_UmlscPtr_UmlscDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Umlsc, 0)
	backRepoUmlsc.Map_UmlscDBID_UmlscPtr = &tmp

	tmpDB := make(map[uint]*UmlscDB, 0)
	backRepoUmlsc.Map_UmlscDBID_UmlscDB = &tmpDB

	tmpID := make(map[*models.Umlsc]uint, 0)
	backRepoUmlsc.Map_UmlscPtr_UmlscDBID = &tmpID

	backRepoUmlsc.db = db
	return
}

// BackRepoUmlsc.CommitPhaseOne commits all staged instances of Umlsc to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for umlsc := range stage.Umlscs {
		backRepoUmlsc.CommitPhaseOneInstance(umlsc)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, umlsc := range *backRepoUmlsc.Map_UmlscDBID_UmlscPtr {
		if _, ok := stage.Umlscs[umlsc]; !ok {
			backRepoUmlsc.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoUmlsc.CommitDeleteInstance commits deletion of Umlsc to the BackRepo
func (backRepoUmlsc *BackRepoUmlscStruct) CommitDeleteInstance(id uint) (Error error) {

	umlsc := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[id]

	// umlsc is not staged anymore, remove umlscDB
	umlscDB := (*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[id]
	query := backRepoUmlsc.db.Unscoped().Delete(&umlscDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoUmlsc.Map_UmlscPtr_UmlscDBID), umlsc)
	delete((*backRepoUmlsc.Map_UmlscDBID_UmlscPtr), id)
	delete((*backRepoUmlsc.Map_UmlscDBID_UmlscDB), id)

	return
}

// BackRepoUmlsc.CommitPhaseOneInstance commits umlsc staged instances of Umlsc to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseOneInstance(umlsc *models.Umlsc) (Error error) {

	// check if the umlsc is not commited yet
	if _, ok := (*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {
		return
	}

	// initiate umlsc
	var umlscDB UmlscDB
	umlscDB.CopyBasicFieldsFromUmlsc(umlsc)

	query := backRepoUmlsc.db.Create(&umlscDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc] = umlscDB.ID
	(*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID] = umlsc
	(*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[umlscDB.ID] = &umlscDB

	return
}

// BackRepoUmlsc.CommitPhaseTwo commits all staged instances of Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, umlsc := range *backRepoUmlsc.Map_UmlscDBID_UmlscPtr {
		backRepoUmlsc.CommitPhaseTwoInstance(backRepo, idx, umlsc)
	}

	return
}

// BackRepoUmlsc.CommitPhaseTwoInstance commits {{structname }} of models.Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, umlsc *models.Umlsc) (Error error) {

	// fetch matching umlscDB
	if umlscDB, ok := (*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[idx]; ok {

		umlscDB.CopyBasicFieldsFromUmlsc(umlsc)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers umlsc.States into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, umlstateAssocEnd := range umlsc.States {

			// get the back repo instance at the association end
			umlstateAssocEnd_DB :=
				backRepo.BackRepoUmlState.GetUmlStateDBFromUmlStatePtr(umlstateAssocEnd)

			// encode reverse pointer in the association end back repo instance
			umlstateAssocEnd_DB.Umlsc_StatesDBID.Int64 = int64(umlscDB.ID)
			umlstateAssocEnd_DB.Umlsc_StatesDBID.Valid = true
			umlstateAssocEnd_DB.Umlsc_StatesDBID_Index.Int64 = int64(idx)
			umlstateAssocEnd_DB.Umlsc_StatesDBID_Index.Valid = true
			if q := backRepoUmlsc.db.Save(umlstateAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoUmlsc.db.Save(&umlscDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Umlsc intance %s", umlsc.Name))
		return err
	}

	return
}

// BackRepoUmlsc.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseOne() (Error error) {

	umlscDBArray := make([]UmlscDB, 0)
	query := backRepoUmlsc.db.Find(&umlscDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	umlscInstancesToBeRemovedFromTheStage := make(map[*models.Umlsc]any)
	for key, value := range models.Stage.Umlscs {
		umlscInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, umlscDB := range umlscDBArray {
		backRepoUmlsc.CheckoutPhaseOneInstance(&umlscDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		umlsc, ok := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID]
		if ok {
			delete(umlscInstancesToBeRemovedFromTheStage, umlsc)
		}
	}

	// remove from stage and back repo's 3 maps all umlscs that are not in the checkout
	for umlsc := range umlscInstancesToBeRemovedFromTheStage {
		umlsc.Unstage()

		// remove instance from the back repo 3 maps
		umlscID := (*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]
		delete((*backRepoUmlsc.Map_UmlscPtr_UmlscDBID), umlsc)
		delete((*backRepoUmlsc.Map_UmlscDBID_UmlscDB), umlscID)
		delete((*backRepoUmlsc.Map_UmlscDBID_UmlscPtr), umlscID)
	}

	return
}

// CheckoutPhaseOneInstance takes a umlscDB that has been found in the DB, updates the backRepo and stages the
// models version of the umlscDB
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseOneInstance(umlscDB *UmlscDB) (Error error) {

	umlsc, ok := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID]
	if !ok {
		umlsc = new(models.Umlsc)

		(*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID] = umlsc
		(*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc] = umlscDB.ID

		// append model store with the new element
		umlsc.Name = umlscDB.Name_Data.String
		umlsc.Stage()
	}
	umlscDB.CopyBasicFieldsToUmlsc(umlsc)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	umlsc.Stage()

	// preserve pointer to umlscDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_UmlscDBID_UmlscDB)[umlscDB hold variable pointers
	umlscDB_Data := *umlscDB
	preservedPtrToUmlsc := &umlscDB_Data
	(*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[umlscDB.ID] = preservedPtrToUmlsc

	return
}

// BackRepoUmlsc.CheckoutPhaseTwo Checkouts all staged instances of Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, umlscDB := range *backRepoUmlsc.Map_UmlscDBID_UmlscDB {
		backRepoUmlsc.CheckoutPhaseTwoInstance(backRepo, umlscDB)
	}
	return
}

// BackRepoUmlsc.CheckoutPhaseTwoInstance Checkouts staged instances of Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, umlscDB *UmlscDB) (Error error) {

	umlsc := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID]
	_ = umlsc // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem umlsc.States in the stage from the encode in the back repo
	// It parses all UmlStateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	umlsc.States = umlsc.States[:0]
	// 2. loop all instances in the type in the association end
	for _, umlstateDB_AssocEnd := range *backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStateDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if umlstateDB_AssocEnd.Umlsc_StatesDBID.Int64 == int64(umlscDB.ID) {
			// 4. fetch the associated instance in the stage
			umlstate_AssocEnd := (*backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStatePtr)[umlstateDB_AssocEnd.ID]
			// 5. append it the association slice
			umlsc.States = append(umlsc.States, umlstate_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(umlsc.States, func(i, j int) bool {
		umlstateDB_i_ID := (*backRepo.BackRepoUmlState.Map_UmlStatePtr_UmlStateDBID)[umlsc.States[i]]
		umlstateDB_j_ID := (*backRepo.BackRepoUmlState.Map_UmlStatePtr_UmlStateDBID)[umlsc.States[j]]

		umlstateDB_i := (*backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStateDB)[umlstateDB_i_ID]
		umlstateDB_j := (*backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStateDB)[umlstateDB_j_ID]

		return umlstateDB_i.Umlsc_StatesDBID_Index.Int64 < umlstateDB_j.Umlsc_StatesDBID_Index.Int64
	})

	return
}

// CommitUmlsc allows commit of a single umlsc (if already staged)
func (backRepo *BackRepoStruct) CommitUmlsc(umlsc *models.Umlsc) {
	backRepo.BackRepoUmlsc.CommitPhaseOneInstance(umlsc)
	if id, ok := (*backRepo.BackRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {
		backRepo.BackRepoUmlsc.CommitPhaseTwoInstance(backRepo, id, umlsc)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitUmlsc allows checkout of a single umlsc (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutUmlsc(umlsc *models.Umlsc) {
	// check if the umlsc is staged
	if _, ok := (*backRepo.BackRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {

		if id, ok := (*backRepo.BackRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {
			var umlscDB UmlscDB
			umlscDB.ID = id

			if err := backRepo.BackRepoUmlsc.db.First(&umlscDB, id).Error; err != nil {
				log.Panicln("CheckoutUmlsc : Problem with getting object with id:", id)
			}
			backRepo.BackRepoUmlsc.CheckoutPhaseOneInstance(&umlscDB)
			backRepo.BackRepoUmlsc.CheckoutPhaseTwoInstance(backRepo, &umlscDB)
		}
	}
}

// CopyBasicFieldsFromUmlsc
func (umlscDB *UmlscDB) CopyBasicFieldsFromUmlsc(umlsc *models.Umlsc) {
	// insertion point for fields commit

	umlscDB.Name_Data.String = umlsc.Name
	umlscDB.Name_Data.Valid = true

	umlscDB.Activestate_Data.String = umlsc.Activestate
	umlscDB.Activestate_Data.Valid = true

	umlscDB.IsInDrawMode_Data.Bool = umlsc.IsInDrawMode
	umlscDB.IsInDrawMode_Data.Valid = true
}

// CopyBasicFieldsFromUmlscWOP
func (umlscDB *UmlscDB) CopyBasicFieldsFromUmlscWOP(umlsc *UmlscWOP) {
	// insertion point for fields commit

	umlscDB.Name_Data.String = umlsc.Name
	umlscDB.Name_Data.Valid = true

	umlscDB.Activestate_Data.String = umlsc.Activestate
	umlscDB.Activestate_Data.Valid = true

	umlscDB.IsInDrawMode_Data.Bool = umlsc.IsInDrawMode
	umlscDB.IsInDrawMode_Data.Valid = true
}

// CopyBasicFieldsToUmlsc
func (umlscDB *UmlscDB) CopyBasicFieldsToUmlsc(umlsc *models.Umlsc) {
	// insertion point for checkout of basic fields (back repo to stage)
	umlsc.Name = umlscDB.Name_Data.String
	umlsc.Activestate = umlscDB.Activestate_Data.String
	umlsc.IsInDrawMode = umlscDB.IsInDrawMode_Data.Bool
}

// CopyBasicFieldsToUmlscWOP
func (umlscDB *UmlscDB) CopyBasicFieldsToUmlscWOP(umlsc *UmlscWOP) {
	umlsc.ID = int(umlscDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	umlsc.Name = umlscDB.Name_Data.String
	umlsc.Activestate = umlscDB.Activestate_Data.String
	umlsc.IsInDrawMode = umlscDB.IsInDrawMode_Data.Bool
}

// Backup generates a json file from a slice of all UmlscDB instances in the backrepo
func (backRepoUmlsc *BackRepoUmlscStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "UmlscDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UmlscDB, 0)
	for _, umlscDB := range *backRepoUmlsc.Map_UmlscDBID_UmlscDB {
		forBackup = append(forBackup, umlscDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Umlsc ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Umlsc file", err.Error())
	}
}

// Backup generates a json file from a slice of all UmlscDB instances in the backrepo
func (backRepoUmlsc *BackRepoUmlscStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UmlscDB, 0)
	for _, umlscDB := range *backRepoUmlsc.Map_UmlscDBID_UmlscDB {
		forBackup = append(forBackup, umlscDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Umlsc")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Umlsc_Fields, -1)
	for _, umlscDB := range forBackup {

		var umlscWOP UmlscWOP
		umlscDB.CopyBasicFieldsToUmlscWOP(&umlscWOP)

		row := sh.AddRow()
		row.WriteStruct(&umlscWOP, -1)
	}
}

// RestoreXL from the "Umlsc" sheet all UmlscDB instances
func (backRepoUmlsc *BackRepoUmlscStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoUmlscid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Umlsc"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoUmlsc.rowVisitorUmlsc)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoUmlsc *BackRepoUmlscStruct) rowVisitorUmlsc(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var umlscWOP UmlscWOP
		row.ReadStruct(&umlscWOP)

		// add the unmarshalled struct to the stage
		umlscDB := new(UmlscDB)
		umlscDB.CopyBasicFieldsFromUmlscWOP(&umlscWOP)

		umlscDB_ID_atBackupTime := umlscDB.ID
		umlscDB.ID = 0
		query := backRepoUmlsc.db.Create(umlscDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[umlscDB.ID] = umlscDB
		BackRepoUmlscid_atBckpTime_newID[umlscDB_ID_atBackupTime] = umlscDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "UmlscDB.json" in dirPath that stores an array
// of UmlscDB and stores it in the database
// the map BackRepoUmlscid_atBckpTime_newID is updated accordingly
func (backRepoUmlsc *BackRepoUmlscStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoUmlscid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "UmlscDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Umlsc file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*UmlscDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_UmlscDBID_UmlscDB
	for _, umlscDB := range forRestore {

		umlscDB_ID_atBackupTime := umlscDB.ID
		umlscDB.ID = 0
		query := backRepoUmlsc.db.Create(umlscDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[umlscDB.ID] = umlscDB
		BackRepoUmlscid_atBckpTime_newID[umlscDB_ID_atBackupTime] = umlscDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Umlsc file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Umlsc>id_atBckpTime_newID
// to compute new index
func (backRepoUmlsc *BackRepoUmlscStruct) RestorePhaseTwo() {

	for _, umlscDB := range *backRepoUmlsc.Map_UmlscDBID_UmlscDB {

		// next line of code is to avert unused variable compilation error
		_ = umlscDB

		// insertion point for reindexing pointers encoding
		// This reindex umlsc.Umlscs
		if umlscDB.DiagramPackage_UmlscsDBID.Int64 != 0 {
			umlscDB.DiagramPackage_UmlscsDBID.Int64 =
				int64(BackRepoDiagramPackageid_atBckpTime_newID[uint(umlscDB.DiagramPackage_UmlscsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoUmlsc.db.Model(umlscDB).Updates(*umlscDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoUmlscid_atBckpTime_newID map[uint]uint
