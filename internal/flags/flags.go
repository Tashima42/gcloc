package flags

import (
	"reflect"

	"github.com/spf13/pflag"
)

type Flag struct {
	ShortName    string
	Kind         reflect.Kind
	DefaultValue interface{}
	Description  string
}

func InitFlags(flagSet *pflag.FlagSet, flags map[string]Flag) {
	for flagName, flag := range flags {
		switch flag.Kind {
		case reflect.String:
			flagSet.StringP(flagName, flag.ShortName, flag.DefaultValue.(string), flag.Description)
		case reflect.Slice:
			flagSet.StringArrayP(flagName, flag.ShortName, flag.DefaultValue.([]string), flag.Description)
		}
	}
}
