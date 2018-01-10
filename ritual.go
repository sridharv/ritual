package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func loadYAML(fileName string) (map[interface{}]interface{}, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	m := map[interface{}]interface{}{}
	if err = yaml.Unmarshal(data, &m); err != nil {
		return nil, errors.WithStack(err)
	}
	return m, nil
}

func getMapList(m map[interface{}]interface{}, key string) []map[interface{}]interface{} {
	v := m[key]
	if ml, ok := v.([]interface{}); ok {
		c := make([]map[interface{}]interface{}, len(ml))
		for i, val := range ml {
			if c[i], ok = val.(map[interface{}]interface{}); !ok {
				panic(fmt.Sprintf("Expected map, got %T", val))
			}
		}
		return c
	}
	return nil
}

func getString(m map[interface{}]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func printFile() error {
	m, err := loadYAML(os.Args[1])
	if err != nil {
		return err
	}
	items := getMapList(m, "items")
	for _, item := range items {
		name := getString(item, "name")
		if name != "" {
			fmt.Println(name)
		}
	}
	return nil
}

// Ritual should be able to list all items in a file
func main() {
	if err := printFile(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
