import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { BarService } from '../bar.service'
import { getBarUniqueID } from '../front-repo.service'
import { WaldoService } from '../waldo.service'
import { getWaldoUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-issuerenaming-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbFromBackService: CommitNbFromBackService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private barService: BarService,
    private waldoService: WaldoService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.barService.BarServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.waldoService.WaldoServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Bar part of the mat tree
      */
      let barGongNodeStruct: GongNode = {
        name: "Bar",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Bar",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(barGongNodeStruct)

      this.frontRepo.Bars_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Bars_array.forEach(
        barDB => {
          let barGongNodeInstance: GongNode = {
            name: barDB.Name,
            type: GongNodeType.INSTANCE,
            id: barDB.ID,
            uniqueIdPerStack: getBarUniqueID(barDB.ID),
            structName: "Bar",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          barGongNodeStruct.children!.push(barGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Waldos
          */
          let WaldosGongNodeAssociation: GongNode = {
            name: "(Waldo) Waldos",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: barDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Bar",
            associationField: "Waldos",
            associatedStructName: "Waldo",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          barGongNodeInstance.children.push(WaldosGongNodeAssociation)

          barDB.Waldos?.forEach(waldoDB => {
            let waldoNode: GongNode = {
              name: waldoDB.Name,
              type: GongNodeType.INSTANCE,
              id: waldoDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getBarUniqueID(barDB.ID)
                + 11 * getWaldoUniqueID(waldoDB.ID),
              structName: "Waldo",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            WaldosGongNodeAssociation.children.push(waldoNode)
          })

        }
      )

      /**
      * fill up the Waldo part of the mat tree
      */
      let waldoGongNodeStruct: GongNode = {
        name: "Waldo",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Waldo",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(waldoGongNodeStruct)

      this.frontRepo.Waldos_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Waldos_array.forEach(
        waldoDB => {
          let waldoGongNodeInstance: GongNode = {
            name: waldoDB.Name,
            type: GongNodeType.INSTANCE,
            id: waldoDB.ID,
            uniqueIdPerStack: getWaldoUniqueID(waldoDB.ID),
            structName: "Waldo",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          waldoGongNodeStruct.children!.push(waldoGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbFromBackService.getCommitNbFromBack().subscribe(
      commitNbFromBack => {
        this.commitNbFromBack = commitNbFromBack
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_issuerenaming_go_table: ["github_com_fullstack_lang_issuerenaming_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_issuerenaming_go_table: ["github_com_fullstack_lang_issuerenaming_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_issuerenaming_go_editor: ["github_com_fullstack_lang_issuerenaming_go-" + structName.toLowerCase() + "-detail", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_issuerenaming_go_editor: ["github_com_fullstack_lang_issuerenaming_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_issuerenaming_go_editor: ["github_com_fullstack_lang_issuerenaming_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}
