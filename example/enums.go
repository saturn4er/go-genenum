package example

import (
	fmt "fmt"
	strconv "strconv"
)

type Platform byte

const (
	PlatformWindows Platform = iota + 1
	PlatformMac
	PlatformLinux
	PlatformAndroid
	PlatformiOS
)

func (p Platform) IsValid() bool {
	return p > 0 && p < 6
}

var MobilePlatforms = []Platform{
	PlatformAndroid,
	PlatformiOS,
}

func (p Platform) IsMobile() bool {
	if p < 1 || p > 6 {
		return false
	}
	return []bool{false, false, false, false, true, true}[p]
}

var DesktopPlatforms = []Platform{
	PlatformWindows,
	PlatformMac,
	PlatformLinux,
}

func (p Platform) IsDesktop() bool {
	if p < 1 || p > 6 {
		return false
	}
	return []bool{false, true, true, true, false, false}[p]
}
func (p Platform) IsWindows() bool {
	return p == PlatformWindows
}
func (p Platform) IsMac() bool {
	return p == PlatformMac
}
func (p Platform) IsLinux() bool {
	return p == PlatformLinux
}
func (p Platform) IsAndroid() bool {
	return p == PlatformAndroid
}
func (p Platform) IsiOS() bool {
	return p == PlatformiOS
}

type InvalidPlatformValueError byte

func (e InvalidPlatformValueError) Error() string {
	return fmt.Sprintf("invalid Platform(%d)", e)
}
func (p Platform) Validate() error {
	if p < 1 || p > 6 {
		return InvalidPlatformValueError(p)
	}
	return nil
}
func (p Platform) String() string {
	if p < 1 || p > 5 {
		return "Platform(" + strconv.FormatInt(int64(p), 10) + ")"
	}
	const names = "WindowsMacLinuxAndroidiOS"

	var indexes = [...]int32{0, 7, 10, 15, 22, 25}

	return names[indexes[p-1]:indexes[p]]
}
