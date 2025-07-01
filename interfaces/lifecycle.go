package interfaces

type ILifecycle interface {
	// Configure initializes this service.  Call it before
	// you use this service!
	Configure()

	// IsActive returns energy saving's status;
	// false means app is inactive,
	// true means app is active.
	//
	// Call it in fyne's thread only!
	IsActive() bool
}
