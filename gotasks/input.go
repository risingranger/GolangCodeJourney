package file
import (
    "os"
    "syscall"
)
type File struct {
    fd   int    // file descriptor number
    name string // file name at Open time
}
func newFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    return &File{fd, name}
}

n := new(File)
n.fd = fd
n.name = name
return n