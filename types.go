package eclosion

type appPaths struct {
	rootPath string
	dirNames []string
}

type CookieConfig struct {
	Name     string
	Lifetime string
	Persist  string
	Secure   string
	Domain   string
}
