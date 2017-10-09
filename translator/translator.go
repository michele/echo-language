package translator

import (
	"github.com/labstack/echo"
	"github.com/michele/http_accept_language/language"
)

const languageKey = "github.com/michele/echo-translator/key"

func SetDefault(lang string) {
	language.Default = lang
}

func TranslatorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		langs := language.ParseHeader(c.Request().Header.Get("Accept-Language"))

		c.Set(languageKey, langs)
		return next(c)
	}
}

func GetLanguages(c echo.Context) language.Languages {
	if c == nil || c.Get(languageKey) == nil {
		return language.ParseHeader(language.Default)
	}
	return c.Get(languageKey).(language.Languages)
}
