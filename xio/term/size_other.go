// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris

package term

// Size returns the number of columns and rows comprising the terminal.
func Size() (columns, rows int) {
	return defColumns, defRows
}
