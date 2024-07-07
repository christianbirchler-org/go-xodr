package xodr

import (
	"encoding/xml"
)

func ParseOdrXml(b []byte, od *OpenDrive) error {
	err := xml.Unmarshal(b, od)
	if err != nil {
		return err
	}
	return nil
}
