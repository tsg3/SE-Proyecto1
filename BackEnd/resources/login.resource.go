package resources

// This is a Presentation Model Resource to Song Model
type LoginResource struct {
	Logged bool   `json:"Logged"`
	Token  string `json:"Token"`
}
