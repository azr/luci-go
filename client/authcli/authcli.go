// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package authcli implements authentication related CLI subcommands and option
// parsing. Can be used from CLI tools that want customize authentication
// configuration from the command line.
package authcli

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/luci/luci-go/common/auth"
	"github.com/maruel/subcommands"
)

// Flags defines command line flags related to authentication.
type Flags struct {
	defaults           auth.Options
	serviceAccountJSON string
}

// Register adds auth related flags to a FlagSet.
func (fl *Flags) Register(f *flag.FlagSet, defaults auth.Options) {
	fl.defaults = defaults
	f.StringVar(&fl.serviceAccountJSON, "service-account-json", "", "Path to JSON file with service account credentials to use.")
}

// Options return instance of auth.Options struct with values set accordingly to
// parsed command line flags.
func (fl *Flags) Options() (auth.Options, error) {
	opts := fl.defaults
	if fl.serviceAccountJSON != "" {
		opts.Method = auth.ServiceAccountMethod
		opts.ServiceAccountJSONPath = fl.serviceAccountJSON
	}
	return opts, nil
}

// SubcommandLogin returns subcommands.Command that can be used to perform
// interactive login.
func SubcommandLogin(opts auth.Options, name string) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: name,
		ShortDesc: "performs interactive login flow",
		LongDesc:  "Performs interactive login flow and caches obtained credentials",
		CommandRun: func() subcommands.CommandRun {
			c := &loginRun{}
			c.flags.Register(&c.Flags, opts)
			return c
		},
	}
}

type loginRun struct {
	subcommands.CommandRunBase
	flags Flags
}

func (c *loginRun) Run(subcommands.Application, []string) int {
	opts, err := c.flags.Options()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	client, err := auth.AuthenticatedClient(auth.InteractiveLogin, auth.NewAuthenticator(opts))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Login failed: %s\n", err.Error())
		return 2
	}
	err = reportIdentity(client)
	if err != nil {
		return 3
	}
	return 0
}

// SubcommandLogout returns subcommands.Command that can be used to purge cached
// credentials.
func SubcommandLogout(opts auth.Options, name string) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: name,
		ShortDesc: "removes cached credentials",
		LongDesc:  "Removes cached credentials from the disk",
		CommandRun: func() subcommands.CommandRun {
			c := &logoutRun{}
			c.flags.Register(&c.Flags, opts)
			return c
		},
	}
}

type logoutRun struct {
	subcommands.CommandRunBase
	flags Flags
}

func (c *logoutRun) Run(a subcommands.Application, args []string) int {
	opts, err := c.flags.Options()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	err = auth.NewAuthenticator(opts).PurgeCredentialsCache()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	return 0
}

// SubcommandInfo returns subcommand.Command that can be used to print current
// cached credentials.
func SubcommandInfo(opts auth.Options, name string) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: name,
		ShortDesc: "prints an email address associated with currently cached token",
		LongDesc:  "Prints an email address associated with currently cached token",
		CommandRun: func() subcommands.CommandRun {
			c := &infoRun{}
			c.flags.Register(&c.Flags, opts)
			return c
		},
	}
}

type infoRun struct {
	subcommands.CommandRunBase
	flags Flags
}

func (c *infoRun) Run(a subcommands.Application, args []string) int {
	opts, err := c.flags.Options()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	client, err := auth.AuthenticatedClient(auth.SilentLogin, auth.NewAuthenticator(opts))
	if err == auth.ErrLoginRequired {
		fmt.Fprintln(os.Stderr, "Not logged in")
		return 2
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 3
	}
	err = reportIdentity(client)
	if err != nil {
		return 4
	}
	return 0
}

// SubcommandToken returns subcommand.Command that can be used to print current
// access token.
func SubcommandToken(opts auth.Options, name string) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: name,
		ShortDesc: "prints an access token",
		LongDesc:  "Generates an access token if requested and prints it.",
		CommandRun: func() subcommands.CommandRun {
			c := &tokenRun{}
			c.flags.Register(&c.Flags, opts)
			c.Flags.DurationVar(
				&c.lifetime, "lifetime", time.Minute,
				"Minimum token lifetime. If existing token expired and refresh token or service account is not present, returns nothing.",
			)
			c.Flags.StringVar(
				&c.jsonOutput, "json-output", "",
				"Destination file to print token and expiration time in JSON. \"-\" for standard output.")
			return c
		},
	}
}

type tokenRun struct {
	subcommands.CommandRunBase
	flags      Flags
	lifetime   time.Duration
	jsonOutput string
}

const (
	TokenExitCodeValidToken = iota
	TokenExitCodeNoValidToken
	TokenExitCodeInvalidInput
	TokenExitCodeInternalError
)

func (c *tokenRun) Run(a subcommands.Application, args []string) int {
	opts, err := c.flags.Options()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return TokenExitCodeInvalidInput
	}
	if c.lifetime > 45*time.Minute {
		fmt.Fprintln(os.Stderr, "lifetime cannot exceed 45m")
		return TokenExitCodeInvalidInput
	}

	authenticator := auth.NewAuthenticator(opts)
	token, expiry, err := authenticator.GetAccessToken(c.lifetime)
	if err != nil {
		if err == auth.ErrLoginRequired {
			fmt.Fprintln(os.Stderr, "interactive login required")
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		return TokenExitCodeNoValidToken
	}
	if token == "" {
		return TokenExitCodeNoValidToken
	}

	if c.jsonOutput == "" {
		fmt.Println(token)
	} else {
		out := os.Stdout
		if c.jsonOutput != "-" {
			out, err = os.Create(c.jsonOutput)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return TokenExitCodeInvalidInput
			}
			defer out.Close()
		}
		data := struct {
			Token  string `json:"token"`
			Expiry int64  `json:"expiry"`
		}{token, expiry.Unix()}
		if err = json.NewEncoder(out).Encode(data); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return TokenExitCodeInternalError
		}
	}
	return TokenExitCodeValidToken
}

// reportIdentity prints identity associated with credentials that the client
// puts into each request (if any).
func reportIdentity(c *http.Client) error {
	service := auth.NewGroupsService("", c, nil)
	ident, err := service.FetchCallerIdentity()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch current identity: %s\n", err)
		return err
	}
	fmt.Printf("Logged in to %s as %s\n", service.ServiceURL(), ident)
	return nil
}