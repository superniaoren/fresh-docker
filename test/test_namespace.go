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
            Cloneflags: syscall.CLONE_NEWUTS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err:= cmd.Run(); err != nil {
        log.Fatal(err)
    }
}
