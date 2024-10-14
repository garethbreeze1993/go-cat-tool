package main

import (
    "os"
    "github.com/urfave/cli/v2"
    "log"
    "io"
    "fmt"
    "bufio"
)


func main() {
    app := &cli.App{
        Name:  "go-cat-tool",
        Usage: "My version of the cat tool",
        Flags:[]cli.Flag{
            &cli.BoolFlag{
            Name: "number",
            Aliases: []string{"n"},
            Usage: "Whether user wants all lines numbered",
            Required: false,
        },
            &cli.BoolFlag{
            Name: "nonblank",
            Aliases: []string{"b"},
            Usage: "Whether user wants only non-blank lines numbered",
            Required: false,
        },
        },
        Action: func(cCtx *cli.Context) error {

                lineCount := 1

                if cCtx.Bool("number") && cCtx.Bool("nonblank") {
                    panic("Cannot choose both number and nonblank flags choose either one or none at all")
                }

                for index, _ := range os.Args[1:] {
                    source := cCtx.Args().Get(index)
//                     fmt.Printf("Hello %q \n", cCtx.Args().Get(index))

                    if !validateFile(source) {
		                panic("File not found")
	                }

	                if len(source) > 0 {
	                    srcFile, err := os.Open(source)

                        if err != nil {
                            panic(err.Error())
                        }

                        if cCtx.Bool("number"){

                            fileScanner := bufio.NewScanner(srcFile)

                            fileScanner.Split(bufio.ScanLines)

                            for fileScanner.Scan() {
                                fmt.Printf("%d. %s\n", lineCount, fileScanner.Text())
                                lineCount++
                            }

                        } else if cCtx.Bool("nonblank"){

                            fileScanner := bufio.NewScanner(srcFile)
                            fileScanner.Split(bufio.ScanLines)
                            for fileScanner.Scan() {
                                if len(fileScanner.Text()) == 0 {
                                    fmt.Printf("%s\n", fileScanner.Text())
                                } else {
	                            fmt.Printf("%d. %s\n", lineCount, fileScanner.Text())
	                            lineCount++
	                            }
                            }
                        }else {
                            io.Copy(os.Stdout, srcFile)

                        }

                        defer srcFile.Close()
                    }
                }

            return nil
        },

    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}


func validateFile(file string) bool {

    if len(file) <= 0 {
        return true
    }

    if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true

}