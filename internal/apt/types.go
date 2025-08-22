package apt

type AptRepository struct {
	Name       string
	Mode       string
	GPGUrl     string
	GPGKeyName string
	RepoUrl    string
	ToSign     bool
}
