package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"github.com/jessevdk/go-flags"
	"path/filepath"
)


var Options struct {
	Provider string `short:"p" long:"provider" required:"yes" description:"Which provider to use e.g. aws or gcc"`
	ConfigFile flags.Filename `short:"c" long:"config" required:"yes" description:"Path to the config file"`
}



func main() {
	validateCommandLineOptions()

	createTerraformFile()

	runTerraform()
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}

func validateCommandLineOptions(){
	// parse cli flags -> break if mandatory flags are missing
	_, err := flags.Parse(&Options)
	check(err)

	// config file is mandatory is should exist
	if _, err := os.Stat(string(Options.ConfigFile)); os.IsNotExist(err) {
		fmt.Println("Config file: no such file: ", Options.ConfigFile)
		return
	}

	if !(Options.Provider == "aws" || Options.Provider == "gcc") {
		fmt.Println("Unknown provider: ", Options.Provider)
		return
	}
}

func runTerraform(){
	// first we need to download the terraform module
	cmd := exec.Command("terraform", "get", "-update=true")
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	err := cmd.Run()
	fmt.Printf("%s", out.String())
	fmt.Printf("%s", outErr.String())
	check(err)


	// then we apply terraform to execute our module with the given parameters
	var out2 bytes.Buffer
	var outErr2 bytes.Buffer
	cmd = exec.Command("terraform", "apply",("-state="+string(Options.ConfigFile)+".tfstate") )
	cmd.Stdout = &out2
	cmd.Stderr = &outErr2
	err =cmd.Run()
	fmt.Printf("%s", out2.String())
	fmt.Printf("%s", outErr2.String())
	check(err)


	// print out the terraform results
	if _, err := os.Stat(filepath.Join(".", "terraform.tfstate")); os.IsNotExist(err) {
		return
	}
	dat, err := ioutil.ReadFile(filepath.Join(".", "terraform.tfstate"))
	check(err)
	fmt.Print(string(dat))
}

func createTerraformFile(){
	config, err := ioutil.ReadFile(string(Options.ConfigFile))
	check(err)


	// create terraform file which will calls our terraform module
	file, err := os.Create(filepath.Join(".", "gcc_terraform.tf"))
	check(err)
	defer func() {
		err = file.Close()
		check(err)
	}()


	file.WriteString("module \"ms-infra-terraform\" { \n")
	file.WriteString(string(config))
	source := strings.Replace("\nsource = \"git::https://gitlab.trifork.se/flg/ms-infra-terraform-ccp.git//{{provider}}?ref=master\"","{{provider}}", Options.Provider,1)
	file.WriteString(source)
	file.WriteString("\n}\n")

	// TODO flush needed?
	fmt.Println("Created terraform file: ", file.Name())
}

