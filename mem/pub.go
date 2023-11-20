// Package memory provides a single method reporting total system memory
// accessible to the kernel.
package mem

// TotalMemory returns the total accessible system memory in bytes.
//
// The total accessible memory is installed physical memory size minus reserved
// areas for the kernel and hardware, if such reservations are reported by
// the operating system.
//
// If accessible memory size could not be determined, then 0 is returned.
func TotalMemory() uint64 {
	return sysTotalMemory()
}

// FreeMemory returns the total free system memory in bytes.
//
// The total free memory is installed physical memory size minus reserved
// areas for other applications running on the same system.
//
// If free memory size could not be determined, then 0 is returned.
func FreeMemory() uint64 {
	return sysFreeMemory()
}

func PercentUsed() float64 {
	return 100 - PercentFree()
}

func PercentFree() float64 {
	total := sysTotalMemory()
	free := sysFreeMemory()
	return 100 * float64(free) / float64(total)
}
