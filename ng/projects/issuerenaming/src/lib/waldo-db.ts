// insertion point for imports
import { FooDB } from './foo-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class WaldoDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Foo_WaldosDBID: NullInt64 = new NullInt64
	Foo_WaldosDBID_Index: NullInt64  = new NullInt64 // store the index of the waldo instance in Foo.Waldos
	Foo_Waldos_reverse?: FooDB 

}
