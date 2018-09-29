
// +build windows

package win

var (
	// Library
	libdnsapi uintptr

	// Functions
	dnsFlushResolverCache                   uintptr
)

func init() {
	// Library
	libdnsapi = doLoadLibrary("dnsapi.dll")

	// Functions
	dnsFlushResolverCache = doGetProcAddress(libdnsapi, "DnsFlushResolverCache")
}

func DnsFlushResolverCache() bool {
	ret1 := syscall3(dnsFlushResolverCache, 0, 0, 0, 0)
	return ret1 != 0
}