package metadata

import (
	"archive/zip"
	"encoding/xml"
	"strings"
)

type OfficeCoreProperty struct {
	XMLName        xml.Name `xml:"coreProperties"`
	Creator        string   `xml:"creator"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
}
type OfficeAppProperty struct {
	XMLName     xml.Name `xml:"Properties"`
	Application string   `xml:"Application"`
	Company     string   `xml:"Company"`
	Version     string   `xml:"AppVersion"`
}

var OfficeVersions = map[string]string{
	"16": "2016",
	"15": "2013",
	"14": "2010",
	"12": "2007",
	"11": "2003",
}

func (a *OfficeAppProperty) GetMajorVersion() string {
	tokens := strings.Split(a.Version, ".")
	if len(tokens) < 2 {
		return "Unknown"
	}
	v, ok := OfficeVersions[tokens[0]]
	if !ok {
		return "Unknown"
	}
	return v
}

// NewProperties() получает *zipReader, который представляет io.Reader для ZIP-архивов
func NewProperties(r *zip.Reader) (*OfficeCoreProperty, *OfficeAppProperty, error) {
	var coreProps OfficeCoreProperty
	var appProps OfficeAppProperty

	for _, f := range r.File { // С помощью экземпляра zip.Reader выполняется перебор всех файлов архива
		switch f.Name { // с проверкой их имен
		case "docProps/core.xml":
			if err := process(f, &coreProps); err != nil {
				return nil, nil, err
			}

		case "docProps/app.xml":
			if err := process(f, &appProps); err != nil {
				return nil, nil, err
			}
		default:
			continue
		}
	}
	return &coreProps, &appProps, nil
}

func process(f *zip.File, prop interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}

	defer rc.Close()

	if err := xml.NewDecoder(rc).Decode(&prop); err != nil {
		return err
	}
	return nil
}
