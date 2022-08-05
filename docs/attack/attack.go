package attack

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
	"os"
	"text/template"
)

type AtkConfig struct {
	TypeTitle  string         `yaml:"type_title"`
	TypePlain  string         `yaml:"type_plain"`
	WorkBaseOn string         `yaml:"work_base_on"`
	FlagDspMap map[string]Arg `yaml:"flag_dsp_map"`
}

func GetAttackConfig(name string) AtkConfig {
	configb, err := os.ReadFile("config/" + name + ".yaml")
	if err != nil {
		panic(err)
	}
	var c AtkConfig
	err = yaml.Unmarshal(configb, &c)
	if err != nil {
		panic(err)
	}
	return c
}

type Attack struct {
	AtkConfig
	Name        string
	HelpCommand string
	Usage       string
	Path        string
	SubCmd      []SubCMD
	Flags       []Flag
}

func GetAttack(cmd *cobra.Command) Attack {
	AtkGen(cmd)
	c := GetAttackConfig(cmd.Name())

	mainHelpCmd := fmt.Sprintf("chaosd attack %s help", cmd.Name())
	mainUsage := fmt.Sprintf("%s \n\n %s", cmd.Short, cmd.UsageString())

	return Attack{
		AtkConfig:   c,
		Name:        cmd.Name(),
		HelpCommand: mainHelpCmd,
		Usage:       mainUsage,
		Path:        fmt.Sprintf("/api/attack/%s", cmd.Name()),
		SubCmd:      GetSubCMD(cmd),
		Flags:       insertFlagDspMap(cmd, c.FlagDspMap),
	}
}

type SubConfig struct {
	TypeTitle  string         `yaml:"type_title"`
	TypePlain  string         `yaml:"type_plain"`
	FlagDspMap map[string]Arg `yaml:"flag_dsp_map"`
}

type Arg struct {
	Dsp  string `yaml:"dsp"`
	Must bool   `yaml:"must"`
}

func GetSubConfig(mName string, subName string) SubConfig {
	configb, err := os.ReadFile(fmt.Sprintf("config/%s.%s.yaml", mName, subName))
	if err != nil {
		panic(err)
	}
	var c SubConfig
	err = yaml.Unmarshal(configb, &c)
	if err != nil {
		panic(err)
	}
	return c
}

type SubCMD struct {
	SubConfig
	Name        string
	HelpCommand string
	Usage       string
	Flags       []Flag
}

func GetSubCMD(mCmd *cobra.Command) []SubCMD {
	subCMDs := mCmd.Commands()
	var subs []SubCMD
	for _, cmd := range subCMDs {
		SubGen(mCmd, cmd)
		c := GetSubConfig(mCmd.Name(), cmd.Name())
		helpCmd := fmt.Sprintf("chaosd attack %s %s help", mCmd.Name(), cmd.Name())
		usage := fmt.Sprintf("%s \n\n %s", cmd.Short, cmd.UsageString())
		subs = append(subs, SubCMD{
			SubConfig:   c,
			Name:        cmd.Name(),
			HelpCommand: helpCmd,
			Usage:       usage,
			Flags:       insertFlagDspMap(cmd, c.FlagDspMap),
		})
	}
	return subs
}

func ParseAtk(buffer *bytes.Buffer, name string, tmpS string, attack Attack) {
	tmp, err := template.New(name).Parse(tmpS)
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(buffer, attack)
	if err != nil {
		panic(err)
	}
}

func ParseSub(buffer *bytes.Buffer, name string, tmpS string, attack SubCMD) {
	tmp, err := template.New(name).Parse(tmpS)
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(buffer, attack)
	if err != nil {
		panic(err)
	}
}

func AtkGen(cmd *cobra.Command) {
	fileName := fmt.Sprintf("config/%s.yaml", cmd.Name())
	var c AtkConfig
	if buffer, err := os.ReadFile(fileName); err == nil {
		err := yaml.Unmarshal(buffer, &c)
		if err != nil {
			panic(err)
		}
	}
	if c.FlagDspMap == nil {
		c.FlagDspMap = make(map[string]Arg)
	}
	_ = insertFlagDspMap(cmd, c.FlagDspMap)
	buffer, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fileName, buffer, 0777)
	if err != nil {
		panic(err)
	}
}

func SubGen(cmd *cobra.Command, subCmd *cobra.Command) {
	fileName := fmt.Sprintf("config/%s.%s.yaml", cmd.Name(), subCmd.Name())
	var c SubConfig
	if buffer, err := os.ReadFile(fileName); err == nil {
		err := yaml.Unmarshal(buffer, &c)
		if err != nil {
			panic(err)
		}
	} else {
	}
	if c.FlagDspMap == nil {
		c.FlagDspMap = make(map[string]Arg)
	}
	_ = insertFlagDspMap(subCmd, c.FlagDspMap)
	buffer, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fileName, buffer, 0777)
	if err != nil {
		panic(err)
	}
}

func insertFlagDspMap(cmd *cobra.Command, dspMap map[string]Arg) []Flag {
	flags := GetFlags(cmd)
	for i := range flags {
		if arg, ok := dspMap[flags[i].Name]; ok {
			flags[i].Description = arg.Dsp
			flags[i].Must = arg.Must
		} else {
			dspMap[flags[i].Name] = Arg{Dsp: flags[i].Description}
		}
		if flags[i].Default == "" {
			flags[i].Default = "\"\""
		}
	}
	return flags
}

type Flag struct {
	Name        string
	Short       string
	Type        string
	Default     string
	Description string
	Must        bool
}

func GetFlags(cmd *cobra.Command) []Flag {
	flagSet := cmd.LocalFlags()
	var flags []pflag.Flag
	flagSet.VisitAll(func(f *pflag.Flag) {
		flags = append(flags, *f)
	})
	var Flags []Flag
	for _, f := range flags {
		Flags = append(Flags, Flag{
			Name:        f.Name,
			Short:       f.Shorthand,
			Type:        f.Value.Type(),
			Default:     f.DefValue,
			Description: f.Usage,
		})
	}
	return Flags
}
