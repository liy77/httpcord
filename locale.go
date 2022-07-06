package httpcord

type (
	Locale string
	Dictionary map[Locale]string
)

const (
	EnglishUSLocale    Locale = "en-US"
	EnglishGBLocale    Locale = "en-GB"
	BulgarianLocale    Locale = "bg"
	ChineseCNLocale    Locale = "zh-CN"
	ChineseTWLocale    Locale = "zh-TW"
	CroatianLocale     Locale = "hr"
	CzechLocale        Locale = "cs"
	DanishLocale       Locale = "da"
	DutchLocale        Locale = "nl"
	FinnishLocale      Locale = "fi"
	FrenchLocale       Locale = "fr"
	GermanLocale       Locale = "de"
	GreekLocale        Locale = "el"
	HindiLocale        Locale = "hi"
	HungarianLocale    Locale = "hu"
	ItalianLocale      Locale = "it"
	JapaneseLocale     Locale = "ja"
	KoreanLocale       Locale = "ko"
	LithuanianLocale   Locale = "lt"
	NorwegianLocale    Locale = "no"
	PolishLocale       Locale = "pl"
	PortugueseBRLocale Locale = "pt-BR"
	RomanianLocale     Locale = "ro"
	RussianLocale      Locale = "ru"
	SpanishESLocale    Locale = "es-ES"
	SwedishLocale      Locale = "sv-SE"
	ThaiLocale         Locale = "th"
	TurkishLocale      Locale = "tr"
	UkrainianLocale    Locale = "uk"
	VietnameseLocale   Locale = "vi"
)
