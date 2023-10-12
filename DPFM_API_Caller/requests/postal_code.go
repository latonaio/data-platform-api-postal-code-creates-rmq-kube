package requests

type PostalCode struct {
	PostalCode          string `json:"PostalCode"`
	Country             string `json:"Country"`
	LocalSubRegion      string `json:"LocalSubRegion"`
	LocalRegion         string `json:"LocalRegion"`
	GlobalRegion        string `json:"GlobalRegion"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
