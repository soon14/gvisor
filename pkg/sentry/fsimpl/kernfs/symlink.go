// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kernfs

import (
	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/context"
	"gvisor.dev/gvisor/pkg/sentry/kernel/auth"
)

// StaticSymlink provides an Inode implementation for symlinks that point to
// a immutable target.
type StaticSymlink struct {
	InodeAttrsReadonly
	InodeNoopRefCount
	InodeSymlink

	target string
}

var _ Inode = (*StaticSymlink)(nil)

// NewStaticSymlink creates a new symlink file pointing to 'target'.
func NewStaticSymlink(creds *auth.Credentials, ino uint64, target string) *Dentry {
	inode := &StaticSymlink{}
	inode.Init(creds, ino, target)

	d := &Dentry{}
	d.Init(inode)
	return d
}

// Init initializes the instance.
func (s *StaticSymlink) Init(creds *auth.Credentials, ino uint64, target string) {
	s.target = target
	s.InodeAttrsReadonly.Init(creds, ino, linux.ModeSymlink|0777)
}

// Readlink implements Inode.
func (s *StaticSymlink) Readlink(_ context.Context) (string, error) {
	return s.target, nil
}
