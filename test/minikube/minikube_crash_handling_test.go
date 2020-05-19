// +build long

package minikube

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/nuodb/nuodb-helm-charts/test/testlib"
	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	"testing"
)

func verifyKillAndInfoInLog(t *testing.T, namespaceName string, adminPodName string, podName string) {
	options := k8s.NewKubectlOptions("", "")
	options.Namespace = namespaceName

	// send SIGABRT
	k8s.RunKubectl(t, options, "exec", podName, "--", "kill", "-6", "1")
	testlib.AwaitPodRestartCountGreaterThan(t, namespaceName, podName, 0)

	testlib.AwaitDatabaseUp(t, namespaceName, adminPodName, "demo", 2)

	stringOccurrence := testlib.GetStringOccurrenceInLog(t, namespaceName, podName, "Core was generated by",
		&corev1.PodLogOptions {Previous:true})

	assert.Assert(t, stringOccurrence > 0, "Could not find core parsing in log file")
}

func TestKubernetesPrintCores(t *testing.T) {
	testlib.AwaitTillerUp(t)
	defer testlib.VerifyTeardown(t)

	defer testlib.Teardown(testlib.TEARDOWN_ADMIN)

	helmChartReleaseName, namespaceName := testlib.StartAdmin(t, &helm.Options{}, 1, "")

	admin0 := fmt.Sprintf("%s-nuodb-cluster0-0", helmChartReleaseName)

	defer testlib.Teardown(testlib.TEARDOWN_DATABASE) // ensure resources allocated in called functions are released when this function exits

	databaseHelmChartReleaseName := testlib.StartDatabase(t, namespaceName, admin0, &helm.Options{
		SetValues: map[string]string{
			"database.sm.resources.requests.cpu":    testlib.MINIMAL_VIABLE_ENGINE_CPU,
			"database.sm.resources.requests.memory": testlib.MINIMAL_VIABLE_ENGINE_MEMORY,
			"database.te.resources.requests.cpu":    testlib.MINIMAL_VIABLE_ENGINE_CPU,
			"database.te.resources.requests.memory": testlib.MINIMAL_VIABLE_ENGINE_MEMORY,
		},
	})

	t.Run("killTEWithCore", func(t *testing.T) {
		tePodNameTemplate := fmt.Sprintf("te-%s-nuodb-%s-%s", databaseHelmChartReleaseName, "cluster0", "demo")
		tePodName := testlib.GetPodName(t, namespaceName, tePodNameTemplate)
		verifyKillAndInfoInLog(t, namespaceName, admin0, tePodName)
	})


	t.Run("killSMWithCore", func(t *testing.T) {
		smPodTemplate := fmt.Sprintf("sm-%s-nuodb-%s-%s", databaseHelmChartReleaseName, "cluster0", "demo")
		smPodName := testlib.GetPodName(t, namespaceName, smPodTemplate)
		verifyKillAndInfoInLog(t, namespaceName, admin0, smPodName)
	})
}
