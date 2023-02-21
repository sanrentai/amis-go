package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".md") {
			ss := file.Name()
			filename := strings.Split(ss, ".")
			fmt.Println(filename[0])
			run(filename[0])
			// break
		}
	}
	// names := []string{
	// 	"color",
	// }
	// run(names[0])
}

func run(filename string) {
	modname := ""
	result := `package form
`
	f, err := os.OpenFile("./"+filename+".md", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	sxb := false
	sx := false
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if strings.HasPrefix(string(line), "title: ") {
			datas := strings.Split(string(line), " ")
			modname = strings.ReplaceAll(datas[1], "-", "")
			zs := modname
			if len(datas) == 3 {
				zs = datas[2]
			}
			result += fmt.Sprintf(`
// %s
func %s(opts ...opt) map[string]interface{} {
	return newForm("%s", opts...)
}

`, zs, modname, filename)
		}
		if strings.HasPrefix(string(line), "## 属性表") {
			sxb = true
			continue
		}
		if sxb {

			if strings.HasPrefix(string(line), "| ---") {
				sx = true
				continue
			}
			if strings.HasPrefix(string(line), "##") {
				sx = false
				continue
			}
			if sx {
				// fmt.Println(string(line))
				// fmt.Println(string(line))
				datas := strings.Split(string(line), "|")
				if len(datas) != 6 {
					continue
				}
				// fmt.Println(datas[0])
				// fmt.Println(datas[1])
				// fmt.Println(datas[2])
				// fmt.Println(datas[3])
				// fmt.Println(datas[4])
				// fmt.Println(datas[5])
				fieldname := strings.Trim(datas[1], " ")
				if fieldname == "type" {
					continue
				}
				result += fmt.Sprintf(`

// %s
func %s_%s(p %s) opt {
	return func(o map[string]interface{}) {
		o["%s"] = p
	}
} `, strings.Trim(datas[4], " "), modname, fieldname, tn(strings.Trim(datas[2], " ")), fieldname)

			}
		}

	}

	ioutil.WriteFile("../../../form/"+filename+".go", []byte(result), 0666)

}

func tn(t string) string {
	if t == "`string`" {
		return "string"
	} else if t == "`boolean`" {
		return "bool"
	}
	return "interface{}"
}
