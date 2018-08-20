package main

import "fmt"
import "os"
import "os/exec"

func main() {
	fmt.Println("start")
	// cmd := exec.Cmd{}

	cmd := exec.Command("docker", "exec", "-it", "bopo", "bash")
	// cmd := exec.Command("/bin/bash", "-c", "docker exec -it bopo bash")
	// cmd := exec.Command("/usr/local/bin/docker exec -it bopo bash")
	// cmd := exec.Command("sleep", "10")
	// cmd := exec.Command("./script.sh")
	cmd.Stdout = os.Stdout
	cmd.Start()
	// cmd.Run()
	// cmd.Wait()
	// fmt.Println("err=", err)
}
