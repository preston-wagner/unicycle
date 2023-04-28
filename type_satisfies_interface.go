package unicycle

func TypeSatisfiesInterface[TYPE any, INTERFACE any]() bool {
	var instance TYPE
	_, ok := any(instance).(INTERFACE)
	return ok
}
