package bigcache

import (
	"fmt"
	"github.com/allegro/bigcache"
	"testing"
	"time"
)

// 默认初始化
func TestInitDefaultCache(t *testing.T) {
	// 创建一个LifeWindow为5秒的cache实例

	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Second * 5))
	defer cache.Close()
	// 设置缓存
	err := cache.Set("key1", []byte("hello word"))
	if err != nil {
		t.Errorf("设置缓存失败:%v", err)
	}
	// 获取缓存
	data, err := cache.Get("key1")
	if err != nil {
		t.Errorf("获取缓存失败:%v", err)
	}
	fmt.Printf("获取结果:%s\n", data)
}

// 创建自定义缓存
func TestInitCustom(t *testing.T) {
	// 指定创建属性
	config := bigcache.Config{
		// 设置分区的数量，必须是2的整倍数
		Shards: 1024,
		// LifeWindow后,缓存对象被认为不活跃,但并不会删除对象
		LifeWindow: 5 * time.Second,
		// CleanWindow后，会删除被认为不活跃的对象，0代表不操作；
		CleanWindow: 3 * time.Second,
		// 设置最大存储对象数量，仅在初始化时可以设置
		//MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntriesInWindow: 1,
		// 缓存对象的最大字节数，仅在初始化时可以设置
		MaxEntrySize: 500,
		// 是否打印内存分配信息
		Verbose: true,
		// 设置缓存最大值(单位为MB),0表示无限制
		HardMaxCacheSize: 8192,
		// 在缓存过期或者被删除时,可设置回调函数，参数是(key、val)，默认是nil不设置
		OnRemove: func(key string, entry []byte) {

		},
		// 在缓存过期或者被删除时,可设置回调函数，参数是(key、val,reason)默认是nil不设置
		OnRemoveWithReason: nil,
	}
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		t.Error(err)
	}
	defer cache.Close()
	// 设置缓存
	_ = cache.Set("key1", []byte("hello word"))
	// 验证CleanWindow是否生效
	time.Sleep(10 * time.Second)
	// 获取缓存
	data, err := cache.Get("key1")
	if err != nil {
		t.Errorf("获取缓存失败:%v", err)
	}
	fmt.Printf("获取结果:%s\n", data)
	fmt.Println("运行结束！")
}

func TestSetAndGet(t *testing.T) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	// 设置缓存
	err := cache.Set("key1", []byte("php"))
	if err != nil {
		t.Errorf("设置缓存失败:%v", err)
	}
	_ = cache.Set("key2", []byte("go"))
	// 获取缓存
	for _, key := range []string{"key1", "key2"} {
		if data, err := cache.Get(key); err == nil {
			fmt.Printf("key: %s 结果:%s\n", key, data)
		}
	}
}

func TestDelCache(t *testing.T) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	key := "key"
	// 设置
	_ = cache.Set(key, []byte("111"))
	// 删除
	_ = cache.Delete(key)
	// 获取
	if _, err := cache.Get(key); err != nil {
		fmt.Println(err)
	}
}
