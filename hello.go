package helloworld

import "go.uber.org/zap"

func newLogger() (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	log, _ := config.Build()
	return log.Sugar(), nil
}

func add(a, b int) int {

	_, _ = newLogger()

	c := a + b

	return c
}

func mul(a, b int) int {
	c := a * b
	return c
}
