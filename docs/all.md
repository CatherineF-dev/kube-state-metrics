# CertificateSigningRequest Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_certificatesigningrequest_annotations | Gauge | `certificatesigningrequest`=&lt;certificatesigningrequest-name&gt; <br> `signer_name`=&lt;certificatesigningrequest-signer-name&gt;| EXPERIMENTAL |
| kube_certificatesigningrequest_created| Gauge | `certificatesigningrequest`=&lt;certificatesigningrequest-name&gt; <br> `signer_name`=&lt;certificatesigningrequest-signer-name&gt;| STABLE |
| kube_certificatesigningrequest_condition | Gauge | `certificatesigningrequest`=&lt;certificatesigningrequest-name&gt; <br> `signer_name`=&lt;certificatesigningrequest-signer-name&gt; <br> `condition`=&lt;approved\|denied&gt; | STABLE |
| kube_certificatesigningrequest_labels | Gauge | `certificatesigningrequest`=&lt;certificatesigningrequest-name&gt; <br> `signer_name`=&lt;certificatesigningrequest-signer-name&gt;| STABLE |
| kube_certificatesigningrequest_cert_length | Gauge | `certificatesigningrequest`=&lt;certificatesigningrequest-name&gt; <br> `signer_name`=&lt;certificatesigningrequest-signer-name&gt;| STABLE |
# ClusterRole Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_clusterrole_annotations | Gauge | `clusterrole`=&lt;clusterrole-name&gt; | EXPERIMENTAL
| kube_clusterrole_labels | Gauge | `clusterrole`=&lt;clusterrole-name&gt; | EXPERIMENTAL
| kube_clusterrole_info | Gauge | `clusterrole`=&lt;clusterrole-name&gt; | EXPERIMENTAL |
| kube_clusterrole_created  | Gauge | `clusterrole`=&lt;clusterrole-name&gt; | EXPERIMENTAL |
| kube_clusterrole_metadata_resource_version | Gauge | `clusterrole`=&lt;clusterrole-name&gt; | EXPERIMENTAL |
# ConfigMap Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_configmap_annotations | Gauge | `configmap`=&lt;configmap-name&gt; <br> `namespace`=&lt;configmap-namespace&gt; <br> `annotation_CONFIGMAP_ANNOTATION`=&lt;CONFIGMAP_ANNOTATION&gt; | EXPERIMENTAL
| kube_configmap_labels | Gauge | `configmap`=&lt;configmap-name&gt; <br> `namespace`=&lt;configmap-namespace&gt; <br> `label_CONFIGMAP_LABEL`=&lt;CONFIGMAP_LABEL&gt; | STABLE
| kube_configmap_info | Gauge | `configmap`=&lt;configmap-name&gt; <br> `namespace`=&lt;configmap-namespace&gt; | STABLE |
| kube_configmap_created  | Gauge | `configmap`=&lt;configmap-name&gt; <br> `namespace`=&lt;configmap-namespace&gt; | STABLE |
| kube_configmap_metadata_resource_version | Gauge | `configmap`=&lt;configmap-name&gt; <br> `namespace`=&lt;configmap-namespace&gt; | EXPERIMENTAL |
# CronJob Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_cronjob_annotations | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; <br> `annotation_CRONJOB_ANNOTATION`=&lt;CRONJOB_ANNOTATION&gt;  | EXPERIMENTAL
| kube_cronjob_info | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; <br> `schedule`=&lt;schedule&gt; <br> `concurrency_policy`=&lt;concurrency-policy&gt; | STABLE
| kube_cronjob_labels | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; <br> `label_CRONJOB_LABEL`=&lt;CRONJOB_LABEL&gt;  | STABLE
| kube_cronjob_created  | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_next_schedule_time  | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_status_active | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_status_last_schedule_time | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_status_last_successful_time  | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | EXPERIMENTAL
| kube_cronjob_spec_suspend | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_spec_starting_deadline_seconds | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_metadata_resource_version| Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | STABLE
| kube_cronjob_spec_successful_job_history_limit | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | EXPERIMENTAL
| kube_cronjob_spec_failed_job_history_limit | Gauge | `cronjob`=&lt;cronjob-name&gt; <br> `namespace`=&lt;cronjob-namespace&gt; | EXPERIMENTAL
# DaemonSet Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_daemonset_annotations | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; <br> `annotation_DAEMONSET_ANNOTATION`=&lt;DAEMONSET_ANNOTATION&gt; | EXPERIMENTAL |
| kube_daemonset_created | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_current_number_scheduled | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_desired_number_scheduled | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_number_available | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_number_misscheduled | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_number_ready | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_number_unavailable | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_observed_generation | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_status_updated_number_scheduled | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_metadata_generation | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; | STABLE |
| kube_daemonset_labels | Gauge | `daemonset`=&lt;daemonset-name&gt; <br> `namespace`=&lt;daemonset-namespace&gt; <br> `label_DAEMONSET_LABEL`=&lt;DAEMONSET_LABEL&gt; | STABLE |
# Deployment Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_deployment_annotations | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; <br> `annotation_DEPLOYMENT_ANNOTATION`=&lt;DEPLOYMENT_ANNOTATION&gt; | EXPERIMENTAL |
| kube_deployment_status_replicas | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_status_replicas_ready | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | EXPERIMENTAL |
| kube_deployment_status_replicas_available | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_status_replicas_unavailable | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_status_replicas_updated | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_status_observed_generation | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_status_condition | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; <br> `condition`=&lt;deployment-condition&gt; <br> `status`=&lt;true\|false\|unknown&gt; | STABLE |
| kube_deployment_spec_replicas | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_spec_paused | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_spec_strategy_rollingupdate_max_unavailable | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_spec_strategy_rollingupdate_max_surge | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_metadata_generation | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
| kube_deployment_labels | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; <br> `label_DEPLOYMENT_LABEL`=&lt;DEPLOYMENT_LABEL&gt; | STABLE |
| kube_deployment_created | Gauge | `deployment`=&lt;deployment-name&gt; <br> `namespace`=&lt;deployment-namespace&gt; | STABLE |
# Endpoint Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_endpoint_annotations | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; <br> `annotation_ENDPOINT_ANNOTATION`=&lt;ENDPOINT_ANNOTATION&gt;  | EXPERIMENTAL |
| kube_endpoint_address_not_ready | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; | STABLE |
| kube_endpoint_address_available | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; | STABLE |
| kube_endpoint_info | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt;  | STABLE |
| kube_endpoint_labels | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; <br> `label_ENDPOINT_LABEL`=&lt;ENDPOINT_LABEL&gt;  | STABLE |
| kube_endpoint_created | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; | STABLE |
| kube_endpoint_ports | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; <br> `port_name`=&lt;endpoint-port-name&gt; <br> `port_protocol`=&lt;endpoint-port-protocol&gt; <br> `port_number`=&lt;endpoint-port-number&gt; | EXPERIMENTAL |
| kube_endpoint_address | Gauge | `endpoint`=&lt;endpoint-name&gt; <br> `namespace`=&lt;endpoint-namespace&gt; <br> `ip`=&lt;endpoint-ip&gt; <br> `ready`=&lt;true if available, false if unavailalbe&gt; | EXPERIMENTAL |
# Horizontal Pod Autoscaler Metrics

