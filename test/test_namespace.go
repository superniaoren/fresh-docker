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
            // new ipc
            //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
            // new pid
            //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
            // mount namespace
            //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
            // new user
            //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
            // new network
            Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | 
                         syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
    }
    // root user ?; must work with CLONE_NEWUSER
    // fork/exec /bin/sh: operation not permitted
    //cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(1), Gid: uint32(1)}

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err:= cmd.Run(); err != nil {
        log.Fatal(err)
    }
    os.Exit(-1)
}
