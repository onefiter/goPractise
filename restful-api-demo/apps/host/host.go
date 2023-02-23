package host

import "context"

type Service interface {
	// 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// 查询主机详情
	DescribeHost(context.Context, *QueryHostRequest) (*Host, error)
	// 主机更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// 主机删除, 比如前端需要 打印当前删除主机的Ip或者其他信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

type HostSet struct {
	Items []*Host
	Total int
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

// Host模型的定义
type Host struct {
	// 资源公共属性部分
	*Resource
	// 资源独有属性部分
	*Describe
}

type Vendor int

const (
	// 枚举的默认值
	PRIVATE_IDC Vendor = iota
	// 阿里云
	ALIYUN
	// 腾讯云
	TXYUN
	// AWS
	AWS
)

type Resource struct {
	Id          string            `json:"id"`          // 全局唯一id
	Vendor      Vendor            `json:"vendor"`      // 厂商
	Region      string            `json:"region"`      // 地域
	CreateAt    int64             `json:"create_at"`   // 创建时间
	ExpireAt    int64             `json:"expire_at"`   // 过期时间
	Type        string            `json:"type"`        // 规格
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	SyncAt      int64             `json:"sync_at"`     // 同步时间
	Account     string            `json:"accout"`      // 资源的所属账号
	PublicIP    string            `json:"public_ip"`   // 公网IP
	PrivateIP   string            `json:"private_ip"`  // 内网IP
	PayType     string            `json:"pay_type"`    // 实例付费方式

}

type Describe struct {
	CPU          int    `json:"cpu"`           // 核数
	Memory       int    `json:"memory"`        // 内存
	GPUAmount    int    `json:"gpu_amount"`    // GPU数量
	GPUSpec      string `json:"gpu_spec"`      // GPU类型
	OSType       string `json:"os_type"`       // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`       // 操作系统名称
	SerialNumber string `json:"serial_number"` // 序列号
}

type QueryHostRequest struct {
}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct {
	Id string
}