| Metric name                       | Metric type | Labels/tags                                                   | Status |
| --------------------------------  | ----------- | ------------------------------------------------------------- | ------ |
| kube_horizontalpodautoscaler_info                     | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; <br> `scaletargetref_api_version`=&lt;hpa-target-api-version&gt; <br> `scaletargetref_kind`=&lt;hpa-target-kind&gt; <br> `scaletargetref_name`=&lt;hpa-target-name&gt; | EXPERIMENTAL |
| kube_horizontalpodautoscaler_annotations              | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | EXPERIMENTAL |
| kube_horizontalpodautoscaler_labels                   | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | STABLE |
| kube_horizontalpodautoscaler_metadata_generation      | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | STABLE |
| kube_horizontalpodautoscaler_spec_max_replicas        | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | STABLE |
| kube_horizontalpodautoscaler_spec_min_replicas        | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | STABLE |
| kube_horizontalpodautoscaler_spec_target_metric       | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; <br> `metric_name`=&lt;metric-name&gt; <br> `metric_target_type`=&lt;value\|utilization\|average&gt; | EXPERIMENTAL |
| kube_horizontalpodautoscaler_status_condition         | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; <br> `condition`=&lt;hpa-condition&gt; <br> `status`=&lt;true\|false\|unknown&gt; | STABLE |
| kube_horizontalpodautoscaler_status_current_replicas  | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | STABLE |
| kube_horizontalpodautoscaler_status_desired_replicas  | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; | STABLE |
| kube_horizontalpodautoscaler_status_current_metric    | Gauge       | `horizontalpodautoscaler`=&lt;hpa-name&gt; <br> `namespace`=&lt;hpa-namespace&gt; <br> `metric_name`=&lt;metric-name&gt; <br> `metric_target_type`=&lt;metric-target-type&gt;  | EXPERIMENTAL |
# Ingress Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_ingress_annotations | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; <br> `annotation_INGRESS_ANNOTATION`=&lt;ANNOTATION_LABEL&gt; | EXPERIMENTAL |
| kube_ingress_info | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; <br> `ingressclass`=&lt;ingress-class&gt; or `_default` if not set | STABLE |
| kube_ingress_labels | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; <br> `label_INGRESS_LABEL`=&lt;INGRESS_LABEL&gt; | STABLE |
| kube_ingress_created  | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; | STABLE |
| kube_ingress_metadata_resource_version  | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; | EXPERIMENTAL |
| kube_ingress_path | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; <br> `host`=&lt;ingress-host&gt; <br> `path`=&lt;ingress-path&gt; <br> `service_name`=&lt;service name for the path&gt; <br> `service_port`=&lt;service port for the path&gt; | STABLE |
| kube_ingress_tls | Gauge | `ingress`=&lt;ingress-name&gt; <br> `namespace`=&lt;ingress-namespace&gt; <br> `tls_host`=&lt;tls hostname&gt; <br> `secret`=&lt;tls secret name&gt;| STABLE |
# Job Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_job_annotations | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; <br> `annotation_JOB_ANNOTATION`=&lt;JOB_ANNOTATION&gt;  | EXPERIMENTAL |
| kube_job_info | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_labels | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; <br> `label_JOB_LABEL`=&lt;JOB_LABEL&gt;  | STABLE |
| kube_job_owner | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; <br> `owner_kind`=&lt;owner kind&gt; <br> `owner_name`=&lt;owner name&gt; <br> `owner_is_controller`=&lt;whether owner is controller&gt;  | STABLE |
| kube_job_spec_parallelism | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_spec_completions | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_spec_active_deadline_seconds | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_status_active | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_status_succeeded | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_status_failed | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; <br> `reason`=&lt;failure reason&gt; | STABLE |
| kube_job_status_start_time | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_status_completion_time | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
| kube_job_complete | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; <br> `condition`=&lt;true\|false\|unknown&gt; | STABLE |
| kube_job_failed | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; <br> `condition`=&lt;true\|false\|unknown&gt; | STABLE |
| kube_job_created | Gauge | `job_name`=&lt;job-name&gt; <br> `namespace`=&lt;job-namespace&gt; | STABLE |
# Lease Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_lease_owner | Gauge | `lease`=&lt;lease-name&gt; <br> `owner_kind`=&lt;onwer kind&gt; <br> `owner_name`=&lt;owner name&gt; | EXPERIMENTAL |
| kube_lease_renew_time | Gauge | `lease`=&lt;lease-name&gt; | EXPERIMENTAL |
# LimitRange Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_limitrange | Gauge | `limitrange`=&lt;limitrange-name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;ResourceName&gt; <br> `type`=&lt;Pod\|Container\|PersistentVolumeClaim&gt; <br> `constraint`=&lt;constraint&gt;| STABLE |
| kube_limitrange_created | Gauge | `limitrange`=&lt;limitrange-name&gt; <br> `namespace`=&lt;namespace&gt; | STABLE |
# MutatingWebhookConfiguration Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_mutatingwebhookconfiguration_info | Gauge | `mutatingwebhookconfiguration`=&lt;mutatingwebhookconfiguration-name&gt; <br> `namespace`=&lt;mutatingwebhookconfiguration-namespace&gt; | EXPERIMENTAL |
| kube_mutatingwebhookconfiguration_created  | Gauge | `mutatingwebhookconfiguration`=&lt;mutatingwebhookconfiguration-name&gt; <br> `namespace`=&lt;mutatingwebhookconfiguration-namespace&gt; | EXPERIMENTAL |
| kube_mutatingwebhookconfiguration_metadata_resource_version | Gauge | `mutatingwebhookconfiguration`=&lt;mutatingwebhookconfiguration-name&gt; <br> `namespace`=&lt;mutatingwebhookconfiguration-namespace&gt; | EXPERIMENTAL |
# Namespace Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_namespace_annotations | Gauge | `namespace`=&lt;namespace-name&gt; <br> `label_NS_ANNOTATION`=&lt;NS_ANNOTATION&gt;  | EXPERIMENTAL |
| kube_namespace_created | Gauge | `namespace`=&lt;namespace-name&gt; | STABLE |
| kube_namespace_labels | Gauge | `namespace`=&lt;namespace-name&gt; <br> `label_NS_LABEL`=&lt;NS_LABEL&gt; | STABLE |
| kube_namespace_status_condition | Gauge | `namespace`=&lt;namespace-name&gt; <br> `condition`=&lt;NamespaceDeletionDiscoveryFailure\|NamespaceDeletionContentFailure\|NamespaceDeletionGroupVersionParsingFailure&gt;  <br> `status`=&lt;true\|false\|unknown&gt; | EXPERIMENTAL |
| kube_namespace_status_phase| Gauge | `namespace`=&lt;namespace-name&gt; <br> `phase`=&lt;Active\|Terminating&gt; | STABLE |
# Network Policy Metrics


