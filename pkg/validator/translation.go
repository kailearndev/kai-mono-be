package validator

import "errors"

type TransactionLike interface {
	GetLang() string
}

func ValidateUniqueLang[T TransactionLike](items []T) error {
	langMap := make(map[string]bool)
	for _, item := range items {
		lang := item.GetLang()
		if _, exists := langMap[lang]; exists {
			return errors.New("duplicate lang found: " + lang)
		}
		langMap[lang] = true
	}
	return nil
}
