/*
 *  Copyright 2020 ChronoWave Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Package parser declares an expression parser with support for macro
 *  expansion.
 */

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/chronowave/chronowave/ssql/parser"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"

	"github.com/chronowave/chronowave/embed"
)

func main() {
	hfmi.SetSegmentCache(5 * 1024 * 1024)

	var rootCmd = &cobra.Command{Use: "waverider"}
	rootCmd.AddCommand(indexCommand(), queryCommand())
	rootCmd.Execute()
}

func indexCommand() *cobra.Command {
	var timestamp string
	var keys []string
	cmd := &cobra.Command{
		Use:   "index [path to json file]",
		Short: "Build index in -d {dir}",
		Long:  `Build semi-structured data index. Only supports UTF-8 encoded JSON`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			close, err := embed.Verify()
			if err != nil {
				panic(err)
			}
			defer close()

			if len(timestamp) == 0 {
				printToScreen("Error: timestamp field must not be empty", true)
				os.Exit(1)
			}

			for _, f := range args {
				info, err := os.Stat(f)
				if err != nil {
					fmt.Printf("skip %v due to err: %v\n", f, err)
					continue
				}

				s := time.Now()
				if err := embed.Build(f, timestamp, keys); err != nil {
					fmt.Printf("indexing %v err: %v\n", f, err)
				} else {
					fmt.Printf("indexed %v with %d Bytes in %v\n", f, info.Size(), time.Since(s))
				}
			}
		},
	}

	cmd.Flags().StringVarP(&embed.Directory, "dir", "d", "data", "index directory")
	cmd.Flags().StringVarP(&timestamp, "timestamp", "t", "", "JSON path to timestamp field, example '/timestamp'")
	cmd.Flags().StringSliceVarP(&keys, "keys", "k", nil, "JSON path to key field in JSON can be queried by key, w/o time range")

	return cmd
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [SSQL query]",
		Short: "Execute SSQL",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			close, err := embed.Verify()
			if err != nil {
				panic(err)
			}
			defer close()

			stmt, errs := parser.Parse(args[0])
			if len(errs) > 0 {
				var anchor [][]byte
				var msg []string
				for _, e := range errs {
					if e.Line > len(anchor) {
						anchor = append(anchor, nil)
					}
					i := e.Line - 1
					sz := e.Column - len(anchor[i])
					if sz > 0 {
						anchor[i] = append(anchor[i], make([]byte, sz)...)
					}

					anchor[i][e.Column-1] = '^'

					msg = append(msg, e.Message)
				}
				scanner := bufio.NewScanner(strings.NewReader(args[0]))
				for i := 0; scanner.Scan(); i++ {
					printToScreen(scanner.Text(), false)
					if i < len(anchor) && len(anchor[i]) > 0 {
						for j, b := range anchor[i] {
							if b == 0 {
								anchor[i][j] = ' '
							}
						}
						printToScreen(string(anchor[i]), true)
					}
				}

				for _, m := range msg {
					printToScreen(m, true)
				}
				return
			}

			answer := embed.Query(context.Background(), stmt)
			fmt.Println(string(answer))
		},
	}

	cmd.Flags().StringVarP(&embed.Directory, "dir", "d", "data", "index directory")

	return cmd
}

func printToScreen(msg string, red bool) {
	if !terminal.IsTerminal(0) || !terminal.IsTerminal(1) {
		fmt.Printf(msg)
		return
	}

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		fmt.Printf(msg)
		return
	}
	defer terminal.Restore(0, oldState)

	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term := terminal.NewTerminal(screen, "")
	escape := ""
	if red {
		escape = string(term.Escape.Red)
	}

	fmt.Fprintln(term, escape, msg, escape)
}