| Metric name                           | Metric type | Labels/tags                                                                    | Status       |
| ------------------------------------- | ----------- | ------------------------------------------------------------------------------ | ------------ |
| kube_networkpolicy_annotations        | Gauge       | `namespace`=&lt;namespace name&gt; `networkpolicy`=&lt;networkpolicy name&gt;  | EXPERIMENTAL |
| kube_networkpolicy_created            | Gauge       | `namespace`=&lt;namespace name&gt; `networkpolicy`=&lt;networkpolicy name&gt;  | EXPERIMENTAL |
| kube_networkpolicy_labels             | Gauge       | `namespace`=&lt;namespace name&gt; `networkpolicy`=&lt;networkpolicy name&gt;  | EXPERIMENTAL |
| kube_networkpolicy_spec_egress_rules  | Gauge       | `namespace`=&lt;namespace name&gt; `networkpolicy`=&lt;networkpolicy name&gt;  | EXPERIMENTAL |
| kube_networkpolicy_spec_ingress_rules | Gauge       | `namespace`=&lt;namespace name&gt; `networkpolicy`=&lt;networkpolicy name&gt;  | EXPERIMENTAL |
# Node Metrics

| Metric name| Metric type | Description | Unit (where applicable) | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------------------- | ----------- | ------ |
| kube_node_annotations | Gauge | Kubernetes annotations converted to Prometheus labels | | `node`=&lt;node-address&gt; <br> `annotation_NODE_ANNOTATION`=&lt;NODE_ANNOTATION&gt;  | EXPERIMENTAL |
| kube_node_info | Gauge |  Information about a cluster node| |`node`=&lt;node-address&gt; <br> `kernel_version`=&lt;kernel-version&gt; <br> `os_image`=&lt;os-image-name&gt; <br> `container_runtime_version`=&lt;container-runtime-and-version-combination&gt; <br> `kubelet_version`=&lt;kubelet-version&gt; <br> `kubeproxy_version`=&lt;kubeproxy-version&gt; <br> `pod_cidr`=&lt;pod-cidr&gt; <br> `provider_id`=&lt;provider-id&gt; <br> `system_uuid`=&lt;system-uuid&gt; <br> `internal_ip`=&lt;internal-ip&gt; | STABLE |
| kube_node_labels | Gauge | Kubernetes labels converted to Prometheus labels | | `node`=&lt;node-address&gt; <br> `label_NODE_LABEL`=&lt;NODE_LABEL&gt;  | STABLE |
| kube_node_role | Gauge | The role of a cluster node | | `node`=&lt;node-address&gt; <br> `role`=&lt;NODE_ROLE&gt; | EXPERIMENTAL |
| kube_node_spec_unschedulable | Gauge | Whether a node can schedule new pods | | `node`=&lt;node-address&gt;| STABLE |
| kube_node_spec_taint | Gauge | The taint of a cluster node. | |`node`=&lt;node-address&gt; <br> `key`=&lt;taint-key&gt; <br> `value=`&lt;taint-value&gt; <br> `effect=`&lt;taint-effect&gt; | STABLE |
| kube_node_status_capacity | Gauge | The capacity for different resources of a node | `cpu`=&lt;core&gt; <br> `ephemeral_storage`=&lt;byte&gt; <br> `pods`=&lt;integer&gt; <br> `attachable_volumes_*`=&lt;byte&gt; <br> `hugepages_*`=&lt;byte&gt; <br> `memory`=&lt;byte&gt; |`node`=&lt;node-address&gt; <br> `resource`=&lt;resource-name&gt; <br> `unit`=&lt;resource-unit&gt;| STABLE |
| kube_node_status_allocatable | Gauge | The allocatable for different resources of a node that are available for scheduling | `cpu`=&lt;core&gt; <br> `ephemeral_storage`=&lt;byte&gt; <br> `pods`=&lt;integer&gt; <br> `attachable_volumes_*`=&lt;byte&gt; <br> `hugepages_*`=&lt;byte&gt; <br> `memory`=&lt;byte&gt; |`node`=&lt;node-address&gt; <br> `resource`=&lt;resource-name&gt; <br> `unit`=&lt;resource-unit&gt;| STABLE |
| kube_node_status_condition | Gauge | The condition of a cluster node | |`node`=&lt;node-address&gt; <br> `condition`=&lt;node-condition&gt; <br> `status`=&lt;true\|false\|unknown&gt; | STABLE |
| kube_node_created | Gauge | Unix creation timestamp | seconds |`node`=&lt;node-address&gt;| STABLE |
# PersistentVolumeClaim Metrics

| Metric name | Metric type | Description | Unit (where applicable) | Labels/tags | Status |
| ----------- | ------------- | ----------- | ----------- | ----------- | ----------- |
| kube_persistentvolumeclaim_annotations | Gauge | | | `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; <br> `namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `annotation_PERSISTENTVOLUMECLAIM_ANNOTATION`=&lt;PERSISTENTVOLUMECLAIM_ANNOATION&gt; | EXPERIMENTAL |
| kube_persistentvolumeclaim_access_mode | Gauge | | | `access_mode`=&lt;persistentvolumeclaim-access-mode&gt; <br>`namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; | STABLE |
| kube_persistentvolumeclaim_info | Gauge | | | `namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; <br> `storageclass`=&lt;persistentvolumeclaim-storageclassname&gt;<br>`volumename`=&lt;volumename&gt; | STABLE |
| kube_persistentvolumeclaim_labels | Gauge | | | `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; <br> `namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `label_PERSISTENTVOLUMECLAIM_LABEL`=&lt;PERSISTENTVOLUMECLAIM_LABEL&gt; | STABLE |
| kube_persistentvolumeclaim_resource_requests_storage_bytes | Gauge | | | `namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; | STABLE |
| kube_persistentvolumeclaim_status_condition | Gauge | | | `namespace` =&lt;persistentvolumeclaim-namespace&gt; <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; <br> `type`=&lt;persistentvolumeclaim-condition-type&gt; <br> `status`=&lt;true\false\unknown&gt; | EXPERIMENTAL |
| kube_persistentvolumeclaim_status_phase | Gauge | | | `namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; <br> `phase`=&lt;Pending\Bound\Lost&gt; | STABLE |
| kube_persistentvolumeclaim_created | Gauge | Unix Creation Timestamp | seconds | `namespace`=&lt;persistentvolumeclaim-namespace&gt; <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-name&gt; | EXPERIMENTAL |

