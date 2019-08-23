// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-08-22 17:13:46.911907902 -0700 PDT m=+6.167961714

package events

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "type"), *new(string), "Sets the type of EventSink to configure [log/admin/file].")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "file-path"), *new(string), "For file types,  specify where the file should be located.")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "rate"), *new(int64), "Max rate at which events can be recorded per second.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "capacity"), *new(int), "The max bucket size for event recording tokens.")
	return cmdFlags
}
