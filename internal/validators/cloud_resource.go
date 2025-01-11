package validators

type AddCloudResourcesInput struct {
	Resources []string `json:"resources" binding:"required"`
}

type UpdateCloudResourceInput struct {
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Region string `json:"region" binding:"required" `
}
