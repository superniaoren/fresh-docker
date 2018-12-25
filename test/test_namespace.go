package main

import ( "os/exec"
         "syscall"
         "os"
         "log")

/*
$go build -o bin/test_namespace test_namespace.go
$sudo ./bin/test_namespace
*/

func main(){
    cmd := exec.Command("sh")
    cmd.SysProcAttr = &syscall.SysProcAttr{ 
            //Cloneflags: syscall.CLONE_NEWUTS,
            //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
            //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
            // mount namespace
            Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err:= cmd.Run(); err != nil {
        log.Fatal(err)
    }
}
