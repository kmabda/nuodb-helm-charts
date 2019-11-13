# Getting Started with NuoDB/Helm on OpenShift

This post will walk you through getting Helm up and running with Kubernetes or OpenShift, and then installing your first NuoDB Helm Chart. It assumes that you already have the OpenShift `oc client` installed locally and that you are logged into your OpenShift instance.

**[Deploying NuoDB using Helm Charts][5]** covers how to configure hosts to permit running NuoDB, and covers deploying your first NuoDB database using the provided Helm charts.

Bear in mind there are sub-charts in subdirectories included in this distribution. Instructions provided in this specific README are more geared towards the prerequisite setup of projects, Helm and Tiller, security settings, etc. Sub-charts have details for each of the deployed components.

## Deploying NuoDB using Helm Charts

    IMPORTANT:

    You MUST first disable THP on nodes where NuoDB will run. Run the transparent-hugepage` chart first, or manually setup tuned profiles to do this for you.

In a nutshell the order of installation is:

- **transparent-hugepage** ([documentation](transparent-hugepage/README.md))
- **admin** ([documentation](admin/README.md))
- **monitoring-influx** ([documentation](monitoring-influx/README.md))
- **database** ([documentation](database/README.md))
- **backup** ([documentation](backup/README.md))
- **restore** ([documentation](restore/README.md))
- **demo-ycsb** ([documentation](demo-ycsb/README.md))
- **demo-quickstart** ([documentation](demo-quickstart/README.md))

See the instructions for the individual charts for deploying the applications.

## Cleanup

See the instructions for the individual charts for deleting the applications.

An alternative cleanup strategy is to delete the entire project:

```bash
oc delete project <project-name>
```

## References

1. Materials herein unscrupulously stolen from [an online article written by Jim Minter, Red Hat, September 21, 2017][0].

[0]: https://blog.openshift.com/getting-started-helm-openshift/
[1]: https://helm.sh/docs/using_helm/
[2]: https://github.com/helm/helm/releases
[3]: https://docs.google.com/document/d/1G1Ljwe0c97KsH881QPUZK6ZtIShCk8jkxskXehuLpKw/edit#
[4]: #getting-started-with-helm-on-openshift
[5]: #deploying-nuodb-using-helm-charts
