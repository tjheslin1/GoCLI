// Copyright © 2017 Thomas Heslin <tjheslin1@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates the version command
func NewVersionCmd(version string) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Displays the version of this CLI tool",
		Long:  `Displays the version of this CLI tool`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CLI tool's version is ", version)
		},
	}

	// versionCmd.Flags()...

	return versionCmd
}
