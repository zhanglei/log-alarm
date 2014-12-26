// log
package main

type Log struct {
	FileName string
	Dir      string
	Email    []string
	LastPos  int64
}

func Newlog(fname string, dir string, email []string) *Log {
	return &Log{fname, dir, email, 0}
}
