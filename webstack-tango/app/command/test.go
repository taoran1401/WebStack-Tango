package command

import (
	"go.uber.org/zap"
)

type TestCmd struct {
	Log *zap.SugaredLogger
}

func NewTestCmd(log *zap.SugaredLogger) *TestCmd {
	return &TestCmd{Log: log}
}

func (this *TestCmd) Handle() {

}