Note:

- A special `<none>` string will be used if PVC has no storage class.
# PersistentVolume Metrics

| Metric name | Metric type | Description | Unit (where applicable) | Labels/tags | Status |
| ----------- | ----------- | ----------- | ----------- | ----------- | ------------ |
| kube_persistentvolume_annotations | Gauge | | | `persistentvolume`=&lt;persistentvolume-name&gt; <br> `annotation_PERSISTENTVOLUME_ANNOTATION`=&lt;PERSISTENTVOLUME_ANNOTATION&gt; | EXPERIMENTAL |
| kube_persistentvolume_capacity_bytes | Gauge | | | `persistentvolume`=&lt;pv-name&gt; | STABLE |
| kube_persistentvolume_status_phase | Gauge | | | `persistentvolume`=&lt;pv-name&gt; <br>`phase`=&lt;Bound\|Failed\|Pending\|Available\|Released&gt; | STABLE |
| kube_persistentvolume_claim_ref | Gauge | | | `persistentvolume`=&lt;pv-name&gt; <br>`claim_namespace`=&lt;<namespace>&gt; <br>`name`=&lt;<name>&gt; | STABLE |
| kube_persistentvolume_labels | Gauge | | | `persistentvolume`=&lt;persistentvolume-name&gt; <br> `label_PERSISTENTVOLUME_LABEL`=&lt;PERSISTENTVOLUME_LABEL&gt; | STABLE |
| kube_persistentvolume_info | Gauge | | | `persistentvolume`=&lt;pv-name&gt; <br> `storageclass`=&lt;storageclass-name&gt; <br> `gce_persistent_disk_name`=&lt;pd-name&gt; <br> `ebs_volume_id`=&lt;ebs-volume-id&gt; <br> `azure_disk_name`=&lt;azure-disk-name&gt; <br> `fc_wwids`=&lt;fc-wwids-comma-separated&gt; <br> `fc_lun`=&lt;fc-lun&gt; <br> `fc_target_wwns`=&lt;fc-target-wwns-comma-separated&gt; <br> `iscsi_target_portal`=&lt;iscsi-target-portal&gt; <br> `iscsi_iqn`=&lt;iscsi-iqn&gt; <br> `iscsi_lun`=&lt;iscsi-lun&gt; <br> `iscsi_initiator_name`=&lt;iscsi-initiator-name&gt; <br> `nfs_server`=&lt;nfs-server&gt; <br> `nfs_path`=&lt;nfs-path&gt; <br> `csi_driver`=&lt;csi-driver&gt; <br> `csi_volume_handle`=&lt;csi-volume-handle&gt; | STABLE |
| kube_persistentvolume_created | Gauge | Unix Creation Timestamp | seconds | `persistentvolume`=&lt;persistentvolume-name&gt; <br> | EXPERIMENTAL |

# PodDisruptionBudget Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_poddisruptionbudget_annotations | Gauge | `poddisruptionbudget`=&lt;poddisruptionbudget-name&gt; <br> `namespace`=&lt;poddisruptionbudget-namespace&gt; <br> `annotation_PODDISRUPTIONBUDGET_ANNOTATION`=&lt;PODDISRUPTIONBUDGET_ANNOATION&gt;  | EXPERIMENTAL |
| kube_poddisruptionbudget_labels | Gauge | `poddisruptionbudget`=&lt;poddisruptionbudget-name&gt; <br> `namespace`=&lt;poddisruptionbudget-namespace&gt; <br> `label_PODDISRUPTIONBUDGET_LABEL`=&lt;PODDISRUPTIONBUDGET_ANNOATION&gt;  | EXPERIMENTAL |
| kube_poddisruptionbudget_created | Gauge | `poddisruptionbudget`=&lt;pdb-name&gt; <br> `namespace`=&lt;pdb-namespace&gt;  | STABLE
| kube_poddisruptionbudget_status_current_healthy | Gauge | `poddisruptionbudget`=&lt;pdb-name&gt; <br> `namespace`=&lt;pdb-namespace&gt;  | STABLE
| kube_poddisruptionbudget_status_desired_healthy | Gauge | `poddisruptionbudget`=&lt;pdb-name&gt; <br> `namespace`=&lt;pdb-namespace&gt;  | STABLE
| kube_poddisruptionbudget_status_pod_disruptions_allowed | Gauge | `poddisruptionbudget`=&lt;pdb-name&gt; <br> `namespace`=&lt;pdb-namespace&gt;  | STABLE
| kube_poddisruptionbudget_status_expected_pods | Gauge | `poddisruptionbudget`=&lt;pdb-name&gt; <br> `namespace`=&lt;pdb-namespace&gt;  | STABLE
| kube_poddisruptionbudget_status_observed_generation | Gauge | `poddisruptionbudget`=&lt;pdb-name&gt; <br> `namespace`=&lt;pdb-namespace&gt;  | STABLE
# Pod Metrics

