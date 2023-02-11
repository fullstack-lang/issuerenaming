// insertion point sub template for components imports 
  import { BarsTableComponent } from './bars-table/bars-table.component'
  import { BarSortingComponent } from './bar-sorting/bar-sorting.component'
  import { WaldosTableComponent } from './waldos-table/waldos-table.component'
  import { WaldoSortingComponent } from './waldo-sorting/waldo-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfBarsComponents: Map<string, any> = new Map([["BarsTableComponent", BarsTableComponent],])
  export const MapOfBarSortingComponents: Map<string, any> = new Map([["BarSortingComponent", BarSortingComponent],])
  export const MapOfWaldosComponents: Map<string, any> = new Map([["WaldosTableComponent", WaldosTableComponent],])
  export const MapOfWaldoSortingComponents: Map<string, any> = new Map([["WaldoSortingComponent", WaldoSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Bar", MapOfBarsComponents],
      ["Waldo", MapOfWaldosComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Bar", MapOfBarSortingComponents],
      ["Waldo", MapOfWaldoSortingComponents],
    ]
  )
