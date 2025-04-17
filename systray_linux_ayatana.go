//go:build linux && !legacy_appindicator
// +build linux,!legacy_appindicator

package mysystray

/*
#cgo linux pkg-config: ayatana-appindicator3-0.1

#include "systray.h"
*/
import "C"