| Metric name| Metric type | Description | Unit (where applicable) | Labels/tags | Status | Opt-in |
| ---------- | ----------- | ----------- | ----------------------- | ----------- | ------ | ------ |
| kube_pod_annotations | Gauge | Kubernetes annotations converted to Prometheus labels | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `annotation_POD_ANNOTATION`=&lt;POD_ANNOTATION&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | -
| kube_pod_info | Gauge | Information about pod | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `host_ip`=&lt;host-ip&gt; <br> `pod_ip`=&lt;pod-ip&gt; <br> `node`=&lt;node-name&gt;<br> `created_by_kind`=&lt;created_by_kind&gt;<br> `created_by_name`=&lt;created_by_name&gt;<br> `uid`=&lt;pod-uid&gt;<br> `priority_class`=&lt;priority_class&gt;<br> `host_network`=&lt;host_network&gt;| STABLE | - |
| kube_pod_ips | Gauge | Pod IP addresses | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `ip`=&lt;pod-ip-address&gt; <br> `ip_family`=&lt;4 OR 6&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_start_time | Gauge | Start time in unix timestamp for a pod | seconds | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_completion_time | Gauge | Completion time in unix timestamp for a pod | seconds | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_owner | Gauge | Information about the Pod's owner | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `owner_kind`=&lt;owner kind&gt; <br> `owner_name`=&lt;owner name&gt; <br> `owner_is_controller`=&lt;whether owner is controller&gt; <br> `uid`=&lt;pod-uid&gt;  | STABLE | - |
| kube_pod_labels | Gauge | Kubernetes labels converted to Prometheus labels | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `label_POD_LABEL`=&lt;POD_LABEL&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_nodeselectors| Gauge | Describes the Pod nodeSelectors | |  `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `nodeselector_NODE_SELECTOR`=&lt;NODE_SELECTOR&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | Opt-in |
| kube_pod_status_phase | Gauge | The pods current phase | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `phase`=&lt;Pending\|Running\|Succeeded\|Failed\|Unknown&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_status_ready | Gauge | Describes whether the pod is ready to serve requests | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `condition`=&lt;true\|false\|unknown&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_status_scheduled | Gauge | Describes the status of the scheduling process for the pod | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `condition`=&lt;true\|false\|unknown&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_info | Gauge | Information about a container in a pod | | `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `image`=&lt;image-name&gt; <br> `image_id`=&lt;image-id&gt; <br> `image_spec`=&lt;image-spec&gt; <br> `container_id`=&lt;containerid&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_status_waiting | Gauge | Describes whether the container is currently in waiting state | | `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_status_waiting_reason | Gauge | Describes the reason the container is currently in waiting state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;container-waiting-reason&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_container_status_running | Gauge | Describes whether the container is currently in running state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_state_started | Gauge | Start time in unix timestamp for a pod container | seconds |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_status_terminated | Gauge | Describes whether the container is currently in terminated state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_status_terminated_reason | Gauge | Describes the reason the container is currently in terminated state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;container-terminated-reason&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_container_status_last_terminated_reason | Gauge | Describes the last reason the container was in terminated state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;last-terminated-reason&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_container_status_ready | Gauge | Describes whether the containers readiness check succeeded | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_status_restarts_total | Counter | The number of container restarts per container | | `container`=&lt;container-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `pod`=&lt;pod-name&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_container_resource_requests | Gauge | The number of requested request resource by a container | `cpu`=&lt;core&gt; <br> `memory`=&lt;bytes&gt; |`resource`=&lt;resource-name&gt; <br> `unit`=&lt;resource-unit&gt; <br> `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `node`=&lt; node-name&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_container_resource_limits | Gauge | The number of requested limit resource by a container | `cpu`=&lt;core&gt; <br> `memory`=&lt;bytes&gt; |`resource`=&lt;resource-name&gt; <br> `unit`=&lt;resource-unit&gt; <br> `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `node`=&lt; node-name&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_overhead_cpu_cores | Gauge | The pod overhead in regards to cpu cores associated with running a pod | core |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_overhead_memory_bytes | Gauge | The pod overhead in regards to memory associated with running a pod | bytes |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_runtimeclass_name_info | Gauge | The runtimeclass associated with the pod | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_created | Gauge | Unix creation timestamp | seconds |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_deletion_timestamp | Gauge | Unix deletion timestamp | seconds |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_restart_policy | Gauge | Describes the restart policy in use by this pod | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `type`=&lt;Always\|Never\|OnFailure&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_info | Gauge | Information about an init container in a pod | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `image`=&lt;image-name&gt; <br> `image_id`=&lt;image-id&gt; <br> `image_spec`=&lt;image-spec&gt; <br> `container_id`=&lt;containerid&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_status_waiting | Gauge | Describes whether the init container is currently in waiting state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_status_waiting_reason | Gauge | Describes the reason the init container is currently in waiting state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;container-waiting-reason&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_init_container_status_running | Gauge | Describes whether the init container is currently in running state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_status_terminated | Gauge | Describes whether the init container is currently in terminated state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_status_terminated_reason | Gauge | Describes the reason the init container is currently in terminated state | | `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;container-terminated-reason&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_init_container_status_last_terminated_reason | Gauge | Describes the last reason the init container was in terminated state | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;last-terminated-reason&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_init_container_status_ready | Gauge | Describes whether the init containers readiness check succeeded | |`container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_status_restarts_total | Counter | The number of restarts for the init container | integer |`container`=&lt;container-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `pod`=&lt;pod-name&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_init_container_resource_limits | Gauge | The number of CPU cores requested limit by an init container | `cpu`=&lt;core&gt; <br> `memory`=&lt;bytes&gt; |`resource`=&lt;resource-name&gt; <br> `unit`=&lt;resource-unit&gt; <br> `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `node`=&lt; node-name&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_init_container_resource_requests | Gauge | The number of CPU cores requested by an init container | `cpu`=&lt;core&gt; <br> `memory`=&lt;bytes&gt; |`resource`=&lt;resource-name&gt; <br> `unit`=&lt;resource-unit&gt; <br> `container`=&lt;container-name&gt; <br> `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `node`=&lt; node-name&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_spec_volumes_persistentvolumeclaims_info | Gauge | Information about persistentvolumeclaim volumes in a pod | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `volume`=&lt;volume-name&gt;  <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-claimname&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_spec_volumes_persistentvolumeclaims_readonly | Gauge | Describes whether a persistentvolumeclaim is mounted read only | bool |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt;  <br> `volume`=&lt;volume-name&gt;  <br> `persistentvolumeclaim`=&lt;persistentvolumeclaim-claimname&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_status_reason | Gauge | The pod status reasons | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `reason`=&lt;Evicted\|NodeAffinity\|NodeLost\|Shutdown\|UnexpectedAdmissionError&gt; <br> `uid`=&lt;pod-uid&gt; | EXPERIMENTAL | - |
| kube_pod_status_scheduled_time | Gauge | Unix timestamp when pod moved into scheduled status | seconds |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_status_unschedulable | Gauge | Describes the unschedulable status for the pod | |`pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; | STABLE | - |
| kube_pod_tolerations | Gauge | Information about the pod tolerations | | `pod`=&lt;pod-name&gt; <br> `namespace`=&lt;pod-namespace&gt; <br> `uid`=&lt;pod-uid&gt; <br> `key`=&lt;toleration-key&gt; <br> `operator`=&lt;toleration-operator&gt; <br> `value`=&lt;toleration-value&gt; <br> `effect`=&lt;toleration-effect&gt; `toleration_seconds`=&lt;toleration-seconds&gt; | EXPERIMENTAL | - |

## Useful metrics queries

