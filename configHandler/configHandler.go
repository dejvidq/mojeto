package configHandler

type config struct {
	path string
}

func New(path string) config {
	conf := config {path: path}
	return conf
}

func (conf config) ShowConfigPath() string {
	return conf.path
}
