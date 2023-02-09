package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/issuerenaming/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Foo": &(ref_models.Foo{}),

	"ref_models.Foo.Name": (ref_models.Foo{}).Name,

	"ref_models.Foo.Waldos": (ref_models.Foo{}).Waldos,

	"ref_models.GONG__ENUM_CAST_INT": ref_models.GONG__ENUM_CAST_INT,

	"ref_models.GONG__ENUM_CAST_STRING": ref_models.GONG__ENUM_CAST_STRING,

	"ref_models.GONG__ExpressionType": ref_models.GONG__ExpressionType(""),

	"ref_models.GONG__FIELD_OR_CONST_VALUE": ref_models.GONG__FIELD_OR_CONST_VALUE,

	"ref_models.GONG__FIELD_VALUE": ref_models.GONG__FIELD_VALUE,

	"ref_models.GONG__IDENTIFIER_CONST": ref_models.GONG__IDENTIFIER_CONST,

	"ref_models.GONG__STRUCT_INSTANCE": ref_models.GONG__STRUCT_INSTANCE,

	"ref_models.Note": ref_models.Note,

	"ref_models.Waldo": &(ref_models.Waldo{}),

	"ref_models.Waldo.Name": (ref_models.Waldo{}).Name,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Foo := (&models.GongStructShape{Name: `NewDiagram-Foo`}).Stage()
	__GongStructShape__000001_NewDiagram_Waldo := (&models.GongStructShape{Name: `NewDiagram-Waldo`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Waldos := (&models.Link{Name: `Waldos`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_Note := (&models.NoteShape{Name: `Note`}).Stage()

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Foo := (&models.NoteShapeLink{Name: `Foo`}).Stage()
	__NoteShapeLink__000001_Foo_Waldos := (&models.NoteShapeLink{Name: `Foo.Waldos`}).Stage()
	__NoteShapeLink__000002_Waldo := (&models.NoteShapeLink{Name: `Waldo`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Foo := (&models.Position{Name: `Pos-NewDiagram-Foo`}).Stage()
	__Position__000001_Pos_NewDiagram_Waldo := (&models.Position{Name: `Pos-NewDiagram-Waldo`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Foo_and_NewDiagram_Waldo := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Foo and NewDiagram-Waldo`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Foo.Name = `NewDiagram-Foo`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Foo]
	__GongStructShape__000000_NewDiagram_Foo.Identifier = `ref_models.Foo`
	__GongStructShape__000000_NewDiagram_Foo.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Foo.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Foo.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Foo.Heigth = 63.000000
	__GongStructShape__000000_NewDiagram_Foo.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Waldo.Name = `NewDiagram-Waldo`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Waldo]
	__GongStructShape__000001_NewDiagram_Waldo.Identifier = `ref_models.Waldo`
	__GongStructShape__000001_NewDiagram_Waldo.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Waldo.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Waldo.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Waldo.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_Waldo.IsSelected = false

	// Link values setup
	__Link__000000_Waldos.Name = `Waldos`
	__Link__000000_Waldos.Structname = `Foo`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Foo.Waldos]
	__Link__000000_Waldos.Identifier = `ref_models.Foo.Waldos`
	__Link__000000_Waldos.Fieldtypename = `Waldo`
	__Link__000000_Waldos.TargetMultiplicity = models.MANY
	__Link__000000_Waldos.SourceMultiplicity = models.ZERO_ONE

	// NoteShape values setup
	__NoteShape__000000_Note.Name = `Note`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Note]
	__NoteShape__000000_Note.Identifier = `ref_models.Note`
	__NoteShape__000000_Note.Body = `[models.Foo] is
related to [models.Waldo] throughs the field
[models.Foo.Waldos]
`
	__NoteShape__000000_Note.X = 300.000000
	__NoteShape__000000_Note.Y = 60.000000
	__NoteShape__000000_Note.Width = 240.000000
	__NoteShape__000000_Note.Heigth = 63.000000
	__NoteShape__000000_Note.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Foo.Name = `Foo`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Foo]
	__NoteShapeLink__000000_Foo.Identifier = `ref_models.Foo`
	__NoteShapeLink__000000_Foo.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000001_Foo_Waldos.Name = `Foo.Waldos`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Foo.Waldos]
	__NoteShapeLink__000001_Foo_Waldos.Identifier = `ref_models.Foo.Waldos`
	__NoteShapeLink__000001_Foo_Waldos.Type = models.NOTE_SHAPE_LINK_TO_GONG_FIELD

	// NoteShapeLink values setup
	__NoteShapeLink__000002_Waldo.Name = `Waldo`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Waldo]
	__NoteShapeLink__000002_Waldo.Identifier = `ref_models.Waldo`
	__NoteShapeLink__000002_Waldo.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// Position values setup
	__Position__000000_Pos_NewDiagram_Foo.X = 60.000000
	__Position__000000_Pos_NewDiagram_Foo.Y = 170.000000
	__Position__000000_Pos_NewDiagram_Foo.Name = `Pos-NewDiagram-Foo`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Waldo.X = 600.000000
	__Position__000001_Pos_NewDiagram_Waldo.Y = 180.000000
	__Position__000001_Pos_NewDiagram_Waldo.Name = `Pos-NewDiagram-Waldo`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Foo_and_NewDiagram_Waldo.X = 465.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Foo_and_NewDiagram_Waldo.Y = 288.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Foo_and_NewDiagram_Waldo.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Foo and NewDiagram-Waldo`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Foo)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Waldo)
	__Classdiagram__000000_NewDiagram.NoteShapes = append(__Classdiagram__000000_NewDiagram.NoteShapes, __NoteShape__000000_Note)
	__GongStructShape__000000_NewDiagram_Foo.Position = __Position__000000_Pos_NewDiagram_Foo
	__GongStructShape__000000_NewDiagram_Foo.Links = append(__GongStructShape__000000_NewDiagram_Foo.Links, __Link__000000_Waldos)
	__GongStructShape__000001_NewDiagram_Waldo.Position = __Position__000001_Pos_NewDiagram_Waldo
	__Link__000000_Waldos.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Foo_and_NewDiagram_Waldo
	__NoteShape__000000_Note.NoteShapeLinks = append(__NoteShape__000000_Note.NoteShapeLinks, __NoteShapeLink__000000_Foo)
	__NoteShape__000000_Note.NoteShapeLinks = append(__NoteShape__000000_Note.NoteShapeLinks, __NoteShapeLink__000001_Foo_Waldos)
	__NoteShape__000000_Note.NoteShapeLinks = append(__NoteShape__000000_Note.NoteShapeLinks, __NoteShapeLink__000002_Waldo)
}


