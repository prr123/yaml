// parseYaml
// program that parses a yaml file and generates go lang code for reading
// author: prr, azul software
// date: 13 July 2022
// copyright 2022 prr, azul software
//
package main

import (
    "os"
    "fmt"
)


func main () {
    var inpFilNam string
    numarg := len(os.Args)

    switch numarg {

        case 1:
            fmt.Println("insufficient arguments - no input file")
            fmt.Println("usage: parseYaml inputfile!")
            os.Exit(-1)

        case 2:
            inpFilNam = os.Args[1]

        default:
            fmt.Println("too many arguments!")
            fmt.Println("usage: parseYaml inputfile!")
            os.Exit(-1)
    }

    fmt.Printf("input file: %s\n", inpFilNam)

    extPos := -1
    for i:= 0; i < len(inpFilNam); i++ {
        if inpFilNam[i] == '.' {
            extPos = i
            break
        }
    }

    extV := false
    if extPos > -1 {
        if (inpFilNam[(extPos +1):] == "yaml") {extV = true}
        if (inpFilNam[(extPos +1):] == "yml") {extV = true}
    } else {
        extV = true
        inpFilNam += ".yaml"
    }
    if !extV {
        fmt.Printf("invalid extension: %s\n", inpFilNam[extPos:])
        os.Exit(-1)
    }

    inpFilPath := "inpTest/" + inpFilNam

    fmt.Printf("input file path: %s\n", inpFilPath)

    


    fmt.Println("**** success ****")
}


