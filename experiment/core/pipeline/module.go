package pipeline

import (
	"github.com/kubesphere/kubekey/experiment/core/cache"
	"github.com/kubesphere/kubekey/experiment/core/config"
	"github.com/kubesphere/kubekey/experiment/core/logger"
)

type Module interface {
	Default(runtime *config.Runtime)
	Init()
	Is() string
	Run() error
}

type BaseTaskModule struct {
	Name    string
	Tasks   []Task
	Cache   *cache.Cache
	Runtime *config.Runtime
}

func (t *BaseTaskModule) Default(runtime *config.Runtime) {
	if t.Name == "" {
		t.Name = DefaultModuleName
	}

	logger.Log.SetModule(t.Name)
	t.Runtime = runtime
	t.Cache = cache.NewCache()
}

func (t *BaseTaskModule) Init() {
}

func (t *BaseTaskModule) Is() string {
	return TaskModule
}

func (t *BaseTaskModule) Run() error {
	logger.Log.Info("Begin Run")
	for i := range t.Tasks {
		task := t.Tasks[i]
		task.Init(t.Runtime, t.Cache)
		if err := task.Execute(); err != nil {
			return err
		}
	}
	return nil
}
