// insertion point sub template for components imports 
  import { FoosTableComponent } from './foos-table/foos-table.component'
  import { FooSortingComponent } from './foo-sorting/foo-sorting.component'
  import { WaldosTableComponent } from './waldos-table/waldos-table.component'
  import { WaldoSortingComponent } from './waldo-sorting/waldo-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfFoosComponents: Map<string, any> = new Map([["FoosTableComponent", FoosTableComponent],])
  export const MapOfFooSortingComponents: Map<string, any> = new Map([["FooSortingComponent", FooSortingComponent],])
  export const MapOfWaldosComponents: Map<string, any> = new Map([["WaldosTableComponent", WaldosTableComponent],])
  export const MapOfWaldoSortingComponents: Map<string, any> = new Map([["WaldoSortingComponent", WaldoSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Foo", MapOfFoosComponents],
      ["Waldo", MapOfWaldosComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Foo", MapOfFooSortingComponents],
      ["Waldo", MapOfWaldoSortingComponents],
    ]
  )
