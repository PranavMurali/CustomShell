package main

import (
	"bufio"
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	rand "math/rand"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"
)

var history[]string
var htime[]string

func main() {


	reader := bufio.NewReader(os.Stdin)
	color.Cyan("yes")
	user,err:=user.Current()
	tmps:=user.Username
	if err != nil {
		panic(err)
	}
	cmd:=exec.Command("toilet",tmps)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()

	for {
		fmt.Print("🔥🐲> ")
		input, err := reader.ReadString('\n')
		history = append(history, input)
		dt := time.Now()
		dtf := dt.Format("01-02-2006 15:04:05")
		htime = append(htime, dtf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
        return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
        err := binary.Read(crand.Reader, binary.BigEndian, &v)
        if err != nil {
                log.Fatal(err)
        }
        return v
}

func execInput(input string,) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {

	case "cd":	

		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])

	case "vscode":

		if len(args) < 2 {
			return errors.New("path required")
		}

		cmd := exec.Command("code", args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()

	case "nano":
		if len(args) < 2 {
			return errors.New("path required")
		}

		cmd := exec.Command("nano", args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()

	case "touch":

		if len(args) < 2 {
			return errors.New("path required")
		}

		cmd := exec.Command("touch", args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()

	case "golang":

		if len(args) < 2 {
			return errors.New("path required")
		}

		cmd := exec.Command("go", args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()

	case "userinfo":
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
		fmt.Println("Username: " + user.Username)
		fmt.Println("Home Dir: " + user.HomeDir)
		fmt.Println("Real User: " + os.Getenv("SUDO_USER"))
		return nil

	case "wther":
		tmp:="http://wttr.in/chennai"
		cmd := exec.Command("curl", tmp)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	
	case "art":
		var src cryptoSource
		rnd := rand.New(src)
		ndate:=rnd.Intn(30)
		date:=strconv.Itoa(ndate)
		if ndate<10{
			date="0"+date	
		}
		tmp:="http://samiare.net/daily/1901"+date+"?width=20"
		cmd := exec.Command("curl", tmp)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()

	case "ospref":
		if len(args) < 2 {
			return errors.New("name required")
		}
	
	case "ls":
		cmd := exec.Command("ls", args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		
		return cmd.Run()

	case "history":
		var tmp []string
		//strings.Join(htime, "\n")
		i := 0
		for i < len(history) {
			tmp = append(tmp, history[i])
			tmp = append(tmp, "     ")
			tmp = append(tmp, htime[i])
			tmp = append(tmp, "\n")
			i = i + 1
		}
		cmd := exec.Command("echo", strings.Join(tmp, ""))
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		
		return cmd.Run()
	

	case "exit":
		os.Exit(0)

	}
	

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