### How to retrieve non-standard Pod state

It is not straightforward to get the Pod states for certain cases like "Terminating" and "Unknown" since it is not stored behind a field in the `Pod.Status`.

So to mimic the [logic](https://github.com/kubernetes/kubernetes/blob/v1.17.3/pkg/printers/internalversion/printers.go#L624) used by the `kubectl` command line, you will need to compose multiple metrics.

For example:

* To get the list of pods that are in the `Unknown` state, you can run the following PromQL query: `sum(kube_pod_status_phase{phase="Unknown"}) by (namespace, pod) or (count(kube_pod_deletion_timestamp) by (namespace, pod) * sum(kube_pod_status_reason{reason="NodeLost"}) by(namespace, pod))`

* For Pods in `Terminating` state: `count(kube_pod_deletion_timestamp) by (namespace, pod) * count(kube_pod_status_reason{reason="NodeLost"} == 0) by (namespace, pod)`

Here is an example of a Prometheus rule that can be used to alert on a Pod that has been in the `Terminated` state for more than `5m`.

```yaml
groups:
- name: Pod state
  rules:
  - alert: PodsBlockInTerminatingState
    expr: count(kube_pod_deletion_timestamp) by (namespace, pod) * count(kube_pod_status_reason{reason="NodeLost"} == 0) by (namespace, pod) > 0
    for: 5m
    labels:
      severity: page
    annotations:
      summary: Pod {{$labels.namespace}}/{{$labels.pod}} block in Terminating state.
```
# ReplicaSet metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_replicaset_annotations | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; <br> `annotation_REPLICASET_ANNOTATION`=&lt;REPLICASET_ANNOTATION&gt; | EXPERIMENTAL |
| kube_replicaset_status_replicas | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_status_fully_labeled_replicas | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_status_ready_replicas | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_status_observed_generation | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_spec_replicas | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_metadata_generation | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_labels | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; <br> `label_REPLICASET_LABEL`=&lt;REPLICASET_LABEL&gt; | STABLE |
| kube_replicaset_created | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; | STABLE |
| kube_replicaset_owner | Gauge | `replicaset`=&lt;replicaset-name&gt; <br> `namespace`=&lt;replicaset-namespace&gt; <br> `owner_kind`=&lt;owner kind&gt; <br> `owner_name`=&lt;owner name&gt; <br> `owner_is_controller`=&lt;whether owner is controller&gt;  | STABLE |
# ReplicationController metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_replicationcontroller_status_replicas | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_status_fully_labeled_replicas | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_status_ready_replicas | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_status_available_replicas | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_status_observed_generation | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_spec_replicas | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_metadata_generation | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_created | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; | STABLE |
| kube_replicationcontroller_owner | Gauge | `replicationcontroller`=&lt;replicationcontroller-name&gt; <br> `namespace`=&lt;replicationcontroller-namespace&gt; <br> `owner_kind`=&lt;owner kind&gt; <br> `owner_name`=&lt;owner name&gt; <br> `owner_is_controller`=&lt;whether owner is controller&gt;  | EXPERIMENTAL |
# ResourceQuota Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_resourcequota | Gauge | `resourcequota`=&lt;quota-name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;ResourceName&gt; <br> `type`=&lt;quota-type&gt; | STABLE |
| kube_resourcequota_created | Gauge | `resourcequota`=&lt;quota-name&gt; <br> `namespace`=&lt;namespace&gt; | STABLE |
# Role Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_role_annotations | Gauge | `role`=&lt;role-name&gt; <br> `namespace`=&lt;role-namespace&gt; | EXPERIMENTAL
| kube_role_labels | Gauge | `role`=&lt;role-name&gt; <br> `namespace`=&lt;role-namespace&gt; | EXPERIMENTAL
| kube_role_info | Gauge | `role`=&lt;role-name&gt; <br> `namespace`=&lt;role-namespace&gt; | EXPERIMENTAL
| kube_role_created  | Gauge | `role`=&lt;role-name&gt; <br> `namespace`=&lt;role-namespace&gt; | EXPERIMENTAL |
| kube_role_metadata_resource_version | Gauge | `role`=&lt;role-name&gt; <br> `namespace`=&lt;role-namespace&gt; | EXPERIMENTAL |
# Secret Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_secret_annotations | Gauge | `secret`=&lt;secret-name&gt; <br> `namespace`=&lt;secret-namespace&gt; <br> `annotations_SECRET_ANNOTATION`=&lt;SECRET_ANNOTATION&gt; | EXPERIMENTAL |
| kube_secret_info | Gauge | `secret`=&lt;secret-name&gt; <br> `namespace`=&lt;secret-namespace&gt; | STABLE |
| kube_secret_type | Gauge | `secret`=&lt;secret-name&gt; <br> `namespace`=&lt;secret-namespace&gt; <br> `type`=&lt;secret-type&gt; | STABLE |
| kube_secret_labels | Gauge | `secret`=&lt;secret-name&gt; <br> `namespace`=&lt;secret-namespace&gt; <br> `label_SECRET_LABEL`=&lt;SECRET_LABEL&gt; | STABLE |
| kube_secret_created  | Gauge | `secret`=&lt;secret-name&gt; <br> `namespace`=&lt;secret-namespace&gt; | STABLE |
| kube_secret_metadata_resource_version  | Gauge | `secret`=&lt;secret-name&gt; <br> `namespace`=&lt;secret-namespace&gt; | EXPERIMENTAL |
# Service Metrics

| Metric name                           | Metric type | Description                                                                    | Unit (where applicable) | Labels/tags                                                                                                                                                                                                          | Status       |
|---------------------------------------|-------------|--------------------------------------------------------------------------------|-------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------|
| kube_serviceaccount_info              | Gauge       | Information about a service account                                            |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt; <br> `automount_token`=&lt;serviceaccount-automount-token&gt;                   | EXPERIMENTAL |
| kube_serviceaccount_created           | Gauge       | Unix creation timestamp                                                        |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt;                                                                                 | EXPERIMENTAL |
| kube_serviceaccount_deleted           | Gauge       | Unix deletion timestamp                                                        |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt;                                                                                 | EXPERIMENTAL |
| kube_serviceaccount_secret            | Gauge       | Secret being referenced by a service account                                   |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt; <br> `name`=&lt;secret-name&gt;                                                 | EXPERIMENTAL |
| kube_serviceaccount_image_pull_secret | Gauge       | Secret being referenced by a service account for the purpose of pulling images |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt; <br> `name`=&lt;secret-name&gt;                                                 | EXPERIMENTAL |
| kube_serviceaccount_annotations       | Gauge       | Kubernetes annotations converted to Prometheus labels                          |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt; <br> `annotation_SERVICE_ACCOUNT_ANNOTATION`=&lt;SERVICE_ACCOUNT_ANNOTATION&gt; | EXPERIMENTAL |
| kube_serviceaccount_labels            | Gauge       | Kubernetes labels converted to Prometheus labels                               |                         | `namespace`=&lt;serviceaccount-namespace&gt; <br> `serviceaccount`=&lt;serviceaccount-name&gt; <br> `uid`=&lt;serviceaccount-uid&gt; <br> `label_SERVICE_ACCOUNT_LABEL`=&lt;SERVICE_ACCOUNT_LABEL&gt;                | EXPERIMENTAL |
# Service Metrics

| Metric name| Metric type | Description | Unit (where applicable) | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------------------- | ----------- | ------ |
| kube_service_annotations | Gauge | Kubernetes annotations converted to Prometheus labels | |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt; <br> `annotation_SERVICE_ANNOTATION`=&lt;SERVICE_ANNOTATION&gt;  | EXPERIMENTAL |
| kube_service_info | Gauge | Information about service | |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt; <br> `cluster_ip`=&lt;service cluster ip&gt; <br> `external_name`=&lt;service external name&gt; <br> `load_balancer_ip`=&lt;service load balancer ip&gt; | STABLE |
| kube_service_labels | Gauge | Kubernetes labels converted to Prometheus labels | |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt; <br> `label_SERVICE_LABEL`=&lt;SERVICE_LABEL&gt;  | STABLE |
| kube_service_created | Gauge | Unix creation timestamp | seconds |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt;  | STABLE |
| kube_service_spec_type | Gauge | Type about service | |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt; <br> `type`=&lt;ClusterIP\|NodePort\|LoadBalancer\|ExternalName&gt; | STABLE |
| kube_service_spec_external_ip | Gauge | Service external ips. One series for each ip | |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt; <br> `external_ip`=&lt;external-ip&gt; | STABLE |
| kube_service_status_load_balancer_ingress | Gauge | Service load balancer ingress status | |`service`=&lt;service-name&gt; <br> `namespace`=&lt;service-namespace&gt; <br> `uid`=&lt;service-uid&gt; <br> `ip`=&lt;load-balancer-ingress-ip&gt; <br> `hostname`=&lt;load-balancer-ingress-hostname&gt; | STABLE |
# Stateful Set Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_statefulset_annotations | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt; <br> `annotation_STATEFULSET_ANNOTATION`=&lt;STATEFULSET_ANNOTATION&gt; | EXPERIMENTAL |
| kube_statefulset_status_replicas | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_status_replicas_current | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_status_replicas_ready | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_status_replicas_available | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | EXPERIMENTAL |
| kube_statefulset_status_replicas_updated | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_status_observed_generation | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_replicas | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_metadata_generation | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_created | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt;  | STABLE |
| kube_statefulset_labels | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt; <br> `label_STATEFULSET_LABEL`=&lt;STATEFULSET_LABEL&gt; | STABLE |
| kube_statefulset_status_current_revision | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt; <br> `revision`=&lt;statefulset-current-revision&gt; | STABLE |
| kube_statefulset_status_update_revision | Gauge | `statefulset`=&lt;statefulset-name&gt; <br> `namespace`=&lt;statefulset-namespace&gt; <br> `revision`=&lt;statefulset-update-revision&gt; | STABLE |
# StorageClass Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_storageclass_annotations | Gauge | `storageclass`=&lt;storageclass-name&gt; <br> `annotation_STORAGECLASS_ANNOTATION`=&lt;STORAGECLASS_ANNOTATION&gt; | EXPERIMENTAL |
| kube_storageclass_info | Gauge | `storageclass`=&lt;storageclass-name&gt; <br> `provisioner`=&lt;storageclass-provisioner&gt; <br> `reclaim_policy`=&lt;storageclass-reclaimPolicy&gt; <br> `volume_binding_mode`=&lt;storageclass-volumeBindingMode&gt; | STABLE |
| kube_storageclass_labels | Gauge | `storageclass`=&lt;storageclass-name&gt; <br> `label_STORAGECLASS_LABEL`=&lt;STORAGECLASS_LABEL&gt; | STABLE |
| kube_storageclass_created  | Gauge | `storageclass`=&lt;storageclass-name&gt; | STABLE |
# ValidatingWebhookConfiguration Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_validatingwebhookconfiguration_info | Gauge | `validatingwebhookconfiguration`=&lt;validatingwebhookconfiguration-name&gt; <br> `namespace`=&lt;validatingwebhookconfiguration-namespace&gt; | EXPERIMENTAL |
| kube_validatingwebhookconfiguration_created  | Gauge | `validatingwebhookconfiguration`=&lt;validatingwebhookconfiguration-name&gt; <br> `namespace`=&lt;validatingwebhookconfiguration-namespace&gt; | EXPERIMENTAL |
| kube_validatingwebhookconfiguration_metadata_resource_version | Gauge | `validatingwebhookconfiguration`=&lt;validatingwebhookconfiguration-name&gt; <br> `namespace`=&lt;validatingwebhookconfiguration-namespace&gt; | EXPERIMENTAL |
# Vertical Pod Autoscaler Metrics

| Metric name                                                                | Metric type | Labels/tags                                                                                                                                                                                                                                                | Status                                                                                                                                                      |
| --------------------------------                                           | ----------- | -------------------------------------------------------------                                                                                                                                                                                              | ------                                                                                                                                                      |
| kube_verticalpodautoscaler_annotations                                          | Gauge       | `annotation_app`=&lt;foo&gt; <br> `namespace`=&lt;namespace&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;   | EXPERIMENTAL                                                                                                                                                |
| kube_verticalpodautoscaler_spec_resourcepolicy_container_policies_minallowed                   | Gauge       | `container`=&lt;container name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;cpu memory&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `unit`=&lt;core byte&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;                | EXPERIMENTAL |
| kube_verticalpodautoscaler_spec_resourcepolicy_container_policies_maxallowed                   | Gauge       | `container`=&lt;container name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;cpu memory&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `unit`=&lt;core byte&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;                | EXPERIMENTAL |
| kube_verticalpodautoscaler_status_recommendation_containerrecommendations_lowerbound     | Gauge       | `container`=&lt;container name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;cpu memory&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `unit`=&lt;core byte&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;                | EXPERIMENTAL |
| kube_verticalpodautoscaler_status_recommendation_containerrecommendations_target          | Gauge       | `container`=&lt;container name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;cpu memory&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `unit`=&lt;core byte&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;                | EXPERIMENTAL |
| kube_verticalpodautoscaler_status_recommendation_containerrecommendations_uncappedtarget | Gauge       | `container`=&lt;container name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;cpu memory&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `unit`=&lt;core byte&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;                | EXPERIMENTAL |
| kube_verticalpodautoscaler_status_recommendation_containerrecommendations_upperbound     | Gauge       | `container`=&lt;container name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;cpu memory&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `unit`=&lt;core byte&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;                | EXPERIMENTAL |
| kube_verticalpodautoscaler_labels                                          | Gauge       | `label_app`=&lt;foo&gt; <br> `namespace`=&lt;namespace&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt;   | EXPERIMENTAL                                                                                                                                                |
| kube_verticalpodautoscaler_spec_updatepolicy_updatemode                                     | Gauge       | `namespace`=&lt;namespace&gt; <br> `target_api_version`=&lt;api version&gt; <br> `target_kind`=&lt;target kind&gt; <br> `target_name`=&lt;target name&gt; <br> `update_mode`=&lt;foo&gt; <br> `verticalpodautoscaler`=&lt;vertical pod autoscaler name&gt; | EXPERIMENTAL                                                                                                                                                |

## Configuration

Vertical Pod Autoscalers(VPAs) are managed as custom resources.

To enable the Vertical Pod Autoscaler collector, please:

1. Ensure that the Vertical Pod Autoscaler CRDs are installed in the cluster. The CRDs are [here](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/deploy/vpa-beta2-crd.yaml).
2. Ensure that `verticalpodautoscalers` is included in list of `Resources` enabled using the flag `--resources` when `kube-state-metrics` is run (see below).

One of the [command line arguments](./docs/cli-arguments.md) for `kube-state-metrics` is `--resources`. If this flag is omitted, a default set of Resources is enabled. This default list does **not** include Vertical Pod Autoscalers.

To enable Vertical Pod Autoscalers, the `kube-state-metrics` flag `--resource` must be included when the binary is run and the list of resources must include `verticalpodautoscalers`.


### Examples

The following configures `kube-state-metrics` on the command line and in the `args` section of a Kubernetes manifest. Because neither command includes the `--resource` flag, the default set of resources will be include **but** metrics for Vertical Pod Autoscalers will **not** be included:

Shell:

```bash
kube-state-metrics \
--telemetry-port=8081 \
--kubeconfig=... \
--apiserver=...
```

Kubernetes:

```YAML
spec:
  template:
    spec:
      containers:
        - args:
          - --telemetry-port=8081
          - --kubeconfig=...
          - --apiserver=...
```

To include Vertical Pod Autoscaler metrics, you must include the `--resources` flag and to include the default resources, you must include the list of default resources **and** `verticalpodautoscalers`, i.e.:

Shell:

```bash
kube-state-metrics \
--telemetry-port=8081 \
--kubeconfig=... \
--apiserver=... \
--resources=certificatesigningrequests,configmaps,cronjobs,daemonsets,deployments,endpoints,horizontalpodautoscalers, ingresses,jobs,leases,limitranges,mutatingwebhookconfigurations,namespaces,networkpolicies,nodes,persistentvolumeclaims,persistentvolumes,poddisruptionbudgets,pods,replicasets,replicationcontrollers,resourcequotas,secrets,services,statefulsets,storageclasses,validatingwebhookconfigurations,verticalpodautoscalers,volumeattachments
```

Kubernetes:

```YAML
spec:
  template:
    spec:
      containers:
        - args:
          - --telemetry-port=8081
          - --kubeconfig=...
          - --apiserver=...
          - --resources=certificatesigningrequests,configmaps,cronjobs,daemonsets,deployments,endpoints,horizontalpodautoscalers, ingresses,jobs,leases,limitranges,mutatingwebhookconfigurations,namespaces,networkpolicies,nodes,persistentvolumeclaims,persistentvolumes,poddisruptionbudgets,pods,replicasets,replicationcontrollers,resourcequotas,secrets,services,statefulsets,storageclasses,validatingwebhookconfigurations,verticalpodautoscalers,volumeattachments
```

### Confirmation

To confirm that a `kube-state-metrics` process includes `verticalpodautoscalers`, you can:

Shell:

```bash
ps aux \
| grep kube-state-metrics \
| grep verticalpodautoscalers
```

Kubernetes: assuming your deployment is called `kube-state-metrics`:

```bash
DEPLOYMENT="kube-state-metrics"
NAMESPACE="default"

kubectl get deployment/${DEPLOYMENT} \
--namespace=${NAMESPACE} \
--output=jsonpath="{range .spec.template.spec.containers[?(@.name=='kube-state-metrics')].args[*]}{@}{'\n'}{end}"
```

Should include (among other `--flags`):

```console
--resources=certificatesigningrequests,configmaps,cronjobs,daemonsets,deployments,endpoints,horizontalpodautoscalers,ingresses,jobs,leases,limitranges,mutatingwebhookconfigurations,namespaces,networkpolicies,nodes,persistentvolumeclaims,persistentvolumes,poddisruptionbudgets,pods,replicasets,replicationcontrollers,resourcequotas,secrets,services,statefulsets,storageclasses,validatingwebhookconfigurations,verticalpodautoscalers,volumeattachments
```
# VolumeAttachment Metrics

| Metric name| Metric type | Labels/tags | Status |
| ---------- | ----------- | ----------- | ----------- |
| kube_volumeattachment_info | Gauge | `volumeattachment`=&lt;volumeattachment-name&gt; <br> `attacher`=&lt;attacher-name&gt; <br> `node`=&lt;node-name&gt; | EXPERIMENTAL |
| kube_volumeattachment_created | Gauge | `volumeattachment`=&lt;volumeattachment-name&gt; | EXPERIMENTAL |
| kube_volumeattachment_labels | Gauge | `volumeattachment`=&lt;volumeattachment-name&gt; <br> `label_VOLUMEATTACHMENT_LABEL`=&lt;VOLUMEATTACHMENT_LABEL&gt;  | EXPERIMENTAL |
| kube_volumeattachment_spec_source_persistentvolume | Gauge | `volumeattachment`=&lt;volumeattachment-name&gt; <br> `volumename`=&lt;persistentvolume-name&gt; | EXPERIMENTAL |
| kube_volumeattachment_status_attached | Gauge | `volumeattachment`=&lt;volumeattachment-name&gt; | EXPERIMENTAL |
| kube_volumeattachment_status_attachment_metadata | Gauge | `volumeattachment`=&lt;volumeattachment-name&gt; <br> `metadata_METADATA_KEY`=&lt;METADATA_VALUE&gt;  | EXPERIMENTAL |
