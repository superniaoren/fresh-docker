package main

import (
    "os/exec"
    "path"
    "os"
    "fmt"
    "io/ioutil"
    "syscall"
    "strconv"
)

// the root dir of hierarch (of sys/mem)
const  cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main(){
    if os.Args[0] == "/proc/self/exe" {
        fmt.Printf("current pid %d", syscall.Getpid())
        fmt.Println()
    }
}
