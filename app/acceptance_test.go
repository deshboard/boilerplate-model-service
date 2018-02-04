// +build acceptance

package app

import (
	"github.com/deshboard/boilerplate-model-service/pkg/app"
	"github.com/goph/fxt/dev"
	"github.com/goph/fxt/test"
	fxacceptance "github.com/goph/fxt/test/acceptance"
	"go.uber.org/fx"
)

func init() {
	dev.LoadEnvFromFile("../.env.test")
	dev.LoadEnvFromFile("../.env.dist")

	runnerFactoryRegistry.Register(test.RunnerFactoryFunc(AcceptanceRunnerFactory))
}

func AcceptanceRunnerFactory() (test.Runner, error) {
	acceptanceRunner := test.NewGodogRunner()

	config, err := newConfig()
	if err != nil {
		panic(err)
	}

	dbContext := new(fxacceptance.DbContext)
	appContext := fxacceptance.NewAppContext(
		fx.NopLogger,
		fx.Provide(func() Config { return config }, newApplicationInfo),
		Module,
		fx.Populate(&dbContext.DB),
	)

	acceptanceRunner.RegisterFeatureContext(appContext.BeforeFeatureContext)
	app.RegisterSuite(acceptanceRunner)
	acceptanceRunner.RegisterFeatureContext(appContext.AfterFeatureContext)

	return acceptanceRunner, nil
}
