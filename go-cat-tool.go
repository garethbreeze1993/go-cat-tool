package main

import (
    "os"
    "github.com/urfave/cli/v2"
    "log"
    "io"
)


func main() {
    app := &cli.App{
        Name:  "go-cat-tool",
        Usage: "My version of the cat tool",
        Flags:[]cli.Flag{
            &cli.StringFlag{
            Name: "port",
            Aliases: []string{"p"},
            Usage: "Port number to check",
            Required: false,
        },
        },
        Action: func(cCtx *cli.Context) error {


//                 fmt.Printf("Port %q \n", cCtx.String("port"))

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

                        io.Copy(os.Stdout, srcFile)
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