package build

import (
	"os"
)

func CreateDepsFile(path string, module string, license string) error {
	depsFile, err := os.Create(path + "/" + module + ".deps")
	if err != nil {
		return err
	}
	_, err = depsFile.WriteString(license)
	if err != nil {
		depsFile.Close()
		return err
	}
	err = depsFile.Close()
	if err != nil {
		return err
	}
	return nil
}
