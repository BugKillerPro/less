package fk

// IGroup 代表前缀分组
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ...Handler)
	Post(string, ...Handler)
	Put(string, ...Handler)
	Delete(string, ...Handler)

	// 实现嵌套group
	Group(string) IGroup

	// 嵌套中间件
	Use(middlewares ...Handler)
}

// Group struct 实现了IGroup
type Group struct {
	core   *Engine // 指向core结构
	parent *Group  //指向上一个Group，如果有的话
	prefix string  // 这个group的通用前缀

	middlewares []Handler // 存放中间件
}

// 初始化Group
func NewGroup(core *Engine, prefix string) *Group {
	return &Group{
		core:        core,
		parent:      nil,
		prefix:      prefix,
		middlewares: []Handler{},
	}
}

// 实现Get方法
func (g *Group) Get(uri string, handlers ...Handler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}

// 实现Post方法
func (g *Group) Post(uri string, handlers ...Handler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}

// 实现Put方法
func (g *Group) Put(uri string, handlers ...Handler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(uri, allHandlers...)
}

// 实现Delete方法
func (g *Group) Delete(uri string, handlers ...Handler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(uri, allHandlers...)
}

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// 获取某个group的middleware
// 这里就是获取除了Get/Post/Put/Delete之外设置的middleware
func (g *Group) getMiddlewares() []Handler {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 实现 Group 方法
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

// 注册中间件
func (g *Group) Use(middlewares ...Handler) {
	g.middlewares = middlewares
}
