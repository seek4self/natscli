// Copyright 2020 The STHG Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/choria-io/fisk"
	"github.com/nats-io/natscli/cli"
)

var version = "development"

func main() {
	help := `STHG-MS Utility

STHG-MS and JetStream administration.

See 'ms-client cheat' for a quick cheatsheet of commands`

	ncli := fisk.New("ms-client", help)
	ncli.Author("STHG-MS Authors")
	ncli.UsageWriter(os.Stdout)
	ncli.UsageTemplate(fisk.CompactMainUsageTemplate)
	ncli.ErrorUsageTemplate(fisk.CompactMainUsageTemplate)
	ncli.Version(getVersion())
	ncli.HelpFlag.Short('h')
	ncli.WithCheats().CheatCommand.Hidden()

	opts, err := cli.ConfigureInApp(ncli, nil, true)
	if err != nil {
		return
	}
	cli.SetVersion(version)

	ncli.Flag("server", "STHG-MS server urls").Short('s').Envar("STHG_URL").PlaceHolder("URL").StringVar(&opts.Servers)
	ncli.Flag("user", "Username or Token").Envar("STHG_USER").PlaceHolder("USER").StringVar(&opts.Username)
	ncli.Flag("password", "Password").Envar("STHG_PASSWORD").PlaceHolder("PASSWORD").StringVar(&opts.Password)
	ncli.Flag("connection-name", "Nickname to use for the underlying STHG Connection").Default("STHG CLI Version " + version).PlaceHolder("NAME").StringVar(&opts.ConnectionName)
	ncli.Flag("creds", "User credentials").Envar("STHG_CREDS").PlaceHolder("FILE").StringVar(&opts.Creds)
	ncli.Flag("nkey", "User NKEY").Envar("STHG_NKEY").PlaceHolder("FILE").StringVar(&opts.Nkey)
	ncli.Flag("tlscert", "TLS public certificate").Envar("STHG_CERT").PlaceHolder("FILE").ExistingFileVar(&opts.TlsCert)
	ncli.Flag("tlskey", "TLS private key").Envar("STHG_KEY").PlaceHolder("FILE").ExistingFileVar(&opts.TlsKey)
	ncli.Flag("tlsca", "TLS certificate authority chain").Envar("STHG_CA").PlaceHolder("FILE").ExistingFileVar(&opts.TlsCA)
	ncli.Flag("timeout", "Time to wait on responses from STHG-MS").Default("5s").Envar("STHG_TIMEOUT").PlaceHolder("DURATION").DurationVar(&opts.Timeout)
	ncli.Flag("js-api-prefix", "Subject prefix for access to JetStream API").PlaceHolder("PREFIX").StringVar(&opts.JsApiPrefix)
	ncli.Flag("js-event-prefix", "Subject prefix for access to JetStream Advisories").PlaceHolder("PREFIX").StringVar(&opts.JsEventPrefix)
	ncli.Flag("js-domain", "JetStream domain to access").PlaceHolder("DOMAIN").StringVar(&opts.JsDomain)
	ncli.Flag("inbox-prefix", "Custom inbox prefix to use for inboxes").PlaceHolder("PREFIX").StringVar(&opts.InboxPrefix)
	ncli.Flag("domain", "JetStream domain to access").PlaceHolder("DOMAIN").Hidden().StringVar(&opts.JsDomain)
	ncli.Flag("context", "Configuration context").Envar("STHG_CONTEXT").PlaceHolder("NAME").StringVar(&opts.CfgCtx)
	ncli.Flag("trace", "Trace API interactions").UnNegatableBoolVar(&opts.Trace)

	log.SetFlags(log.Ltime)

	ncli.MustParseWithUsage(os.Args[1:])
}

func getVersion() string {
	if version != "development" {
		return version
	}

	nfo, ok := debug.ReadBuildInfo()
	if !ok || (nfo != nil && nfo.Main.Version == "") {
		return version
	}

	return nfo.Main.Version
}
