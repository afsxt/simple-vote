package constants

type ThemeState int

const (
	ThemeNotStarted ThemeState = iota
	ThemeStarted
	ThemeFinished
	ThemeUnknown
)
