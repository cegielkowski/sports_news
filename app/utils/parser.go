package utils

import (
	"encoding/xml"
	"sports_news/domain"
)

type XmlAdaptors interface {
	domain.NewListInformation | domain.NewsArticleInformation
}

// ReadXml Generic function to read xml.
func ReadXml[T XmlAdaptors](data []byte, x T) (T, error) {
	err := xml.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}

	return x, nil
}
