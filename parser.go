package xodr

import (
	"encoding/xml"
)

func IsXodrFormat([]byte) bool {
	return true
}

func ParseOdrXml(b []byte, od *OpenDrive) error {
	err := xml.Unmarshal(b, od)
	if err != nil {
		return err
	}
	return nil
}
