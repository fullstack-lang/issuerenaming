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
var dummy_Bar_sql sql.NullBool
var dummy_Bar_time time.Duration
var dummy_Bar_sort sort.Float64Slice

// BarAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model barAPI
type BarAPI struct {
	gorm.Model

	models.Bar

	// encoding of pointers
	BarPointersEnconding
}

// BarPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type BarPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// BarDB describes a bar in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model barDB
type BarDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field barDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	BarPointersEnconding
}

// BarDBs arrays barDBs
// swagger:response barDBsResponse
type BarDBs []BarDB

// BarDBResponse provides response
// swagger:response barDBResponse
type BarDBResponse struct {
	BarDB
}

// BarWOP is a Bar without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type BarWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Bar_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoBarStruct struct {
	// stores BarDB according to their gorm ID
	Map_BarDBID_BarDB *map[uint]*BarDB

	// stores BarDB ID according to Bar address
	Map_BarPtr_BarDBID *map[*models.Bar]uint

	// stores Bar according to their gorm ID
	Map_BarDBID_BarPtr *map[uint]*models.Bar

	db *gorm.DB
}

func (backRepoBar *BackRepoBarStruct) GetDB() *gorm.DB {
	return backRepoBar.db
}

// GetBarDBFromBarPtr is a handy function to access the back repo instance from the stage instance
func (backRepoBar *BackRepoBarStruct) GetBarDBFromBarPtr(bar *models.Bar) (barDB *BarDB) {
	id := (*backRepoBar.Map_BarPtr_BarDBID)[bar]
	barDB = (*backRepoBar.Map_BarDBID_BarDB)[id]
	return
}

// BackRepoBar.Init set up the BackRepo of the Bar
func (backRepoBar *BackRepoBarStruct) Init(db *gorm.DB) (Error error) {

	if backRepoBar.Map_BarDBID_BarPtr != nil {
		err := errors.New("In Init, backRepoBar.Map_BarDBID_BarPtr should be nil")
		return err
	}

	if backRepoBar.Map_BarDBID_BarDB != nil {
		err := errors.New("In Init, backRepoBar.Map_BarDBID_BarDB should be nil")
		return err
	}

	if backRepoBar.Map_BarPtr_BarDBID != nil {
		err := errors.New("In Init, backRepoBar.Map_BarPtr_BarDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Bar, 0)
	backRepoBar.Map_BarDBID_BarPtr = &tmp

	tmpDB := make(map[uint]*BarDB, 0)
	backRepoBar.Map_BarDBID_BarDB = &tmpDB

	tmpID := make(map[*models.Bar]uint, 0)
	backRepoBar.Map_BarPtr_BarDBID = &tmpID

	backRepoBar.db = db
	return
}

// BackRepoBar.CommitPhaseOne commits all staged instances of Bar to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoBar *BackRepoBarStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for bar := range stage.Bars {
		backRepoBar.CommitPhaseOneInstance(bar)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, bar := range *backRepoBar.Map_BarDBID_BarPtr {
		if _, ok := stage.Bars[bar]; !ok {
			backRepoBar.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoBar.CommitDeleteInstance commits deletion of Bar to the BackRepo
func (backRepoBar *BackRepoBarStruct) CommitDeleteInstance(id uint) (Error error) {

	bar := (*backRepoBar.Map_BarDBID_BarPtr)[id]

	// bar is not staged anymore, remove barDB
	barDB := (*backRepoBar.Map_BarDBID_BarDB)[id]
	query := backRepoBar.db.Unscoped().Delete(&barDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoBar.Map_BarPtr_BarDBID), bar)
	delete((*backRepoBar.Map_BarDBID_BarPtr), id)
	delete((*backRepoBar.Map_BarDBID_BarDB), id)

	return
}

// BackRepoBar.CommitPhaseOneInstance commits bar staged instances of Bar to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoBar *BackRepoBarStruct) CommitPhaseOneInstance(bar *models.Bar) (Error error) {

	// check if the bar is not commited yet
	if _, ok := (*backRepoBar.Map_BarPtr_BarDBID)[bar]; ok {
		return
	}

	// initiate bar
	var barDB BarDB
	barDB.CopyBasicFieldsFromBar(bar)

	query := backRepoBar.db.Create(&barDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoBar.Map_BarPtr_BarDBID)[bar] = barDB.ID
	(*backRepoBar.Map_BarDBID_BarPtr)[barDB.ID] = bar
	(*backRepoBar.Map_BarDBID_BarDB)[barDB.ID] = &barDB

	return
}

// BackRepoBar.CommitPhaseTwo commits all staged instances of Bar to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBar *BackRepoBarStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, bar := range *backRepoBar.Map_BarDBID_BarPtr {
		backRepoBar.CommitPhaseTwoInstance(backRepo, idx, bar)
	}

	return
}

