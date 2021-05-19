package main

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	bmh_v1alpha1 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
	"github.com/openshift/assisted-service/internal/common"
	hiveext "github.com/openshift/assisted-service/internal/controller/api/hiveextension/v1beta1"
	aiv1beta1 "github.com/openshift/assisted-service/internal/controller/api/v1beta1"
	aictrl "github.com/openshift/assisted-service/internal/controller/controllers"
	"github.com/openshift/assisted-service/pkg/cli"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var log *logrus.Entry

const (
	deployment_type_k8s    = "k8s"
	deployment_type_onprem = "onprem"
	deployment_type_ocp    = "ocp"
	storage_filesystem     = "filesystem"
	storage_s3             = "s3"
)

func failOnError(err error, msg string, args ...interface{}) {
	if err != nil {
		log.WithError(err).Fatalf(msg, args...)
	}
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	var ctrlClient ctrlclient.Client
	var schemes = runtime.NewScheme()
	if cli.Options.EnableKubeAPI {
		scheme.AddToScheme(schemes)
		aiv1beta1.AddToScheme(schemes)
		hivev1.AddToScheme(schemes)
		hiveext.AddToScheme(schemes)
		bmh_v1alpha1.AddToScheme(schemes)

		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		// if you want to change the loading rules (which files in which order), you can do so here

		configOverrides := &clientcmd.ConfigOverrides{}
		// if you want to change override values or bind them to flags, there are methods to help you

		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
		cfg, err := kubeConfig.ClientConfig()
		failOnError(err, "Failed to create kubernetes cluster config")

		ctrlClient, err = ctrlclient.New(cfg, ctrlclient.Options{Scheme: schemes})
		failOnError(err, "Failed to create kubernetes cluster config")
	}

	bmhList := []types.NamespacedName{}

	if len(args) > 0 {
		for _, bmhName := range args {
			bmhList = append(bmhList, types.NamespacedName{
				Namespace: "assisted-installer",
				Name:      bmhName,
			})
		}
	} else {
		ctx := context.Background()
		bmhCRs := bmh_v1alpha1.BareMetalHostList{}
		err := ctrlClient.List(ctx, &bmhCRs)

		if err != nil {
			failOnError(err, "Couldn't get list of BMH resources")
		}

		for _, bmh := range bmhCRs.Items {
			bmhList = append(bmhList, types.NamespacedName{
				Namespace: bmh.Namespace,
				Name:      bmh.Name,
			})
		}

	}

	for _, bmh := range bmhList {
		debugBMH(ctrlClient, bmh)
	}

}

func debugBMH(ctrlClient ctrlclient.Client, nsName types.NamespacedName) {
	var err error
	ctx := context.Background()
	bmh := &bmh_v1alpha1.BareMetalHost{}

	if err = ctrlClient.Get(ctx, nsName, bmh); err != nil {
		failOnError(err, "Failed to get BareMetalHost for %v", nsName)
	}

	log.Infof("Found BMH: %s -> %s", bmh.Namespace, bmh.Name)

	if _, ok := bmh.Labels[aictrl.BMH_INFRA_ENV_LABEL]; !ok {
		log.Fatalf("The BMH doesn't have the %s label. Set it to point to the infraenv", aictrl.BMH_INFRA_ENV_LABEL)
	}

	infraEnv := &aiv1beta1.InfraEnv{}
	depRes := types.NamespacedName{
		Namespace: bmh.Namespace,
		Name:      bmh.Labels[aictrl.BMH_INFRA_ENV_LABEL],
	}

	if err = ctrlClient.Get(ctx, depRes, infraEnv); err != nil {
		failOnError(err, "Failed to get the InfraEnv using the following namespace and name %v", depRes)
	}

	log.Infof("Found InfraEnv: %s -> %s", infraEnv.Namespace, infraEnv.Name)

	clusterDeployment := &hivev1.ClusterDeployment{}
	depRes = types.NamespacedName{
		Namespace: infraEnv.Spec.ClusterRef.Namespace,
		Name:      infraEnv.Spec.ClusterRef.Name,
	}

	if err = ctrlClient.Get(ctx, depRes, clusterDeployment); err != nil {
		failOnError(err, "Failed to get ClusterDeployment using the following namespace and name %v", depRes)
	}

	log.Infof("Found ClusterDeployment: %s -> %s", clusterDeployment.Namespace, clusterDeployment.Name)

	if clusterDeployment.Spec.ClusterInstallRef == nil ||
		clusterDeployment.Spec.ClusterInstallRef.Group != hiveext.Group ||
		clusterDeployment.Spec.ClusterInstallRef.Kind != "AgentClusterInstall" {
		log.Fatalf("Make sure the ClusterDeployment has a valid ClusterInstall reference of kind AgentClusterInstall %v", clusterDeployment.Spec.ClusterInstallRef)
	}

	agentClusterInstall := &hiveext.AgentClusterInstall{}
	aciName := clusterDeployment.Spec.ClusterInstallRef.Name
	depRes = types.NamespacedName{
		Namespace: clusterDeployment.Namespace,
		Name:      aciName,
	}
	if err = ctrlClient.Get(ctx, depRes, agentClusterInstall); err != nil {
		failOnError(err, "Failed to get AgentClusterInstall using the following namespace and name %v", depRes)
	}

	log.Infof("Found AgentClusterInstall: %s -> %s", agentClusterInstall.Namespace, agentClusterInstall.Name)

	bmac := aictrl.BMACReconciler{Client: ctrlClient}
	agent := bmac.FindAgent(ctx, bmh)

	if agent == nil {
		log.Fatalf("Agent match error: \n No agent found for BMH: %s/%s \n Make sure the BMH BootMacAddress (%s) matches one of the interfaces in the node", bmh.Namespace, bmh.Name, bmh.Spec.BootMACAddress)
	}

	log.Infof("Found Agent: %s -> %s", agent.Namespace, agent.Name)

	if !agent.Spec.Approved {
		failOnError(err, "Agent %s is not approved", agent.Name)
	}

}

var rootCmd = &cobra.Command{
	Use:   "assisted-service-debugger [subcommand]",
	Short: "Debugger for Assisted Service based deployments",
	Run:   rootCmdRun,
}

func init() {
	_ = envconfig.Process(common.EnvConfigPrefix, &cli.Options)
	log = cli.InitLogs()
	// cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("bmh", "", "Starting BareMetalHost", "The BMH to deploy")
	viper.BindPFlag("bmh", rootCmd.PersistentFlags().Lookup("bmh"))
}

func main() {
	rootCmd.Execute()
}
