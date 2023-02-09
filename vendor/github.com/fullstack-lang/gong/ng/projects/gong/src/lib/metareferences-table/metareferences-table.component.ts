// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { MetaReferenceDB } from '../metareference-db'
import { MetaReferenceService } from '../metareference.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-metareferencestable',
  templateUrl: './metareferences-table.component.html',
  styleUrls: ['./metareferences-table.component.css'],
})
export class MetaReferencesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of MetaReference instances
  selection: SelectionModel<MetaReferenceDB> = new (SelectionModel)
  initialSelection = new Array<MetaReferenceDB>()

  // the data source for the table
  metareferences: MetaReferenceDB[] = []
  matTableDataSource: MatTableDataSource<MetaReferenceDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.metareferences
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (metareferenceDB: MetaReferenceDB, property: string) => {
      switch (property) {
        case 'ID':
          return metareferenceDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return metareferenceDB.Name;

        case 'Meta_MetaReferences':
          if (this.frontRepo.Metas.get(metareferenceDB.Meta_MetaReferencesDBID.Int64) != undefined) {
            return this.frontRepo.Metas.get(metareferenceDB.Meta_MetaReferencesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (metareferenceDB: MetaReferenceDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the metareferenceDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += metareferenceDB.Name.toLowerCase()
      if (metareferenceDB.Meta_MetaReferencesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Metas.get(metareferenceDB.Meta_MetaReferencesDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private metareferenceService: MetaReferenceService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of metareference instances
    public dialogRef: MatDialogRef<MetaReferencesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.metareferenceService.MetaReferenceServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getMetaReferences()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Meta_MetaReferences",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Meta_MetaReferences",
      ]
      this.selection = new SelectionModel<MetaReferenceDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getMetaReferences()
    this.matTableDataSource = new MatTableDataSource(this.metareferences)
  }

  getMetaReferences(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.metareferences = this.frontRepo.MetaReferences_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let metareference of this.metareferences) {
            let ID = this.dialogData.ID
            let revPointer = metareference[this.dialogData.ReversePointer as keyof MetaReferenceDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(metareference)
            }
            this.selection = new SelectionModel<MetaReferenceDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, MetaReferenceDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as MetaReferenceDB[]
          for (let associationInstance of sourceField) {
            let metareference = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as MetaReferenceDB
            this.initialSelection.push(metareference)
          }

          this.selection = new SelectionModel<MetaReferenceDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.metareferences
      }
    )
  }

  // newMetaReference initiate a new metareference
  // create a new MetaReference objet
  newMetaReference() {
  }

  deleteMetaReference(metareferenceID: number, metareference: MetaReferenceDB) {
    // list of metareferences is truncated of metareference before the delete
    this.metareferences = this.metareferences.filter(h => h !== metareference);

    this.metareferenceService.deleteMetaReference(metareferenceID).subscribe(
      metareference => {
        this.metareferenceService.MetaReferenceServiceChanged.next("delete")
      }
    );
  }

  editMetaReference(metareferenceID: number, metareference: MetaReferenceDB) {

  }

  // display metareference in router
  displayMetaReferenceInRouter(metareferenceID: number) {
    this.router.navigate(["github_com_fullstack_lang_gong_go-" + "metareference-display", metareferenceID])
  }

  // set editor outlet
  setEditorRouterOutlet(metareferenceID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "metareference-detail", metareferenceID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.metareferences.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.metareferences.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<MetaReferenceDB>()

      // reset all initial selection of metareference that belong to metareference
      for (let metareference of this.initialSelection) {
        let index = metareference[this.dialogData.ReversePointer as keyof MetaReferenceDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(metareference)

      }

      // from selection, set metareference that belong to metareference
      for (let metareference of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = metareference[this.dialogData.ReversePointer as keyof MetaReferenceDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(metareference)
      }


      // update all metareference (only update selection & initial selection)
      for (let metareference of toUpdate) {
        this.metareferenceService.updateMetaReference(metareference)
          .subscribe(metareference => {
            this.metareferenceService.MetaReferenceServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, MetaReferenceDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedMetaReference = new Set<number>()
      for (let metareference of this.initialSelection) {
        if (this.selection.selected.includes(metareference)) {
          // console.log("metareference " + metareference.Name + " is still selected")
        } else {
          console.log("metareference " + metareference.Name + " has been unselected")
          unselectedMetaReference.add(metareference.ID)
          console.log("is unselected " + unselectedMetaReference.has(metareference.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let metareference = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as MetaReferenceDB
      if (unselectedMetaReference.has(metareference.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<MetaReferenceDB>) = new Array<MetaReferenceDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          metareference => {
            if (!this.initialSelection.includes(metareference)) {
              // console.log("metareference " + metareference.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + metareference.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = metareference.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = metareference.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("metareference " + metareference.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<MetaReferenceDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
