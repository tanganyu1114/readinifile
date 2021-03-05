package main

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	iniFiles    string
	sectionName string
	keyName     string
)

func init() {
	flag.StringVar(&iniFiles, "f", "", "Plz enter the ini file one or more;eg:one.ini,two.ini")
	flag.StringVar(&sectionName, "s", "", "Plz enter the sectionName")
	flag.StringVar(&keyName, "k", "", "Plz enter the keyName")
	flag.Parse()
}

func main() {
	if iniFiles == "" {
		flag.Usage()
		return
	}
	files := strings.Split(iniFiles, ",")
	filei := make([]interface{}, 0)
	for _, file := range files {
		filei = append(filei, file)
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, files[0], filei[1:]...)
	if err != nil {
		fmt.Println(err)
	}
	val := cfg.Section(sectionName).Key(keyName).Value()
	fmt.Println(val)
}
