# Kubernetes from Scratch

> “What I cannot create, I do not understand” - Richard Feynman

![Go/Kubernetes](https://github.com/reece-iriye/kubernetes/blob/main/assets/images/go-kubernetes.png?raw=true)

A basic, lightweight, & distributed implementation of Kubernetes written fully in Go. I am doing this to learn more about Kubernetes' internal architecture and how the control and data plane work alongside each other. This implementation will require unit tests to be implemented for every single feature. I am also creating a YouTube series where I dive into how Kubernetes works by presenting concepts then coding them alongside the viewer.

## Kubernetes Resources and Operations

### Core Resources

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| Pods                      | ✅            |          |
| Namespaces                |             |          |
| Nodes                     |             |          |
| Services                  |             |          |
| Endpoints                 |             |          |
| ConfigMaps                |             |          |
| Secrets                   |             |          |
| PersistentVolumes (PV)    |             |          |
| PersistentVolumeClaims (PVC) |         |          |
| ResourceQuotas            |             |          |
| LimitRanges               |             |          |

## Workload Resources

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| Deployments               |             |          |
| ReplicaSets               |             |          |
| StatefulSets              |             |          |
| DaemonSets                |             |          |
| Jobs                      |             |          |
| CronJobs                  |             |          |

## Networking Resources

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| Ingress                   |             |          |
| NetworkPolicies           |             |          |

## Authentication & Authorization Resources

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| ServiceAccounts           |             |          |
| Roles                     |             |          |
| RoleBindings              |             |          |
| ClusterRoles              |             |          |
| ClusterRoleBindings       |             |          |

## Custom Resource Definitions (CRDs)

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| CustomResourceDefinitions (CRDs) |       |          |
| Any custom resources defined by CRDs |   |          |

## Cluster Management

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| Events                    |             |          |
| ClusterVersions (OpenShift-specific) | |          |
| ComponentStatuses         |             |          |

## Storage and Volume Resources

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| VolumeAttachments         |             |          |
| StorageClasses            |             |          |
| CSIDrivers                |             |          |
| CSIStorageCapacities      |             |          |

## Metrics & Autoscaling

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| HorizontalPodAutoscalers (HPA) |        |          |
| VerticalPodAutoscalers (if enabled) |    |          |

## API Access Control

| Resource                  | In Progress | Completed |
|---------------------------|-------------|-----------|
| AdmissionControllers      |             |          |
| MutatingWebhookConfigurations |        |          |
| ValidatingWebhookConfigurations |      |          |

## Supported Operations

| Operation                 | In Progress | Completed |
|---------------------------|-------------|-----------|
| Get                       |             |          |
| Describe                  |             |          |
| Create                    |             |          |
| Apply                    |              |          |
| Delete                   |              |          |
| Edit                     |              |          |
| Logs                     |              |          |
| Exec                     |              |          |
