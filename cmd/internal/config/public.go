//
//  Copyright (c) 2018, Joyent, Inc. All rights reserved.
//
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

package config

const (
	KeyAccount        = "general.account"
	KeyUrl            = "general.url"
	KeySshKeyMaterial = "general.key-material"
	KeySshKeyID       = "general.key-id"

	DefaultManDir = "./docs/man"
	ManSect       = 8

	DefaultMarkdownDir       = "./docs/md"
	DefaultMarkdownURLPrefix = "/command"

	KeyDocManDir            = "doc.mandir"
	KeyDocMarkdownDir       = "doc.markdown-dir"
	KeyDocMarkdownURLPrefix = "doc.markdown-url-prefix"

	KeyBashAutoCompletionTarget = "shell.autocomplete.bash.target"

	KeyUsePager = "general.use-pager"
	KeyUseUTC   = "general.utc"

	KeyLogFormat    = "log.format"
	KeyLogLevel     = "log.level"
	KeyLogStats     = "log.stats"
	KeyLogTermColor = "log.use-color"

	KeyInstanceName         = "compute.instance.name"
	KeyInstanceID           = "compute.instance.id"
	KeyInstanceWait         = "compute.instance.wait"
	KeyInstanceFirewall     = "compute.instance.firewall"
	KeyInstanceState        = "compute.instance.state"
	KeyInstanceBrand        = "compute.instance.brand"
	KeyInstanceNetwork      = "compute.instance.networks"
	KeyInstanceTag          = "compute.instance.tag"
	KeyInstanceSearchTag    = "compute.instance.search-tags"
	KeyInstanceMetadata     = "compute.instance.metadata"
	KeyInstanceAffinityRule = "compute.instance.affinity"
	KeyInstanceUserdata     = "compute.instance.userdata"
	KeyInstanceNamePrefix   = "compute.instance.name-prefix"

	KeyPackageName = "compute.package.name"
	KeyPackageId   = "compute.package.id"

	KeyImageName = "compute.image.name"
	KeyImageId   = "compute.image.id"
)
