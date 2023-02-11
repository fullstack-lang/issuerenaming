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
import { WaldoDB } from '../waldo-db'
import { WaldoService } from '../waldo.service'

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
  selector: 'app-waldostable',
  templateUrl: './waldos-table.component.html',
  styleUrls: ['./waldos-table.component.css'],
})
export class WaldosTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Waldo instances
  selection: SelectionModel<WaldoDB> = new (SelectionModel)
  initialSelection = new Array<WaldoDB>()

  // the data source for the table
  waldos: WaldoDB[] = []
  matTableDataSource: MatTableDataSource<WaldoDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.waldos
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
    this.matTableDataSource.sortingDataAccessor = (waldoDB: WaldoDB, property: string) => {
      switch (property) {
        case 'ID':
          return waldoDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return waldoDB.Name;

        case 'Bar_Waldos':
          if (this.frontRepo.Bars.get(waldoDB.Bar_WaldosDBID.Int64) != undefined) {
            return this.frontRepo.Bars.get(waldoDB.Bar_WaldosDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (waldoDB: WaldoDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the waldoDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += waldoDB.Name.toLowerCase()
      if (waldoDB.Bar_WaldosDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Bars.get(waldoDB.Bar_WaldosDBID.Int64)!.Name.toLowerCase()
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
    private waldoService: WaldoService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of waldo instances
    public dialogRef: MatDialogRef<WaldosTableComponent>,
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
    this.waldoService.WaldoServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getWaldos()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Bar_Waldos",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Bar_Waldos",
      ]
      this.selection = new SelectionModel<WaldoDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getWaldos()
    this.matTableDataSource = new MatTableDataSource(this.waldos)
  }

  getWaldos(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.waldos = this.frontRepo.Waldos_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let waldo of this.waldos) {
            let ID = this.dialogData.ID
            let revPointer = waldo[this.dialogData.ReversePointer as keyof WaldoDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(waldo)
            }
            this.selection = new SelectionModel<WaldoDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, WaldoDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as WaldoDB[]
          for (let associationInstance of sourceField) {
            let waldo = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as WaldoDB
            this.initialSelection.push(waldo)
          }

          this.selection = new SelectionModel<WaldoDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.waldos
      }
    )
  }

  // newWaldo initiate a new waldo
  // create a new Waldo objet
  newWaldo() {
  }

  deleteWaldo(waldoID: number, waldo: WaldoDB) {
    // list of waldos is truncated of waldo before the delete
    this.waldos = this.waldos.filter(h => h !== waldo);

    this.waldoService.deleteWaldo(waldoID).subscribe(
      waldo => {
        this.waldoService.WaldoServiceChanged.next("delete")
      }
    );
  }

  editWaldo(waldoID: number, waldo: WaldoDB) {

  }

  // display waldo in router
  displayWaldoInRouter(waldoID: number) {
    this.router.navigate(["github_com_fullstack_lang_issuerenaming_go-" + "waldo-display", waldoID])
  }

  // set editor outlet
  setEditorRouterOutlet(waldoID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_issuerenaming_go_editor: ["github_com_fullstack_lang_issuerenaming_go-" + "waldo-detail", waldoID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.waldos.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.waldos.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<WaldoDB>()

      // reset all initial selection of waldo that belong to waldo
      for (let waldo of this.initialSelection) {
        let index = waldo[this.dialogData.ReversePointer as keyof WaldoDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(waldo)

      }

      // from selection, set waldo that belong to waldo
      for (let waldo of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = waldo[this.dialogData.ReversePointer as keyof WaldoDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(waldo)
      }


      // update all waldo (only update selection & initial selection)
      for (let waldo of toUpdate) {
        this.waldoService.updateWaldo(waldo)
          .subscribe(waldo => {
            this.waldoService.WaldoServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, WaldoDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedWaldo = new Set<number>()
      for (let waldo of this.initialSelection) {
        if (this.selection.selected.includes(waldo)) {
          // console.log("waldo " + waldo.Name + " is still selected")
        } else {
          console.log("waldo " + waldo.Name + " has been unselected")
          unselectedWaldo.add(waldo.ID)
          console.log("is unselected " + unselectedWaldo.has(waldo.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let waldo = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as WaldoDB
      if (unselectedWaldo.has(waldo.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<WaldoDB>) = new Array<WaldoDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          waldo => {
            if (!this.initialSelection.includes(waldo)) {
              // console.log("waldo " + waldo.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + waldo.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = waldo.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = waldo.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("waldo " + waldo.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<WaldoDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
