package validators

type AddResourcesInput struct {
	CustomerID uint     `json:"customer_id" binding:"required"`
	Resources  []string `json:"resources" binding:"required"`
}
