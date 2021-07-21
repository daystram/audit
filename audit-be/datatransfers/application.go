package datatransfers

type ApplicationInfo struct {
	ID       string        `json:"id,omitempty" binding:"-"`
	Services []ServiceInfo `json:"services,omitempty" binding:"-"`

	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`

	CreatedAt int64 `json:"createdAt,omitempty" binding:"-"`
	UpdatedAt int64 `json:"updatedAt,omitempty" binding:"-"`
}
