package datatransfers

type ServiceInfo struct {
	ID            string `json:"id,omitempty" binding:"-"`
	ApplicationID string `json:"-" binding:"-"`

	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Endpoint    string `json:"endpoint" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Config      string `json:"config" binding:"required"` // TODO: needs standard format, check via marshalling
	Showcase    bool   `json:"showcase" binding:"required"`

	CreatedAt int64 `json:"createdAt" binding:"-"`
	UpdatedAt int64 `json:"updatedAt" binding:"-"`
}
