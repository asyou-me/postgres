package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	inpath := flag.String("in", "", "配置文件路径")
	outpath := flag.String("o", "", "配置文件路径")

	flag.Parse()

	if (*inpath) == "" {
		Usage()
		return
	}

	if (*outpath) == "" {
		Usage()
		return
	}

	var err error
	if string((*inpath)[0]) != "/" {
		var curr_path string = ""
		curr_path, err = filepath.Abs(filepath.Dir(os.Args[0]))
		*inpath = curr_path + "/" + *inpath
	}

	if string((*outpath)[0]) != "/" {
		var curr_path string = ""
		curr_path, err = filepath.Abs(filepath.Dir(os.Args[0]))
		*outpath = curr_path + "/" + *outpath
	}

	obj, err := getJson(*inpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fout, err := os.Create(*outpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	tmpl, err := template.New("type").Parse(tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(fout, obj)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getJson(path string) (*InputType, error) {
	fi, err := os.Open(path)
	var obj InputType = InputType{}

	if err != nil {
		return &obj, errors.New("路径" + path + "不存在")
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return &obj, err
	}

	if err = json.Unmarshal(fd, &obj); err != nil {
		if err != nil {
			return &obj, err
		}
	}
	return &obj, nil

}
