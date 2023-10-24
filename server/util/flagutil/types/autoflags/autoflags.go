package autoflags

import (
	"flag"
	"log"
	"net/url"
	"reflect"
	"time"

	flagtypes "github.com/buildbuddy-io/buildbuddy/server/util/flagutil/types"
	flagyaml "github.com/buildbuddy-io/buildbuddy/server/util/flagutil/yaml"
)

type Taggable interface {
	Tag(name string)
}

type secretTag struct{}

func (_ *secretTag) Tag(name string) {
	log.Fatalf(
		"'secretTag.Taggable' was called for flag '%s', but is not meant to be "+
			"called for 'secretTag'; tagging a flag as secret is instead handled by "+
			"'New' or 'Var'.",
		name,
	)
}

var SecretTag = &secretTag{}

type deprecatedTag struct {
	migrationPlan string
}

func (d *deprecatedTag) Tag(name string) {
	log.Fatalf(
		"'deprecatedTag.Taggable' was called for flag '%s', but is not meant to "+
			"be called for 'deprecatedTag'; tagging a flag as deprecated is instead "+
			"handled by 'New' or 'Var'.",
		name,
	)
}

func DeprecatedTag(migrationPlan string) Taggable {
	return &deprecatedTag{migrationPlan: migrationPlan}
}

type yamlIgnoreTag struct{}

func (_ *yamlIgnoreTag) Tag(name string) {
	flagyaml.IgnoreFlagForYAML(name)
}

var YAMLIgnoreTag = &yamlIgnoreTag{}

// New declares a new flag named `name` with the specified value `defaultValue`
// of type `T` and the help text `usage`. It returns a pointer to where the
// value is stored. Tags may be used to mark flags; for example, use
// `SecretTag` to mark a flag that contains a secret that should be redacted in
// output, or use `DeprecatedTag(migrationPlan)` to mark a flag that has been
// deprecated and provide its migration plan.
func New[T any](flagset *flag.FlagSet, name string, defaultValue T, usage string, tags ...Taggable) *T {
	value := reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Interface().(*T)
	Var(flagset, value, name, defaultValue, usage, tags...)
	return value
}

// Var declares a new flag named `name` with the specified value `defaultValue`
// of type `T` stored at the pointer `value` and the help text `usage`. Tags may
// be used to mark flags; for example, use `SecretTag` to mark a flag that
// contains a secret that should be redacted in output, or use
// `DeprecatedTag(migrationPlan)` to mark a flag that has been deprecated and
// provide its migration plan.
func Var[T any](flagset *flag.FlagSet, value *T, name string, defaultValue T, usage string, tags ...Taggable) {
	switch v := any(value).(type) {
	case *bool:
		flagset.BoolVar(v, name, any(defaultValue).(bool), usage)
	case *time.Duration:
		flagset.DurationVar(v, name, any(defaultValue).(time.Duration), usage)
	case *float64:
		flagset.Float64Var(v, name, any(defaultValue).(float64), usage)
	case *int:
		flagset.IntVar(v, name, any(defaultValue).(int), usage)
	case *int64:
		flagset.Int64Var(v, name, any(defaultValue).(int64), usage)
	case *uint:
		flagset.UintVar(v, name, any(defaultValue).(uint), usage)
	case *uint64:
		flagset.Uint64Var(v, name, any(defaultValue).(uint64), usage)
	case *string:
		flagset.StringVar(v, name, any(defaultValue).(string), usage)
	case *[]string:
		flagtypes.StringSliceVar(flagset, v, name, any(defaultValue).([]string), usage)
	case *url.URL:
		flagtypes.URLVar(flagset, v, name, any(defaultValue).(url.URL), usage)
	default:
		if reflect.TypeOf(value).Elem().Kind() == reflect.Slice {
			flagtypes.JSONSliceVar(flagset, value, name, defaultValue, usage)
			break
		}
		if reflect.TypeOf(value).Elem().Kind() == reflect.Struct {
			flagtypes.JSONStructVar(flagset, value, name, defaultValue, usage)
			break
		}
		log.Fatalf("Var was called from flag registry for flag %s with value %v of unrecognized type %T.", name, defaultValue, defaultValue)
	}
	for _, tg := range tags {
		switch v := any(tg).(type) {
		case *secretTag:
			flagtypes.Secret[T](flagset, name)
		case *deprecatedTag:
			flagtypes.Deprecate[T](flagset, name, v.migrationPlan)
		default:
			tg.Tag(name)
		}
	}
}