// BackRepoBar.CommitPhaseTwoInstance commits {{structname }} of models.Bar to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBar *BackRepoBarStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, bar *models.Bar) (Error error) {

	// fetch matching barDB
	if barDB, ok := (*backRepoBar.Map_BarDBID_BarDB)[idx]; ok {

		barDB.CopyBasicFieldsFromBar(bar)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers bar.Waldos into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, waldoAssocEnd := range bar.Waldos {

			// get the back repo instance at the association end
			waldoAssocEnd_DB :=
				backRepo.BackRepoWaldo.GetWaldoDBFromWaldoPtr(waldoAssocEnd)

			// encode reverse pointer in the association end back repo instance
			waldoAssocEnd_DB.Bar_WaldosDBID.Int64 = int64(barDB.ID)
			waldoAssocEnd_DB.Bar_WaldosDBID.Valid = true
			waldoAssocEnd_DB.Bar_WaldosDBID_Index.Int64 = int64(idx)
			waldoAssocEnd_DB.Bar_WaldosDBID_Index.Valid = true
			if q := backRepoBar.db.Save(waldoAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoBar.db.Save(&barDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Bar intance %s", bar.Name))
		return err
	}

	return
}

// BackRepoBar.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoBar *BackRepoBarStruct) CheckoutPhaseOne() (Error error) {

	barDBArray := make([]BarDB, 0)
	query := backRepoBar.db.Find(&barDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	barInstancesToBeRemovedFromTheStage := make(map[*models.Bar]any)
	for key, value := range models.Stage.Bars {
		barInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, barDB := range barDBArray {
		backRepoBar.CheckoutPhaseOneInstance(&barDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		bar, ok := (*backRepoBar.Map_BarDBID_BarPtr)[barDB.ID]
		if ok {
			delete(barInstancesToBeRemovedFromTheStage, bar)
		}
	}

	// remove from stage and back repo's 3 maps all bars that are not in the checkout
	for bar := range barInstancesToBeRemovedFromTheStage {
		bar.Unstage()

		// remove instance from the back repo 3 maps
		barID := (*backRepoBar.Map_BarPtr_BarDBID)[bar]
		delete((*backRepoBar.Map_BarPtr_BarDBID), bar)
		delete((*backRepoBar.Map_BarDBID_BarDB), barID)
		delete((*backRepoBar.Map_BarDBID_BarPtr), barID)
	}

	return
}

// CheckoutPhaseOneInstance takes a barDB that has been found in the DB, updates the backRepo and stages the
// models version of the barDB
func (backRepoBar *BackRepoBarStruct) CheckoutPhaseOneInstance(barDB *BarDB) (Error error) {

	bar, ok := (*backRepoBar.Map_BarDBID_BarPtr)[barDB.ID]
	if !ok {
		bar = new(models.Bar)

		(*backRepoBar.Map_BarDBID_BarPtr)[barDB.ID] = bar
		(*backRepoBar.Map_BarPtr_BarDBID)[bar] = barDB.ID

		// append model store with the new element
		bar.Name = barDB.Name_Data.String
		bar.Stage()
	}
	barDB.CopyBasicFieldsToBar(bar)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	bar.Stage()

	// preserve pointer to barDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_BarDBID_BarDB)[barDB hold variable pointers
	barDB_Data := *barDB
	preservedPtrToBar := &barDB_Data
	(*backRepoBar.Map_BarDBID_BarDB)[barDB.ID] = preservedPtrToBar

	return
}

// BackRepoBar.CheckoutPhaseTwo Checkouts all staged instances of Bar to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBar *BackRepoBarStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, barDB := range *backRepoBar.Map_BarDBID_BarDB {
		backRepoBar.CheckoutPhaseTwoInstance(backRepo, barDB)
	}
	return
}

// BackRepoBar.CheckoutPhaseTwoInstance Checkouts staged instances of Bar to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBar *BackRepoBarStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, barDB *BarDB) (Error error) {

	bar := (*backRepoBar.Map_BarDBID_BarPtr)[barDB.ID]
	_ = bar // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem bar.Waldos in the stage from the encode in the back repo
	// It parses all WaldoDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	bar.Waldos = bar.Waldos[:0]
	// 2. loop all instances in the type in the association end
	for _, waldoDB_AssocEnd := range *backRepo.BackRepoWaldo.Map_WaldoDBID_WaldoDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if waldoDB_AssocEnd.Bar_WaldosDBID.Int64 == int64(barDB.ID) {
			// 4. fetch the associated instance in the stage
			waldo_AssocEnd := (*backRepo.BackRepoWaldo.Map_WaldoDBID_WaldoPtr)[waldoDB_AssocEnd.ID]
			// 5. append it the association slice
			bar.Waldos = append(bar.Waldos, waldo_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(bar.Waldos, func(i, j int) bool {
		waldoDB_i_ID := (*backRepo.BackRepoWaldo.Map_WaldoPtr_WaldoDBID)[bar.Waldos[i]]
		waldoDB_j_ID := (*backRepo.BackRepoWaldo.Map_WaldoPtr_WaldoDBID)[bar.Waldos[j]]

		waldoDB_i := (*backRepo.BackRepoWaldo.Map_WaldoDBID_WaldoDB)[waldoDB_i_ID]
		waldoDB_j := (*backRepo.BackRepoWaldo.Map_WaldoDBID_WaldoDB)[waldoDB_j_ID]

		return waldoDB_i.Bar_WaldosDBID_Index.Int64 < waldoDB_j.Bar_WaldosDBID_Index.Int64
	})

	return
}

// CommitBar allows commit of a single bar (if already staged)
func (backRepo *BackRepoStruct) CommitBar(bar *models.Bar) {
	backRepo.BackRepoBar.CommitPhaseOneInstance(bar)
	if id, ok := (*backRepo.BackRepoBar.Map_BarPtr_BarDBID)[bar]; ok {
		backRepo.BackRepoBar.CommitPhaseTwoInstance(backRepo, id, bar)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitBar allows checkout of a single bar (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutBar(bar *models.Bar) {
	// check if the bar is staged
	if _, ok := (*backRepo.BackRepoBar.Map_BarPtr_BarDBID)[bar]; ok {

		if id, ok := (*backRepo.BackRepoBar.Map_BarPtr_BarDBID)[bar]; ok {
			var barDB BarDB
			barDB.ID = id

			if err := backRepo.BackRepoBar.db.First(&barDB, id).Error; err != nil {
				log.Panicln("CheckoutBar : Problem with getting object with id:", id)
			}
			backRepo.BackRepoBar.CheckoutPhaseOneInstance(&barDB)
			backRepo.BackRepoBar.CheckoutPhaseTwoInstance(backRepo, &barDB)
		}
	}
}

// CopyBasicFieldsFromBar
func (barDB *BarDB) CopyBasicFieldsFromBar(bar *models.Bar) {
	// insertion point for fields commit

	barDB.Name_Data.String = bar.Name
	barDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromBarWOP
func (barDB *BarDB) CopyBasicFieldsFromBarWOP(bar *BarWOP) {
	// insertion point for fields commit

	barDB.Name_Data.String = bar.Name
	barDB.Name_Data.Valid = true
}

// CopyBasicFieldsToBar
func (barDB *BarDB) CopyBasicFieldsToBar(bar *models.Bar) {
	// insertion point for checkout of basic fields (back repo to stage)
	bar.Name = barDB.Name_Data.String
}

// CopyBasicFieldsToBarWOP
func (barDB *BarDB) CopyBasicFieldsToBarWOP(bar *BarWOP) {
	bar.ID = int(barDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	bar.Name = barDB.Name_Data.String
}

// Backup generates a json file from a slice of all BarDB instances in the backrepo
func (backRepoBar *BackRepoBarStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "BarDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*BarDB, 0)
	for _, barDB := range *backRepoBar.Map_BarDBID_BarDB {
		forBackup = append(forBackup, barDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Bar ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Bar file", err.Error())
	}
}

// Backup generates a json file from a slice of all BarDB instances in the backrepo
func (backRepoBar *BackRepoBarStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*BarDB, 0)
	for _, barDB := range *backRepoBar.Map_BarDBID_BarDB {
		forBackup = append(forBackup, barDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Bar")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Bar_Fields, -1)
	for _, barDB := range forBackup {

		var barWOP BarWOP
		barDB.CopyBasicFieldsToBarWOP(&barWOP)

		row := sh.AddRow()
		row.WriteStruct(&barWOP, -1)
	}
}

// RestoreXL from the "Bar" sheet all BarDB instances
func (backRepoBar *BackRepoBarStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoBarid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Bar"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoBar.rowVisitorBar)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoBar *BackRepoBarStruct) rowVisitorBar(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var barWOP BarWOP
		row.ReadStruct(&barWOP)

		// add the unmarshalled struct to the stage
		barDB := new(BarDB)
		barDB.CopyBasicFieldsFromBarWOP(&barWOP)

		barDB_ID_atBackupTime := barDB.ID
		barDB.ID = 0
		query := backRepoBar.db.Create(barDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoBar.Map_BarDBID_BarDB)[barDB.ID] = barDB
		BackRepoBarid_atBckpTime_newID[barDB_ID_atBackupTime] = barDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "BarDB.json" in dirPath that stores an array
// of BarDB and stores it in the database
// the map BackRepoBarid_atBckpTime_newID is updated accordingly
func (backRepoBar *BackRepoBarStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoBarid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "BarDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Bar file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*BarDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_BarDBID_BarDB
	for _, barDB := range forRestore {

		barDB_ID_atBackupTime := barDB.ID
		barDB.ID = 0
		query := backRepoBar.db.Create(barDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoBar.Map_BarDBID_BarDB)[barDB.ID] = barDB
		BackRepoBarid_atBckpTime_newID[barDB_ID_atBackupTime] = barDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Bar file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Bar>id_atBckpTime_newID
// to compute new index
func (backRepoBar *BackRepoBarStruct) RestorePhaseTwo() {

	for _, barDB := range *backRepoBar.Map_BarDBID_BarDB {

		// next line of code is to avert unused variable compilation error
		_ = barDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoBar.db.Model(barDB).Updates(*barDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoBarid_atBckpTime_newID map[uint]uint
