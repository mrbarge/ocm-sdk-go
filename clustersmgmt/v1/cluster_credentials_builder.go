/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterCredentialsBuilder contains the data and logic needed to build 'cluster_credentials' objects.
//
// Credentials of the a cluster.
type ClusterCredentialsBuilder struct {
	id         *string
	href       *string
	link       bool
	ssh        *SSHCredentialsBuilder
	admin      *AdminCredentialsBuilder
	kubeconfig *string
}

// NewClusterCredentials creates a new builder of 'cluster_credentials' objects.
func NewClusterCredentials() *ClusterCredentialsBuilder {
	return new(ClusterCredentialsBuilder)
}

// ID sets the identifier of the object.
func (b *ClusterCredentialsBuilder) ID(value string) *ClusterCredentialsBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *ClusterCredentialsBuilder) HREF(value string) *ClusterCredentialsBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *ClusterCredentialsBuilder) Link(value bool) *ClusterCredentialsBuilder {
	b.link = value
	return b
}

// SSH sets the value of the 'SSH' attribute to the given value.
//
// SSH key pair of a cluster.
func (b *ClusterCredentialsBuilder) SSH(value *SSHCredentialsBuilder) *ClusterCredentialsBuilder {
	b.ssh = value
	return b
}

// Admin sets the value of the 'admin' attribute to the given value.
//
// Temporary administrator credentials generated during the installation of the
// cluster.
func (b *ClusterCredentialsBuilder) Admin(value *AdminCredentialsBuilder) *ClusterCredentialsBuilder {
	b.admin = value
	return b
}

// Kubeconfig sets the value of the 'kubeconfig' attribute to the given value.
//
//
func (b *ClusterCredentialsBuilder) Kubeconfig(value string) *ClusterCredentialsBuilder {
	b.kubeconfig = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterCredentialsBuilder) Copy(object *ClusterCredentials) *ClusterCredentialsBuilder {
	if object == nil {
		return b
	}
	b.id = object.id
	b.href = object.href
	b.link = object.link
	if object.ssh != nil {
		b.ssh = NewSSHCredentials().Copy(object.ssh)
	} else {
		b.ssh = nil
	}
	if object.admin != nil {
		b.admin = NewAdminCredentials().Copy(object.admin)
	} else {
		b.admin = nil
	}
	b.kubeconfig = object.kubeconfig
	return b
}

// Build creates a 'cluster_credentials' object using the configuration stored in the builder.
func (b *ClusterCredentialsBuilder) Build() (object *ClusterCredentials, err error) {
	object = new(ClusterCredentials)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.ssh != nil {
		object.ssh, err = b.ssh.Build()
		if err != nil {
			return
		}
	}
	if b.admin != nil {
		object.admin, err = b.admin.Build()
		if err != nil {
			return
		}
	}
	object.kubeconfig = b.kubeconfig
	return
}
