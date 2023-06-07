// Copyright 2020 The Kube-burner Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"time"

	"github.com/cloud-bulldozer/go-commons/indexers"
	"github.com/cloud-bulldozer/kube-burner/pkg/alerting"
	"github.com/cloud-bulldozer/kube-burner/pkg/config"
	"github.com/cloud-bulldozer/kube-burner/pkg/prometheus"
)

// ScraperConfig holds data related to scraper and target indexer
type ScraperConfig struct {
	ConfigSpec      config.Spec
	Password        string
	PrometheusStep  time.Duration
	MetricsEndpoint string
	MetricsProfile  string
	AlertProfile    string
	SkipTLSVerify   bool
	URL             string
	Token           string
	Username        string
	StartTime       int64
	EndTime         int64
	JobName         string
	ActionIndex     bool
	UserMetaData    string
}

// ScraperResponse holds parsed data related to scraper and target indexer
type Scraper struct {
	PrometheusClients   []*prometheus.Prometheus
	AlertMs             []*alerting.AlertManager
	Indexer             *indexers.Indexer
	UserMetadataContent map[string]interface{}
}
