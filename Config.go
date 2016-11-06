package main

type Config struct {
	basepath string

}

func (conf *Config) load() {
	conf.basepath = `C:\Users\xime\Desktop\`
}

