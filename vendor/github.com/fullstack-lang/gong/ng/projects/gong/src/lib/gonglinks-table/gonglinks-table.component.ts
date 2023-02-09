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
import { GongLinkDB } from '../gonglink-db'
import { GongLinkService } from '../gonglink.service'

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
  selector: 'app-gonglinkstable',
  templateUrl: './gonglinks-table.component.html',
  styleUrls: ['./gonglinks-table.component.css'],
})
export class GongLinksTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of GongLink instances
  selection: SelectionModel<GongLinkDB> = new (SelectionModel)
  initialSelection = new Array<GongLinkDB>()

  // the data source for the table
  gonglinks: GongLinkDB[] = []
  matTableDataSource: MatTableDataSource<GongLinkDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.gonglinks
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
    this.matTableDataSource.sortingDataAccessor = (gonglinkDB: GongLinkDB, property: string) => {
      switch (property) {
        case 'ID':
          return gonglinkDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return gonglinkDB.Name;

        case 'Recv':
          return gonglinkDB.Recv;

        case 'ImportPath':
          return gonglinkDB.ImportPath;

        case 'GongNote_Links':
          if (this.frontRepo.GongNotes.get(gonglinkDB.GongNote_LinksDBID.Int64) != undefined) {
            return this.frontRepo.GongNotes.get(gonglinkDB.GongNote_LinksDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (gonglinkDB: GongLinkDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the gonglinkDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += gonglinkDB.Name.toLowerCase()
      mergedContent += gonglinkDB.Recv.toLowerCase()
      mergedContent += gonglinkDB.ImportPath.toLowerCase()
      if (gonglinkDB.GongNote_LinksDBID.Int64 != 0) {
        mergedContent += this.frontRepo.GongNotes.get(gonglinkDB.GongNote_LinksDBID.Int64)!.Name.toLowerCase()
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
    private gonglinkService: GongLinkService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gonglink instances
    public dialogRef: MatDialogRef<GongLinksTableComponent>,
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
    this.gonglinkService.GongLinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongLinks()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Recv",
        "ImportPath",
        "GongNote_Links",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Recv",
        "ImportPath",
        "GongNote_Links",
      ]
      this.selection = new SelectionModel<GongLinkDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongLinks()
    this.matTableDataSource = new MatTableDataSource(this.gonglinks)
  }

  getGongLinks(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.gonglinks = this.frontRepo.GongLinks_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let gonglink of this.gonglinks) {
            let ID = this.dialogData.ID
            let revPointer = gonglink[this.dialogData.ReversePointer as keyof GongLinkDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(gonglink)
            }
            this.selection = new SelectionModel<GongLinkDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongLinkDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as GongLinkDB[]
          for (let associationInstance of sourceField) {
            let gonglink = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongLinkDB
            this.initialSelection.push(gonglink)
          }

          this.selection = new SelectionModel<GongLinkDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.gonglinks
      }
    )
  }

  // newGongLink initiate a new gonglink
  // create a new GongLink objet
  newGongLink() {
  }

  deleteGongLink(gonglinkID: number, gonglink: GongLinkDB) {
    // list of gonglinks is truncated of gonglink before the delete
    this.gonglinks = this.gonglinks.filter(h => h !== gonglink);

    this.gonglinkService.deleteGongLink(gonglinkID).subscribe(
      gonglink => {
        this.gonglinkService.GongLinkServiceChanged.next("delete")
      }
    );
  }

  editGongLink(gonglinkID: number, gonglink: GongLinkDB) {

  }

  // display gonglink in router
  displayGongLinkInRouter(gonglinkID: number) {
    this.router.navigate(["github_com_fullstack_lang_gong_go-" + "gonglink-display", gonglinkID])
  }

  // set editor outlet
  setEditorRouterOutlet(gonglinkID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gonglink-detail", gonglinkID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gonglinks.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gonglinks.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<GongLinkDB>()

      // reset all initial selection of gonglink that belong to gonglink
      for (let gonglink of this.initialSelection) {
        let index = gonglink[this.dialogData.ReversePointer as keyof GongLinkDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(gonglink)

      }

      // from selection, set gonglink that belong to gonglink
      for (let gonglink of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = gonglink[this.dialogData.ReversePointer as keyof GongLinkDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(gonglink)
      }


      // update all gonglink (only update selection & initial selection)
      for (let gonglink of toUpdate) {
        this.gonglinkService.updateGongLink(gonglink)
          .subscribe(gonglink => {
            this.gonglinkService.GongLinkServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongLinkDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedGongLink = new Set<number>()
      for (let gonglink of this.initialSelection) {
        if (this.selection.selected.includes(gonglink)) {
          // console.log("gonglink " + gonglink.Name + " is still selected")
        } else {
          console.log("gonglink " + gonglink.Name + " has been unselected")
          unselectedGongLink.add(gonglink.ID)
          console.log("is unselected " + unselectedGongLink.has(gonglink.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let gonglink = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongLinkDB
      if (unselectedGongLink.has(gonglink.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<GongLinkDB>) = new Array<GongLinkDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          gonglink => {
            if (!this.initialSelection.includes(gonglink)) {
              // console.log("gonglink " + gonglink.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + gonglink.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = gonglink.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = gonglink.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("gonglink " + gonglink.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<GongLinkDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
