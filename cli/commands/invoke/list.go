/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package invoke

import (
	"flag"

	"errors"
	"fmt"
	"github.com/verizonlabs/northstar/cli/commands"
	"github.com/verizonlabs/northstar/cli/util"
	"github.com/verizonlabs/northstar/data/invocations/client"
)

type ListInvocationsCmd struct {
	client    *client.InvocationClient
	cmd       *flag.FlagSet
	snippetId *string
	limit     *int
}

func NewListInvocation(client *client.InvocationClient) commands.Command {
	cmd := flag.NewFlagSet("invoke-list", flag.ExitOnError)
	snippetId := cmd.String("snippetId", "", "The snippet id")
	limit := cmd.Int("limit", 100, "The limit")

	return &ListInvocationsCmd{client: client,
		cmd:       cmd,
		snippetId: snippetId,
		limit:     limit}
}

func (output *ListInvocationsCmd) Run(args []string) error {
	output.cmd.Parse(args)

	if !output.cmd.Parsed() {
		return errors.New("Failed to parse cmd")
	}

	if *output.snippetId == "" {
		return errors.New("Please set a snippetId using -snippetId.")
	}

	results, err := output.client.GetInvocationResults(util.GetAccountID(),
		*output.snippetId,
		*output.limit)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Println("No invocations found")
		return nil
	}

	for _, result := range results {
		fmt.Println(result.Print())
	}
	return nil
}
