/*
Copyright 2020 The KubeSphere Authors.

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

package manager

import (
	kubekeyapiv1alpha1 "github.com/kubesphere/kubekey/apis/kubekey/v1alpha1"
	kubekeyclientset "github.com/kubesphere/kubekey/clients/clientset/versioned"
	"github.com/kubesphere/kubekey/pkg/connector"
	"github.com/kubesphere/kubekey/pkg/util/runner"
	log "github.com/sirupsen/logrus"
	"sync"
)

const (
	IsInitCluster = true
	Docker        = "docker"
	Conatinerd    = "containerd"
	Crio          = "crio"
	Isula         = "isula"
)

// Manager defines all the parameters needed for the installation.
type Manager struct {
	ObjName                  string
	Cluster                  *kubekeyapiv1alpha1.ClusterSpec
	Logger                   log.FieldLogger
	Connector                connector.Connector
	Runner                   *runner.Runner
	AllNodes                 []kubekeyapiv1alpha1.HostCfg
	EtcdNodes                []kubekeyapiv1alpha1.HostCfg
	MasterNodes              []kubekeyapiv1alpha1.HostCfg
	WorkerNodes              []kubekeyapiv1alpha1.HostCfg
	K8sNodes                 []kubekeyapiv1alpha1.HostCfg
	EtcdContainer            bool
	ClusterHosts             []string
	WorkDir                  string
	KsEnable                 bool
	KsVersion                string
	Debug                    bool
	SkipCheck                bool
	SkipPullImages           bool
	SourcesDir               string
	AddImagesRepo            bool
	InCluster                bool
	ContainerManager         string
	ContainerRuntimeEndpoint string
	DeployLocalStorage       bool
	Kubeconfig               string
	ClusterStatus            *ClusterStatus
	UpgradeStatus            *UpgradeStatus
	Conditions               []kubekeyapiv1alpha1.Condition
	ClientSet                *kubekeyclientset.Clientset
	DownloadCommand          func(path, url string) string
}

// ClusterStatus is used to store cluster status
type ClusterStatus struct {
	IsExist        bool
	Version        string
	AllNodesInfo   map[string]string
	Kubeconfig     string
	BootstrapToken string
	CertificateKey string
}

// UpgradeStatus is used to store cluster upgrade status
type UpgradeStatus struct {
	CurrentVersions   map[string]string
	CurrentVersionStr string
	NextVersionStr    string
	MU                *sync.Mutex
	Kubeconfig        string
}

// Copy is used to create a copy for Manager.
func (mgr *Manager) Copy() *Manager {
	newManager := *mgr
	return &newManager
}

// ExistNode is used determine if the node already exists.
func ExistNode(mgr *Manager, node *kubekeyapiv1alpha1.HostCfg) bool {
	var version bool
	_, name := mgr.ClusterStatus.AllNodesInfo[node.Name]
	if name && mgr.ClusterStatus.AllNodesInfo[node.Name] != "" {
		version = true
	}
	_, ip := mgr.ClusterStatus.AllNodesInfo[node.InternalAddress]
	return version || ip
}
