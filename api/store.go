package api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

func SaveTokenInfos(infos *TokenInfos) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	dir := path.Join(usr.HomeDir, ".keepctl")
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		} else {
			return err
		}
	}
	file := path.Join(dir, "creds.json")
	os.Remove(file)

	b, err := json.Marshal(infos)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	return err
}

func LoadTokenInfos() (*TokenInfos, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	filename := path.Join(usr.HomeDir, ".keepctl", "creds.json")

	if _, err := os.Stat(filename); err == nil {
		file, e := ioutil.ReadFile(filename)
		if e != nil {
			return nil, err
		}
		var infos TokenInfos
		err = json.Unmarshal(file, &infos)
		return &infos, err
	}
	return nil, nil
}

func RemoveTokenInfos() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	filename := path.Join(usr.HomeDir, ".keepctl", "creds.json")
	if _, err := os.Stat(filename); err == nil {
		return os.Remove(filename)
	}
	return nil
}
