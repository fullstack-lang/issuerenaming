package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Bar:
		if stage.OnAfterBarCreateCallback != nil {
			stage.OnAfterBarCreateCallback.OnAfterCreate(stage, target)
		}
	case *Waldo:
		if stage.OnAfterWaldoCreateCallback != nil {
			stage.OnAfterWaldoCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Bar:
		newTarget := any(new).(*Bar)
		if stage.OnAfterBarUpdateCallback != nil {
			stage.OnAfterBarUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Waldo:
		newTarget := any(new).(*Waldo)
		if stage.OnAfterWaldoUpdateCallback != nil {
			stage.OnAfterWaldoUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Bar:
		if stage.OnAfterBarDeleteCallback != nil {
			staged := any(staged).(*Bar)
			stage.OnAfterBarDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Waldo:
		if stage.OnAfterWaldoDeleteCallback != nil {
			staged := any(staged).(*Waldo)
			stage.OnAfterWaldoDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Bar:
		if stage.OnAfterBarReadCallback != nil {
			stage.OnAfterBarReadCallback.OnAfterRead(stage, target)
		}
	case *Waldo:
		if stage.OnAfterWaldoReadCallback != nil {
			stage.OnAfterWaldoReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarUpdateCallback = any(callback).(OnAfterUpdateInterface[Bar])
	
	case *Waldo:
		stage.OnAfterWaldoUpdateCallback = any(callback).(OnAfterUpdateInterface[Waldo])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarCreateCallback = any(callback).(OnAfterCreateInterface[Bar])
	
	case *Waldo:
		stage.OnAfterWaldoCreateCallback = any(callback).(OnAfterCreateInterface[Waldo])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarDeleteCallback = any(callback).(OnAfterDeleteInterface[Bar])
	
	case *Waldo:
		stage.OnAfterWaldoDeleteCallback = any(callback).(OnAfterDeleteInterface[Waldo])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarReadCallback = any(callback).(OnAfterReadInterface[Bar])
	
	case *Waldo:
		stage.OnAfterWaldoReadCallback = any(callback).(OnAfterReadInterface[Waldo])
	
	}
}
