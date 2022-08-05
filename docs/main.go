// Copyright 2022 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"docs/attack"
	"docs/cmd"
	"embed"
	_ "embed"
	"os"
	"strconv"
)

//go:embed config/*
var configfs embed.FS

//go:embed template/title
var title string

//go:embed template/section_1
var sec1 string

//go:embed template/section_2
var sec2 string

//go:embed template/flags
var flagSec string

//go:embed template/section_detail
var secD string

//go:embed template/server_m
var serverM string

//go:embed template/server_sub
var serverSUB string

func main() {
	rootCmd := cmd.GetRootCMD()
	atkCmds := rootCmd.Commands()[0].Commands()
	clocks := atkCmds[4] //.Commands()[0]
	ac := attack.GetAttack(clocks)
	var buffer bytes.Buffer
	attack.ParseAtk(&buffer, "title", title, ac)
	if len(ac.SubCmd) > 0 {
		attack.ParseAtk(&buffer, "sec1", sec1, ac)
	} else {
		attack.ParseAtk(&buffer, "sec2", sec2, ac)
	}

	if len(ac.Flags) > 0 {
		attack.ParseAtk(&buffer, "flagSec", flagSec, ac)
	}
	for i, sub := range ac.SubCmd {
		secDt := secD
		attack.ParseSub(&buffer, "secD"+strconv.Itoa(i), secDt, sub)
		if len(sub.Flags) > 0 {
			flagSect := flagSec
			attack.ParseSub(&buffer, "flagSec"+strconv.Itoa(i), flagSect, sub)
		}
	}

	attack.ParseAtk(&buffer, "serverM", serverM, ac)

	for i, sub := range ac.SubCmd {
		serverSUBt := serverSUB
		attack.ParseSub(&buffer, "serverSub"+strconv.Itoa(i), serverSUBt, sub)
	}

	err := os.WriteFile("test.md", buffer.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
