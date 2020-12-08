package globalFilepath

import (
	"path/filepath"
	"runtime"
	"sync"
)

var (
	base  string
	mutex sync.RWMutex
)

// Init produces a base path that will be used in the other functions.
// It uses the calling package as the base directory and then correctly
// joins any given suffix arguments.
func Init(suffixes ...string) (err error) {
	defer mutex.Unlock()
	mutex.Lock()

	if _, filename, _, ok := runtime.Caller(1); ok {
		base = filepath.Dir(filename)
		base = join(suffixes)
	} else {
		err = ErrCannotLoadBasePath
	}

	return
}

// Get returns the base path.
func Get() string {
	defer mutex.RUnlock()
	mutex.RLock()

	return base
}

func join(suffixes []string) (value string) {
	value = base

	if len(suffixes) > 0 {
		suffix := filepath.Join(suffixes...)
		value = filepath.Join(value, suffix)
	}

	return
}

// Join returns the base path with the suffixes arguments correctly joined.
func Join(suffixes ...string) string {
	defer mutex.RUnlock()
	mutex.RLock()

	return join(suffixes)
}
