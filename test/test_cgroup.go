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
        // the containere process !!
        fmt.Printf("current pid %d", syscall.Getpid())
        fmt.Println()

        cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
        cmd.SysProcAttr = &syscall.SysProcAttr{}
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
    cmd := exec.Command("/proc/self/exe")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err:= cmd.Start(); err != nil {
        fmt.Println("error ", err)
        os.Exit(1)
    } else{
        // get the forked process's pid
        fmt.Printf("%v", cmd.Process.Pid)
        // create cgroup on the Hierarchy
        os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemlimit"), 0755)
        // add the container process to this cgroup
        ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemlimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
        // limit the cgroup-process 
        ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemlimit", "memory.limit_in_bytes"), []byte("100m"), 0644)
    }
    cmd.Process.Wait()
}
