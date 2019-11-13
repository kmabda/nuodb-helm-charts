# NuoDB Storage Class Helm Chart

This chart installs NuoDB storage classes in a Kubernetes cluster using the Helm package manager.

## TL;DR;

```bash
helm install nuodb/storage-class
```

## Prerequisites

- Kubernetes 1.9+

## Installing the Chart

### Configuration

The configuration is structured where configuration values are implemented following a single-definition rule, that is, values are structured and scoped, and shared across charts; e.g. for admin, its parameters are specified once in a single values file which is used for all the charts, and the database chart can use admin values for configuring connectivity of engines to a specific admin process. The same goes for other values **shared** amongst Helm charts. A few key points here:

- values files have structure, values are scoped
- different values files for different deployments
- values files follow the single definition rule (no repeats)
- global configuration exists under its own scoped section
- each chart has its own scoped section named after it
- cloud information is used to drive availability zones (particularly)

All configurable parameters for each top-level scope is detailed below, organized by scope.

#### cloud.*

The purpose of this section is to specify the cloud provider, and specify the availability zones where a solution is deployed.

The following tables list the configurable parameters for the `cloud` option:

| Parameter | Description | Default |
| ----- | ----------- | ------ |
| `provider` | Cloud provider; permissible values include: `azure`, `amazon`, or `google` |`nil`|
| `zones` | List of availability zones to deploy to |`[]`|

For example, for the Google Cloud:

```yaml
cloud:
  provider: google
  zones:
    - us-central1-a
    - us-central1-b
    - us-central1-c
```

### Permissions

This chart installs storage classes required for the operation of NuoDB.
Since storage classes are cluster-scoped objects, in order to install the
chart, the user installing the chart must have cluster-role permissions.

### Running

Verify the Helm chart:

```bash
helm install storage-class nuodb/storage-class \
    --debug --dry-run
```

Deploy the storage classes:

```bash
helm install storage-class nuodb/storage-class
```

## Uninstalling the Chart

To uninstall/delete the deployment:

```bash
helm del storage-class
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

[0]: #permissions
