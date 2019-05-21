package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/tersergo/terser-cli/schema"
	"github.com/tersergo/terser-cli/tpl"
)

func main() {

	dsn := "root:root@tcp(localhost:3306)/open_campus?parseTime=true"
	dbName := "open_campus"

	query := schema.NewDBQuery(dbName, dsn)

	tableList, err := query.GetDBSchema()
	if err != nil {
		fmt.Println("schema query: ", err.Error())
		return
	}

	generateDBConfig(query)

	newTemp, err := template.New("model").Parse(tpl.ModelTemplate)
	if err != nil {
		fmt.Println("template error: ", err.Error())
	}
	modelTpl := template.Must(newTemp, err)

	generateTime := time.Now().Format("2006/01/02 15:04:05")

	for tableName, table := range tableList {
		table.GenerateTime = generateTime
		tpl, _ := modelTpl.Clone()
		fmt.Println(tableName, "\t\t=>\t", table.StructName)

		var stream bytes.Buffer
		err = tpl.Execute(&stream, table)
		if err != nil {
			fmt.Println("template error: ", err.Error())
			continue
		} else {
			writeFile("model", table.FileName, stream)
		}
	}
}

func writeFile(dir, file string, stream bytes.Buffer) {
	os.Mkdir("../"+dir, os.ModePerm)

	filePath := filepath.Join("../", dir, file+".go")
	data, err := format.Source(stream.Bytes())
	if err != nil {
		fmt.Println("stream error: " + err.Error())
		return
	}
	ioutil.WriteFile(filePath, data, os.ModePerm)
}

func generateDBConfig(query *schema.Query) {

	dbTemp, err := template.New("db_config").Parse(tpl.DBConfigTemplate)
	if err != nil {
		fmt.Println("template error: ", err.Error())
	}
	tpl := template.Must(dbTemp, err)

	var stream bytes.Buffer
	err = tpl.Execute(&stream, query)
	if err != nil {
		fmt.Println("db_config template error: ", err.Error())
	}

	writeFile("model", "db_config", stream)

}
