package malloc

func Malloc(size uintptr) (uintptr, error) {
	return PlatformMalloc(size)
}

func Free(size uintptr) error {
	return PlatformFree(size)
}
