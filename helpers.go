package eclosion

import "os"

func (c *Eclosion) CreateDirIfNotExists(path string) error {

	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)

		if err != nil {
			return err
		}

	}
	return nil
}

func (c *Eclosion) CreateFileIfNotExists(path string) error {

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}

func coalesce(v, def string) string {
	if v == "" {
		return def
	}
	return v
}
func ternary(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}