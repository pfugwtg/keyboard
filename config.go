package main

import (
	"gopkg.in/ini.v1"
)

var (
	cfg *ini.File
	section *ini.Section
)
const (
	CFG_KEY_RATE_POS_SELECTED = "key_rate_pos_selected"
	CFG_SECTION_NAME = "default"
	CFG_FILE_NAME = "config.ini"
)

func InitConfig() {
	var err error
	cfg, err = ini.Load(CFG_FILE_NAME)
    if err != nil {
        cfg = ini.Empty()
		section, _ = cfg.NewSection(CFG_SECTION_NAME)
		initKeyValues()
    } else {
		section = cfg.Section(CFG_SECTION_NAME)
	}
}

func initKeyValues() {
	section.Key(CFG_KEY_RATE_POS_SELECTED).SetValue("0")
	cfg.SaveTo(CFG_FILE_NAME)
}

func GetConfig(key string) string {
	return section.Key(key).String()
}

func SetConfig(key string, value string) {
	section.Key(key).SetValue(value)
	cfg.SaveTo(CFG_FILE_NAME)
}