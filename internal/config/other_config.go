package config

type OtherConfig struct {
	Environment     string `mapstructure:"environment"`
	DefaultLanguage string `mapstructure:"default_lang"`
	BundleDirAbs    string `mapstructure:"bundle_dir_abs"`
}
