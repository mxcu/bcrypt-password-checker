package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {

	password := ""
	hash := ""
	cost := bcrypt.DefaultCost
	passbytes := []byte{}
	file := ""

	flag.IntVar(&cost, "cost", cost, "cost (aka rounds)")
	flag.StringVar(&file, "file", "", "file to scan for matching hashes, - for STDIN")
	flag.StringVar(&hash, "hash", "", "bcrypt hash")
	flag.StringVar(&password, "password", "", "the password to check hash against")
	flag.Parse()

	if password == "" && file != "-" {
		fmt.Print("password: ")
		raw, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(raw) == 0 {
			fmt.Println("password is empty")
			os.Exit(1)
		}
		passbytes = raw
	} else {
		passbytes = []byte(password)
	}

	switch {

	case len(hash) > 0 && file == "":
		if err := bcrypt.CompareHashAndPassword([]byte(hash), passbytes); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("password matches")
		}

	case len(hash) == 0 && file == "":
		if cost < bcrypt.MinCost {
			fmt.Printf("cost too low (min=%v)\n", bcrypt.MinCost)
			os.Exit(1)
		}
		if cost > bcrypt.MaxCost {
			fmt.Printf("cost too high (max=%v)\n", bcrypt.MaxCost)
			os.Exit(1)
		}
		hash, err := bcrypt.GenerateFromPassword(passbytes, cost)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(hash))
		}

	case file != "":
		var in *os.File
		if file == "-" {
			in = os.Stdin
		} else {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			in = f
		}
		scanner := bufio.NewScanner(in)
		bcryptRe := regexp.MustCompile(`\$[\w]{1,4}\$\d{1,2}\$[./A-Za-z0-9]{53}`)
		for scanner.Scan() {
			line := scanner.Bytes()
			matches := bcryptRe.FindAll(line, -1)
			if matches == nil {
				continue
			}
			for _, match := range matches {
				if bcrypt.CompareHashAndPassword(match, passbytes) == nil {
					fmt.Printf("%v\n", string(match))
				}
			}
		}

	default:
		flag.Usage()
	}
}
