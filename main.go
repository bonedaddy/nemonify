package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := newApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Author = "bonedaddy"
	app.Name = "nemonify"
	app.Usage = "takes input, turns into mnemonic phrase"
	return app
}

func cmdFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "file.name, fn",
			Usage: "the file to read data from",
		},
		cli.StringFlag{
			Name:  "save.path, sp",
			Usage: "the path to save data to",
			Value: "data.out",
		},
	}
}

func generate() cli.Command {
	return cli.Command{
		Name:  "generate",
		Usage: "generates a mnemonic phrase from a file",
		Action: func(c *cli.Context) error {
			if c.String("file.name") == "" {
				return errors.New("file.name flag is empty")
			}
			return generateMnemonic(
				c.String("file.name"),
				c.String("save.path"),
			)
		},
		Flags: cmdFlags(),
	}
}

func decode() cli.Command {
	return cli.Command{
		Name:  "decode",
		Usage: "decode a mnemonic phrase",
		Action: func(c *cli.Context) error {
			return decodeMnemonic(
				c.String("file.name"),
				c.String("save.path"),
			)
		},
		Flags: cmdFlags(),
	}
}

func generateMnemonic(fileName, savePath string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	data := Encode(fbytes)
	phrase, err := ToMnemonic(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(
		savePath,
		[]byte(phrase),
		os.FileMode(0600),
	)
}

func decodeMnemonic(fileName, savePath string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	data, err := Decode(string(fbytes))
	if err != nil {
		return err
	}
	contents, err := FromMnemonic(string(data))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(
		savePath,
		[]byte(contents),
		os.FileMode(0600),
	)
}
