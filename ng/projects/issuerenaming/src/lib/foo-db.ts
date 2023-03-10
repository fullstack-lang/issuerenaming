// insertion point for imports
import { WaldoDB } from './waldo-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FooDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Waldos?: Array<WaldoDB>
}
