package create_resource_pattern

import (
	"fmt"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/hook"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/resource"
	"sync"
)

// ResourceManager 是资源对象的管理器
type ResourceManager struct {
	mu        sync.Mutex
	resources map[string]*resource.Resource
	beforeFn  []func(resource *resource.Resource) error
	afterFn   []func(resource *resource.Resource) error
}

// NewResourceManager 创建一个新的资源管理器
func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		resources: make(map[string]*resource.Resource),
	}
}

// RegisterHook 注册资源对象的钩子
func (rm *ResourceManager) RegisterHook(hook hook.ResourceHook) {
	rm.beforeFn = append(rm.beforeFn, hook.BeforeCreate)
	rm.afterFn = append(rm.afterFn, hook.AfterCreate)
}

// CreateResource 创建一个新的资源对象
func (rm *ResourceManager) CreateResource(id, name, email string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	var err error
	// 如果 hook 流程出错，默认不创建资源对象
	defer func() {
		if err != nil {
			delete(rm.resources, id)
		}
	}()

	if _, exists := rm.resources[id]; exists {
		return fmt.Errorf("resource with ID %s already exists", id)
	}

	rr := &resource.Resource{
		ID:    id,
		Name:  name,
		Email: email,
		// 设置其他属性...
	}

	// 执行创建前的钩子函数
	for _, beforeFn := range rm.beforeFn {
		err = beforeFn(rr)
		if err != nil {
			return err
		}
	}

	rm.resources[id] = rr

	// 执行创建后的钩子函数
	for _, afterFn := range rm.afterFn {
		err = afterFn(rr)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetResourceByID 根据资源 ID 获取资源对象
func (rm *ResourceManager) GetResourceByID(id string) (*resource.Resource, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	r, exists := rm.resources[id]
	if !exists {
		return nil, fmt.Errorf("resource with ID %s does not exist", id)
	}

	return r, nil
}

func (rm *ResourceManager) DeleteResourceByID(id string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	delete(rm.resources, id)
	return nil
}

// ListResources 返回所有资源对象的列表
func (rm *ResourceManager) ListResources() []*resource.Resource {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	resources := make([]*resource.Resource, 0, len(rm.resources))
	for _, r := range rm.resources {
		resources = append(resources, r)
	}

	return resources
}
