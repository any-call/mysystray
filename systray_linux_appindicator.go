//go:build linux && legacy_appindicator
// +build linux,legacy_appindicator

package mysystray

/*
#cgo linux pkg-config: appindicator3-0.1
#cgo linux CFLAGS: -DUSE_LEGACY_APPINDICATOR

#include "systray.h"
*/
import "C"
