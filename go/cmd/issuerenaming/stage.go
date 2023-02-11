package main

import (
	"time"

	"github.com/fullstack-lang/issuerenaming/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage models.StageStruct
var ___dummy__Time_stage time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = stageInjection
// }

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Foo
	__Foo__000000_A := (&models.Bar{Name: `A`}).Stage()

	// Declarations of staged instances of Waldo
	__Waldo__000000_B := (&models.Waldo{Name: `B`}).Stage()

	// Setup of values

	// Foo values setup
	__Foo__000000_A.Name = `A`

	// Waldo values setup
	__Waldo__000000_B.Name = `B`

	// Setup of pointers
	__Foo__000000_A.Waldos = append(__Foo__000000_A.Waldos, __Waldo__000000_B)
}
