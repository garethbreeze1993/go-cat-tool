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
            Usage: "Whether user wants lines numbered",
            Required: false,
        },
        },
        Action: func(cCtx *cli.Context) error {

                lineCount := 1

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
//                             fmt.Println("Print line by line \n")

                        } else {
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