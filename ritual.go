package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Item holds a single item which can be a task or some other thing that needs to be planned.
type Item struct {
	Name string                 `yaml:"name"`
	Done bool                   `yaml:"done"`
	Rest map[string]interface{} `yaml:",inline"`
}

// Collection holds a single logical collection of tasks and other things.
type Collection struct {
	Items []Item                 `yaml:"items"`
	Rest  map[string]interface{} `yaml:",inline"`
}

func loadCollectionFile(fileName string) (*Collection, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var c Collection
	if err = yaml.Unmarshal(data, &c); err != nil {
		return nil, errors.WithStack(err)
	}
	return &c, nil
}

func printFile() error {
	c, err := loadCollectionFile(os.Args[1])
	if err != nil {
		return err
	}
	for _, item := range c.Items {
		if item.Name == "" || item.Done {
			continue
		}
		fmt.Println(item.Name)
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
