package main

import (
     "fmt"
//     "io/ioutil"
//     "log"
	  "os"
     "gopkg.in/yaml.v3"
)

type Record struct {
     Item string `yaml:"item"`
     Col  string `yaml:"colour"`
     Size string `yaml:"size"`
}

type Config struct {
     Record Record `yaml:"Settings"`
}

func main() {

     config := Config{Record: Record{Item: "window", Col: "blue", Size: "small"}}

     data, err := yaml.Marshal(&config)

     if err != nil {
		fmt.Printf("error - marshall: %v\n", err)
		os.Exit(1)
//          log.Fatal(err)
     }

	outfil, err := os.Create("config.yaml")
	if err != nil {
		fmt.Printf("error open file: %v\n", err)
		os.Exit(1)
	}

	nb, err := outfil.Write(data)
	if err != nil {
		fmt.Printf("error write file: %v\n", err)
		os.Exit(1)
	}

	outfil.Close()
     fmt.Printf("success! %d data written!\n", nb)

// read yaml

	infil, err := os.Open("config.yaml")
	if err != nil {
		fmt.Printf("error open file read: %v\n", err)
		os.Exit(1)
	}

	filinfo, err := infil.Stat()
	if os.IsNotExist(err) {
		fmt.Printf("error file does not exist: %v\n", err)
		os.Exit(1)
	}
	nb = int(filinfo.Size())

	buf := make([]byte, nb)

	nr, err := infil.Read(buf)
	if err != nil {
		fmt.Printf("error could not read: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("read bytes: %d\n",nr)

	var yamlConfig Config

    err = yaml.Unmarshal(buf, &yamlConfig)
    if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
		os.Exit(1)
    }


    fmt.Printf("Result: %v\n", yamlConfig)
	fmt.Printf("item: %s\n", yamlConfig.Record.Item)
}
