// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"

	"github.com/kasworld/argdefault"
)

type Config struct {
	BaseLogDir            string  `argname:""`
	DataFolder            string  `default:"./serverdata"`
	ClientDataFolder      string  `default:"./clientdata" argname:""`
	GroundRPC             string  `default:"localhost:14002"  argname:""`
	ServicePort           string  `default:":14101"  argname:"port"`
	AdminPort             string  `default:":14201"  argname:""`
	TowerFilename         string  `default:"starting" argname:""`
	TowerNumber           int     `default:"1" argname:""`
	DisplayName           string  `default:"Default" argname:""`
	ConcurrentConnections int     `default:"10000" argname:""`
	ActTurnPerSec         float64 `default:"2.0" argname:""`
	StandAlone            bool    `default:"true" argname:""`
}

func main() {
	config := &Config{
		TowerNumber: 43,
	}
	argdefault.SetZeroFieldToDefault(config)
	c := argdefault.AddArgsWith(config)
	flag.Parse()
	c.ApplyArgsTo(config)
	// fmt.Println(prettystring.PrettyString(config, 4))
	fmt.Printf("%#v\n", config)
}
