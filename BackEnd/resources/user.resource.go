package resources

type UserResource struct {
	Id       int    `json:"Id,omitempty"`
	UserName string `json:"Username"`
	Password string `json:"Password"`
}
