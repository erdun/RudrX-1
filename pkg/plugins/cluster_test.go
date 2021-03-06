package plugins

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"

	"cuelang.org/go/cue"

	"github.com/cloud-native-application/rudrx/api/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var _ = Describe("DefinitionFiles", func() {

	route := types.Template{
		Name: "route",
		Type: types.TypeTrait,
		Parameters: []types.Parameter{
			{
				Name:     "domain",
				Required: true,
				Default:  "",
				Type:     cue.StringKind,
			},
		},
	}
	deployment := types.Template{
		Name: "deployment",
		Type: types.TypeWorkload,
		Parameters: []types.Parameter{
			{
				Name:     "name",
				Required: true,
				Type:     cue.StringKind,
				Default:  "",
			},
			{
				Type: cue.ListKind,
				Name: "env",
			},
			{
				Name:     "image",
				Type:     cue.StringKind,
				Default:  "",
				Short:    "i",
				Required: true,
				Usage:    "specify app image",
			},
			{
				Name:    "port",
				Type:    cue.IntKind,
				Short:   "p",
				Default: int64(8080),
				Usage:   "specify port for container",
			},
		},
	}
	req, _ := labels.NewRequirement("usecase", selection.Equals, []string{"forplugintest"})
	selector := labels.NewSelector().Add(*req)

	// Notice!!  DefinitionPath Object is Cluster Scope object
	// which means objects created in other DefinitionNamespace will also affect here.
	It("gettrait", func() {
		traitDefs, err := GetTraitsFromCluster(context.Background(), DefinitionNamespace, k8sClient, definitionDir, selector)
		Expect(err).Should(BeNil())
		logf.Log.Info(fmt.Sprintf("Getting trait definitions %v", traitDefs))
		for i := range traitDefs {
			traitDefs[i].Template = ""
			traitDefs[i].DefinitionPath = ""
		}
		Expect(traitDefs).Should(Equal([]types.Template{route}))
	})
	// Notice!!  DefinitionPath Object is Cluster Scope object
	// which means objects created in other DefinitionNamespace will also affect here.
	It("getworkload", func() {
		workloadDefs, err := GetWorkloadsFromCluster(context.Background(), DefinitionNamespace, k8sClient, definitionDir, selector)
		Expect(err).Should(BeNil())
		logf.Log.Info(fmt.Sprintf("Getting workload definitions  %v", workloadDefs))
		for i := range workloadDefs {
			workloadDefs[i].Template = ""
			workloadDefs[i].DefinitionPath = ""
		}
		Expect(workloadDefs).Should(Equal([]types.Template{deployment}))
	})
	It("getall", func() {
		alldef, err := GetTemplatesFromCluster(context.Background(), DefinitionNamespace, k8sClient, definitionDir, selector)
		Expect(err).Should(BeNil())
		logf.Log.Info(fmt.Sprintf("Getting all definitions %v", alldef))
		for i := range alldef {
			alldef[i].Template = ""
			alldef[i].DefinitionPath = ""
		}
		Expect(alldef).Should(Equal([]types.Template{deployment, route}))
	})
})
