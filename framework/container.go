package framework

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// lessContainer 是服务容器的具体实现
type lessContainer struct {
	Container
	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// NewlessContainer 创建一个服务容器
func NewlessContainer() *lessContainer {
	return &lessContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (less *lessContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range less.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做了绑定
func (less *lessContainer) Bind(provider ServiceProvider) error {
	less.lock.Lock()
	key := provider.Name()

	less.providers[key] = provider
	less.lock.Unlock()

	// if provider is not defer
	if provider.IsDefer() == false {
		if err := provider.Boot(less); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(less)
		method := provider.Register(less)
		instance, err := method(params...)
		if err != nil {
			fmt.Println("bind service provider ", key, " error: ", err)
			return errors.New(err.Error())
		}
		less.instances[key] = instance
	}
	return nil
}

func (less *lessContainer) IsBind(key string) bool {
	return less.findServiceProvider(key) != nil
}

func (less *lessContainer) findServiceProvider(key string) ServiceProvider {
	less.lock.RLock()
	defer less.lock.RUnlock()
	if sp, ok := less.providers[key]; ok {
		return sp
	}
	return nil
}

func (less *lessContainer) Make(key string) (interface{}, error) {
	return less.make(key, nil, false)
}

func (less *lessContainer) MustMake(key string) interface{} {
	serv, err := less.make(key, nil, false)
	if err != nil {
		panic("container not contain key " + key)
	}
	return serv
}

func (less *lessContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return less.make(key, params, true)
}

func (less *lessContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(less); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(less)
	}
	method := sp.Register(less)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// 真正的实例化一个服务
func (less *lessContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	less.lock.RLock()
	defer less.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := less.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return less.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := less.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := less.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	less.instances[key] = inst
	return inst, nil
}

// NameList 列出容器中所有服务提供者的字符串凭证
func (less *lessContainer) NameList() []string {
	ret := []string{}
	for _, provider := range less.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}
