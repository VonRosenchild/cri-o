// +build !linux

package libpod

import (
	"github.com/containers/libpod/v2/libpod/define"
)

func (v *Volume) mount() error {
	return define.ErrNotImplemented
}

func (v *Volume) unmount(force bool) error {
	return define.ErrNotImplemented
}
