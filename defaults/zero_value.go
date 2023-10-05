package defaults

// sometimes you just want to return the zero value for a type as a one-liner
func ZeroValue[OUTPUT_TYPE any]() OUTPUT_TYPE {
	var output OUTPUT_TYPE
	return output
}
