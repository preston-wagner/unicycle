package multithread

type filterResult[OUTPUT_TYPE any] struct {
	value OUTPUT_TYPE
	ok    bool
}

type errorResult[OUTPUT_TYPE any] struct {
	value OUTPUT_TYPE
	err   error
}
