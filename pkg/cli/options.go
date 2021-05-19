package cli

import (
	"fmt"
	"time"

	"github.com/openshift/assisted-service/internal/assistedserviceiso"
	"github.com/openshift/assisted-service/internal/bminventory"
	"github.com/openshift/assisted-service/internal/cluster"
	"github.com/openshift/assisted-service/internal/cluster/validations"
	"github.com/openshift/assisted-service/internal/controller/controllers"
	"github.com/openshift/assisted-service/internal/garbagecollector"
	"github.com/openshift/assisted-service/internal/hardware"
	"github.com/openshift/assisted-service/internal/host"
	"github.com/openshift/assisted-service/internal/host/hostcommands"
	"github.com/openshift/assisted-service/internal/isoeditor"
	"github.com/openshift/assisted-service/internal/network"
	"github.com/openshift/assisted-service/internal/operators"
	"github.com/openshift/assisted-service/internal/versions"
	"github.com/openshift/assisted-service/pkg/auth"
	dbPkg "github.com/openshift/assisted-service/pkg/db"
	"github.com/openshift/assisted-service/pkg/generator"
	"github.com/openshift/assisted-service/pkg/leader"
	logconfig "github.com/openshift/assisted-service/pkg/log"
	"github.com/openshift/assisted-service/pkg/ocm"
	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
)

var Options struct {
	Auth                        auth.Config
	BMConfig                    bminventory.Config
	DBConfig                    dbPkg.Config
	HWValidatorConfig           hardware.ValidatorCfg
	GeneratorConfig             generator.Config
	InstructionConfig           hostcommands.InstructionConfig
	OperatorsConfig             operators.Options
	GCConfig                    garbagecollector.Config
	ClusterStateMonitorInterval time.Duration `envconfig:"CLUSTER_MONITOR_INTERVAL" default:"10s"`
	S3Config                    s3wrapper.Config
	HostStateMonitorInterval    time.Duration `envconfig:"HOST_MONITOR_INTERVAL" default:"8s"`
	Versions                    versions.Versions
	OpenshiftVersions           string        `envconfig:"OPENSHIFT_VERSIONS"`
	ReleaseImageMirror          string        `envconfig:"OPENSHIFT_INSTALL_RELEASE_IMAGE_MIRROR" default:""`
	CreateS3Bucket              bool          `envconfig:"CREATE_S3_BUCKET" default:"false"`
	ImageExpirationInterval     time.Duration `envconfig:"IMAGE_EXPIRATION_INTERVAL" default:"30m"`
	ClusterConfig               cluster.Config
	DeployTarget                string `envconfig:"DEPLOY_TARGET" default:"k8s"`
	Storage                     string `envconfig:"STORAGE" default:"s3"`
	OCMConfig                   ocm.Config
	HostConfig                  host.Config
	LogConfig                   logconfig.Config
	LeaderConfig                leader.Config
	ValidationsConfig           validations.Config
	AssistedServiceISOConfig    assistedserviceiso.Config
	manifestsGeneratorConfig    network.Config
	EnableKubeAPI               bool `envconfig:"ENABLE_KUBE_API" default:"false"`
	InfraEnvConfig              controllers.InfraEnvConfig
	ISOEditorConfig             isoeditor.Config
	CheckClusterVersion         bool          `envconfig:"CHECK_CLUSTER_VERSION" default:"false"`
	DeletionWorkerInterval      time.Duration `envconfig:"DELETION_WORKER_INTERVAL" default:"1h"`
	DeregisterWorkerInterval    time.Duration `envconfig:"DEREGISTER_WORKER_INTERVAL" default:"1h"`
	EnableDeletedUnregisteredGC bool          `envconfig:"ENABLE_DELETE_UNREGISTER_GC" default:"true"`
	EnableDeregisterInactiveGC  bool          `envconfig:"ENABLE_DEREGISTER_INACTIVE_GC" default:"true"`
	ServeHTTPS                  bool          `envconfig:"SERVE_HTTPS" default:"false"`
	HTTPSKeyFile                string        `envconfig:"HTTPS_KEY_FILE" default:""`
	HTTPSCertFile               string        `envconfig:"HTTPS_CERT_FILE" default:""`
	FileSystemUsageThreshold    int           `envconfig:"FILESYSTEM_USAGE_THRESHOLD" default:"80"`
	EnableElasticAPM            bool          `envconfig:"ENABLE_ELASTIC_APM" default:"false"`
	WorkDir                     string        `envconfig:"WORK_DIR" default:"/data/"`
}

func InitLogs() *logrus.Entry {
	log := logrus.New()

	fmt.Println(Options.EnableElasticAPM)
	if Options.EnableElasticAPM {
		log.AddHook(&apmlogrus.Hook{})
	}

	log.SetReportCaller(true)

	logger := log.WithFields(logrus.Fields{})

	//set log format according to configuration
	logger.Info("Setting log format: ", Options.LogConfig.LogFormat)
	if Options.LogConfig.LogFormat == logconfig.LogFormatJson {
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	//set log level according to configuration
	logger.Info("Setting Log Level: ", Options.LogConfig.LogLevel)
	logLevel, err := logrus.ParseLevel(Options.LogConfig.LogLevel)
	if err != nil {
		logger.Error("Invalid Log Level: ", Options.LogConfig.LogLevel)
	} else {
		log.SetLevel(logLevel)
	}

	return logger
}
