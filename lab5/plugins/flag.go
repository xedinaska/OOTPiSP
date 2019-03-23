package plugins

type Flag []string

func (f *Flag) String() string {
	return "pluginFlag"
}

func (f *Flag) Set(value string) error {
	*f = append(*f, value)
	return nil
}
