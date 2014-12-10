package main
import (
  "syscall"
  "os"
  "os/exec"
)

func main() {
  lockfile, found := syscall.Getenv("CRONLOCK_LOCKFILE")
  if !found {
    lockfile = "lock.cronlock"
  }

  fd,err := syscall.Open(lockfile, syscall.O_RDWR | syscall.O_CREAT, 0666)
  if err != nil{
    print("Cronlock: can't open ["+lockfile+"] for locking")
    print(err)
    syscall.Exit(1)
  }
  err = syscall.Flock(fd, syscall.LOCK_EX | syscall.LOCK_NB)
  if err != nil{
    print("Cronlock: can't acquire ["+lockfile+"]")
    syscall.Exit(1)
  }

  out, err := exec.Command(os.Args[1], os.Args[2:]...).Output()

  if err != nil{
    print("Cronlock: failed command ["+err.Error()+"]")
    syscall.Exit(1)
  }
  println("Cronlock: finished running. Output:\n"+string(out))
}


