package framework

// IGroup 代表前缀分组
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	// 实现嵌套group
	Group(string) IGroup

	// 嵌套中间件
	Use(middlewares ...ControllerHandler)
}

// Group struct 实现了IGroup

type Group struct {
	core   *Core  // 指向core结构
	parent *Group // 指向上一个Group，如果存在的话
	prefix string // 整个group的通用前缀

	middlewares []ControllerHandler
}

// 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// 实现Get方法
func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers)
}

// 实现Post方法
func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handlers)
}

// 实现Put方法
func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handlers)
}

// 实现Delete方法
func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handlers)
}

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

// 获取某个group的middleware
// 这里就是获取了除了Get/Post/Put/Delete之外的middleware

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 实现Group方法
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

// 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}
