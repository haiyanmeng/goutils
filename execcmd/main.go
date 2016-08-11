package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func execCmd(name string, s ...string) (string, error) {
	cmd := exec.Command(name, s...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", stderr.String())
		return "", err
	}

	return out.String(), nil
}

func main() {
	test1()
	//syscallExec()
}

func syscallExec() {
	binary := "/usr/local/sbin/runc"
	args := []string{"/usr/local/sbin/runc", "create", "--bundle", "/home/hmeng/go/src/github.com/hmeng-19/myrunc/c4", "1"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func test1() {
	fmt.Printf("pid: %v\n", os.Getpid())
	//cmd := exec.Command("/usr/local/sbin/runc", "create", "--bundle", "/home/hmeng/go/src/github.com/hmeng-19/myrunc/c4", "1")
	cmd := exec.Command("date")

	//	var stdout bytes.Buffer
	//	var stderr bytes.Buffer
	//	cmd.Stdout = &stdout
	//	cmd.Stderr = &stderr

	/*
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
	*/

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(stderr)
	line, _, err := r.ReadLine()
	fmt.Println(line)

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	//	CSIGNAL := uintptr(0x000000ff)              /* signal mask to be sent at exit */
	//	CLONE_VM := uintptr(0x00000100)             /* set if VM shared between processes */
	//	CLONE_FS := uintptr(0x00000200)             /* set if fs info shared between processes */
	//	CLONE_FILES := uintptr(0x00000400)          /* set if open files shared between processes */
	//	CLONE_SIGHAND := uintptr(0x00000800)        /* set if signal handlers and blocked signals shared */
	//	CLONE_PTRACE := uintptr(0x00002000)         /* set if we want to let tracing continue on the child too */
	//	CLONE_VFORK := uintptr(0x00004000)          /* set if the parent wants the child to wake it up on mm_release */
	//	CLONE_PARENT := uintptr(0x00008000)         /* set if we want to have the same parent as the cloner */
	//	CLONE_THREAD := uintptr(0x00010000)         /* Same thread group? */
	//	CLONE_NEWNS := uintptr(0x00020000)          /* New mount namespace group */
	//	CLONE_SYSVSEM := uintptr(0x00040000)        /* share system V SEM_UNDO semantics */
	//	CLONE_SETTLS := uintptr(0x00080000)         /* create a new TLS for the child */
	//	CLONE_PARENT_SETTID := uintptr(0x00100000)  /* set the TID in the parent */
	//	CLONE_CHILD_CLEARTID := uintptr(0x00200000) /* clear the TID in the child */
	//	CLONE_DETACHED := uintptr(0x00400000)       /* Unused, ignored */
	//	CLONE_UNTRACED := uintptr(0x00800000)       /* set if the tracing process can't force CLONE_PTRACE on this clone */
	//	CLONE_CHILD_SETTID := uintptr(0x01000000)   /* set the TID in the child */
	//	CLONE_NEWCGROUP := uintptr(0x02000000)      /* New cgroup namespace */
	//	CLONE_NEWUTS := uintptr(0x04000000)         /* New utsname namespace */
	//	CLONE_NEWIPC := uintptr(0x08000000)         /* New ipc namespace */
	//	CLONE_NEWUSER := uintptr(0x10000000)        /* New user namespace */
	//	CLONE_NEWPID := uintptr(0x20000000)         /* New pid namespace */
	//	CLONE_NEWNET := uintptr(0x40000000)         /* New network namespace */
	//	CLONE_IO := uintptr(0x80000000)             /* Clone io context */

	/*
				CLONE_VM := uintptr(0x00000100)
				CLONE_FS := uintptr(0x00000200)
				CLONE_FILES := uintptr(0x00000400)
				CLONE_THREAD := uintptr(0x00010000)
			CLONE_SIGHAND := uintptr(0x00000800)
			SIGCHLD := uintptr(17)

		cmd.SysProcAttr = &syscall.SysProcAttr{
			//Cloneflags: CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_SIGHAND | CLONE_THREAD,
			Cloneflags: uintptr(17) | uintptr(0x00008000),
		}
	*/

}
