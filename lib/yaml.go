package lib

import (
	"fmt"
	"io/ioutil"
	"os"

	"scsstools/lib/actions"

	"gopkg.in/yaml.v2"
)

type YAML struct {
	Actions []actions.TAction
}

func ParseYAMLActions(yamlfile string) (*YAML, error) {
	y := &YAML{}

	f, err := os.Open(yamlfile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, y)
	if err != nil {
		return nil, err
	}

	for _, act := range y.Actions {

		if !IsAction(act.Action) {
			return nil, fmt.Errorf("parse yaml: the %s action %s is not valid", act.Name, act.Action)
		}
	}

	return y, nil
}
