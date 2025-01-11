package data

import "github.com/babyfaceeasy/crims/internal/models"

var CloudResources = []models.CloudResource{
	{Name: "Compute Engine", Type: "VM", Region: "us-central1"},
	{Name: "Storage Bucket", Type: "Storage", Region: "us-west1"},
}
