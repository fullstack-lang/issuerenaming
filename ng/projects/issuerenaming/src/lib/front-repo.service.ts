import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { BarDB } from './bar-db'
import { BarService } from './bar.service'

import { WaldoDB } from './waldo-db'
import { WaldoService } from './waldo.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Bars_array = new Array<BarDB>(); // array of repo instances
  Bars = new Map<number, BarDB>(); // map of repo instances
  Bars_batch = new Map<number, BarDB>(); // same but only in last GET (for finding repo instances to delete)
  Waldos_array = new Array<WaldoDB>(); // array of repo instances
  Waldos = new Map<number, WaldoDB>(); // map of repo instances
  Waldos_batch = new Map<number, WaldoDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private barService: BarService,
    private waldoService: WaldoService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<BarDB[]>,
    Observable<WaldoDB[]>,
  ] = [ // insertion point sub template 
      this.barService.getBars(),
      this.waldoService.getWaldos(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            bars_,
            waldos_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var bars: BarDB[]
            bars = bars_ as BarDB[]
            var waldos: WaldoDB[]
            waldos = waldos_ as WaldoDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Bars_array = bars

            // clear the map that counts Bar in the GET
            FrontRepoSingloton.Bars_batch.clear()

            bars.forEach(
              bar => {
                FrontRepoSingloton.Bars.set(bar.ID, bar)
                FrontRepoSingloton.Bars_batch.set(bar.ID, bar)
              }
            )

            // clear bars that are absent from the batch
            FrontRepoSingloton.Bars.forEach(
              bar => {
                if (FrontRepoSingloton.Bars_batch.get(bar.ID) == undefined) {
                  FrontRepoSingloton.Bars.delete(bar.ID)
                }
              }
            )

            // sort Bars_array array
            FrontRepoSingloton.Bars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Waldos_array = waldos

            // clear the map that counts Waldo in the GET
            FrontRepoSingloton.Waldos_batch.clear()

            waldos.forEach(
              waldo => {
                FrontRepoSingloton.Waldos.set(waldo.ID, waldo)
                FrontRepoSingloton.Waldos_batch.set(waldo.ID, waldo)
              }
            )

            // clear waldos that are absent from the batch
            FrontRepoSingloton.Waldos.forEach(
              waldo => {
                if (FrontRepoSingloton.Waldos_batch.get(waldo.ID) == undefined) {
                  FrontRepoSingloton.Waldos.delete(waldo.ID)
                }
              }
            )

            // sort Waldos_array array
            FrontRepoSingloton.Waldos_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            bars.forEach(
              bar => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            waldos.forEach(
              waldo => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Bar.Waldos redeeming
                {
                  let _bar = FrontRepoSingloton.Bars.get(waldo.Bar_WaldosDBID.Int64)
                  if (_bar) {
                    if (_bar.Waldos == undefined) {
                      _bar.Waldos = new Array<WaldoDB>()
                    }
                    _bar.Waldos.push(waldo)
                    if (waldo.Bar_Waldos_reverse == undefined) {
                      waldo.Bar_Waldos_reverse = _bar
                    }
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // BarPull performs a GET on Bar of the stack and redeem association pointers 
  BarPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.barService.getBars()
        ]).subscribe(
          ([ // insertion point sub template 
            bars,
          ]) => {
            // init the array
            FrontRepoSingloton.Bars_array = bars

            // clear the map that counts Bar in the GET
            FrontRepoSingloton.Bars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bars.forEach(
              bar => {
                FrontRepoSingloton.Bars.set(bar.ID, bar)
                FrontRepoSingloton.Bars_batch.set(bar.ID, bar)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear bars that are absent from the GET
            FrontRepoSingloton.Bars.forEach(
              bar => {
                if (FrontRepoSingloton.Bars_batch.get(bar.ID) == undefined) {
                  FrontRepoSingloton.Bars.delete(bar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // WaldoPull performs a GET on Waldo of the stack and redeem association pointers 
  WaldoPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.waldoService.getWaldos()
        ]).subscribe(
          ([ // insertion point sub template 
            waldos,
          ]) => {
            // init the array
            FrontRepoSingloton.Waldos_array = waldos

            // clear the map that counts Waldo in the GET
            FrontRepoSingloton.Waldos_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            waldos.forEach(
              waldo => {
                FrontRepoSingloton.Waldos.set(waldo.ID, waldo)
                FrontRepoSingloton.Waldos_batch.set(waldo.ID, waldo)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Bar.Waldos redeeming
                {
                  let _bar = FrontRepoSingloton.Bars.get(waldo.Bar_WaldosDBID.Int64)
                  if (_bar) {
                    if (_bar.Waldos == undefined) {
                      _bar.Waldos = new Array<WaldoDB>()
                    }
                    _bar.Waldos.push(waldo)
                    if (waldo.Bar_Waldos_reverse == undefined) {
                      waldo.Bar_Waldos_reverse = _bar
                    }
                  }
                }
              }
            )

            // clear waldos that are absent from the GET
            FrontRepoSingloton.Waldos.forEach(
              waldo => {
                if (FrontRepoSingloton.Waldos_batch.get(waldo.ID) == undefined) {
                  FrontRepoSingloton.Waldos.delete(waldo.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getBarUniqueID(id: number): number {
  return 31 * id
}
export function getWaldoUniqueID(id: number): number {
  return 37 * id
}
