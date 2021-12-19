// Package gflag is a go1.18 wrapper to provide simple generics based API for
// defining command line flags.
package gflag

import (
	"flag"
	"time"
)

// FlagConstraint defines the constraints for the Define and DefineVar type
// paramater.
type FlagConstraint interface {
	bool | time.Duration | float64 | int | int64 | string | uint | uint64
}

// FlagSet describes a flag set which can be wrapped by Define and DefineVar
// funcs.
type FlagSet interface {
	Bool(name string, value bool, usage string) *bool
	BoolVar(p *bool, name string, value bool, usage string)
	Duration(name string, value time.Duration, usage string) *time.Duration
	DurationVar(p *time.Duration, name string, value time.Duration, usage string)
	Float64(name string, value float64, usage string) *float64
	Float64Var(p *float64, name string, value float64, usage string)
	Int(name string, value int, usage string) *int
	Int64(name string, value int64, usage string) *int64
	Int64Var(p *int64, name string, value int64, usage string)
	IntVar(p *int, name string, value int, usage string)
	String(name string, value string, usage string) *string
	StringVar(p *string, name string, value string, usage string)
	Uint(name string, value uint, usage string) *uint
	Uint64(name string, value uint64, usage string) *uint64
	Uint64Var(p *uint64, name string, value uint64, usage string)
	UintVar(p *uint, name string, value uint, usage string)
}

var wrappedFlagSet FlagSet = flag.CommandLine

// SetFlagSet sets a flag set to be wrapped. flag.CommandLine is used by
// default.
func SetFlagSet(flagSet FlagSet) {
	wrappedFlagSet = flagSet
}

// Define defines a flag of type T with specified name, default value, and usage
// string. The return value is the address of a type T variable that stores th
// value of the flag.
func Define[T FlagConstraint](name string, defaultVal T, usage string) *T {
	var v any

	switch d := any(defaultVal).(type) {
	case bool:
		v = wrappedFlagSet.Bool(name, d, usage)
	case time.Duration:
		v = wrappedFlagSet.Duration(name, d, usage)
	case float64:
		v = wrappedFlagSet.Float64(name, d, usage)
	case int:
		v = wrappedFlagSet.Int(name, d, usage)
	case int64:
		v = wrappedFlagSet.Int64(name, d, usage)
	case string:
		v = wrappedFlagSet.String(name, d, usage)
	case uint:
		v = wrappedFlagSet.Uint(name, d, usage)
	case uint64:
		v = wrappedFlagSet.Uint64(name, d, usage)
	}

	return v.(*T)
}

// DefineVar defines a type T flag with specified name, default value, and usage
// string. The argument p points to a type T variable in which to store the
// value of the flag.
func DefineVar[T FlagConstraint](p *T, name string, defaultVal T, usage string) {
	switch d := any(defaultVal).(type) {
	case bool:
		wrappedFlagSet.BoolVar(any(p).(*bool), name, d, usage)
	case time.Duration:
		wrappedFlagSet.DurationVar(any(p).(*time.Duration), name, d, usage)
	case float64:
		wrappedFlagSet.Float64Var(any(p).(*float64), name, d, usage)
	case int:
		wrappedFlagSet.IntVar(any(p).(*int), name, d, usage)
	case int64:
		wrappedFlagSet.Int64Var(any(p).(*int64), name, d, usage)
	case string:
		wrappedFlagSet.StringVar(any(p).(*string), name, d, usage)
	case uint:
		wrappedFlagSet.UintVar(any(p).(*uint), name, d, usage)
	case uint64:
		wrappedFlagSet.Uint64Var(any(p).(*uint64), name, d, usage)
	}
}
