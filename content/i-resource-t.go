package content

func NewResource(resourceType Material) *Resource {
	return &Resource{
		Type: resourceType,
	}
}

