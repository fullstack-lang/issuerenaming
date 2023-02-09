package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Foo:
		if stage.OnAfterFooCreateCallback != nil {
			stage.OnAfterFooCreateCallback.OnAfterCreate(stage, target)
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
	case *Foo:
		newTarget := any(new).(*Foo)
		if stage.OnAfterFooUpdateCallback != nil {
			stage.OnAfterFooUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Foo:
		if stage.OnAfterFooDeleteCallback != nil {
			staged := any(staged).(*Foo)
			stage.OnAfterFooDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Foo:
		if stage.OnAfterFooReadCallback != nil {
			stage.OnAfterFooReadCallback.OnAfterRead(stage, target)
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
	case *Foo:
		stage.OnAfterFooUpdateCallback = any(callback).(OnAfterUpdateInterface[Foo])
	
	case *Waldo:
		stage.OnAfterWaldoUpdateCallback = any(callback).(OnAfterUpdateInterface[Waldo])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Foo:
		stage.OnAfterFooCreateCallback = any(callback).(OnAfterCreateInterface[Foo])
	
	case *Waldo:
		stage.OnAfterWaldoCreateCallback = any(callback).(OnAfterCreateInterface[Waldo])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Foo:
		stage.OnAfterFooDeleteCallback = any(callback).(OnAfterDeleteInterface[Foo])
	
	case *Waldo:
		stage.OnAfterWaldoDeleteCallback = any(callback).(OnAfterDeleteInterface[Waldo])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Foo:
		stage.OnAfterFooReadCallback = any(callback).(OnAfterReadInterface[Foo])
	
	case *Waldo:
		stage.OnAfterWaldoReadCallback = any(callback).(OnAfterReadInterface[Waldo])
	
	}
}
