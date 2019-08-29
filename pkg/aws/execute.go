package aws

import (
	"github.com/deislabs/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

func (m *Mixin) loadAction() (*Action, error) {
	var action Action
	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &action)
		return &action, err
	})
	return &action, err
}

func (m *Mixin) Execute() error {
	action, err := m.loadAction()
	if err != nil {
		return err
	}

	output, err := builder.ExecuteSingleStepAction(m.Context, action)
	if err != nil {
		return err
	}
	step := action.Steps[0]

	return builder.ProcessJsonPathOutputs(m.Context, step, output)
}
