// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPICoreV1PodSpec PodSpec is a description of a pod.
//
// swagger:model io.k8s.api.core.v1.PodSpec
type IoK8sAPICoreV1PodSpec struct {

	// Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
	ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty" json,yaml:"activeDeadlineSeconds,omitempty"`

	// If specified, the pod's scheduling constraints
	Affinity *IoK8sAPICoreV1Affinity `json:"affinity,omitempty" json,yaml:"affinity,omitempty"`

	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	AutomountServiceAccountToken bool `json:"automountServiceAccountToken,omitempty" json,yaml:"automountServiceAccountToken,omitempty"`

	// List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.
	// Required: true
	Containers []*IoK8sAPICoreV1Container `json:"containers" json,yaml:"containers"`

	// Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.
	DNSConfig *IoK8sAPICoreV1PodDNSConfig `json:"dnsConfig,omitempty" json,yaml:"dnsConfig,omitempty"`

	// Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	DNSPolicy string `json:"dnsPolicy,omitempty" json,yaml:"dnsPolicy,omitempty"`

	// EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.
	EnableServiceLinks bool `json:"enableServiceLinks,omitempty" json,yaml:"enableServiceLinks,omitempty"`

	// List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource. This field is alpha-level and is only honored by servers that enable the EphemeralContainers feature.
	EphemeralContainers []*IoK8sAPICoreV1EphemeralContainer `json:"ephemeralContainers" json,yaml:"ephemeralContainers"`

	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.
	HostAliases []*IoK8sAPICoreV1HostAlias `json:"hostAliases" json,yaml:"hostAliases"`

	// Use the host's ipc namespace. Optional: Default to false.
	HostIPC bool `json:"hostIPC,omitempty" json,yaml:"hostIPC,omitempty"`

	// Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.
	HostNetwork bool `json:"hostNetwork,omitempty" json,yaml:"hostNetwork,omitempty"`

	// Use the host's pid namespace. Optional: Default to false.
	HostPID bool `json:"hostPID,omitempty" json,yaml:"hostPID,omitempty"`

	// Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
	Hostname string `json:"hostname,omitempty" json,yaml:"hostname,omitempty"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []*IoK8sAPICoreV1LocalObjectReference `json:"imagePullSecrets" json,yaml:"imagePullSecrets"`

	// List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	InitContainers []*IoK8sAPICoreV1Container `json:"initContainers" json,yaml:"initContainers"`

	// NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	NodeName string `json:"nodeName,omitempty" json,yaml:"nodeName,omitempty"`

	// NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector map[string]string `json:"nodeSelector,omitempty" json,yaml:"nodeSelector,omitempty"`

	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.16, and is only honored by servers that enable the PodOverhead feature.
	Overhead map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"overhead,omitempty" json,yaml:"overhead,omitempty"`

	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	PreemptionPolicy string `json:"preemptionPolicy,omitempty" json,yaml:"preemptionPolicy,omitempty"`

	// The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
	Priority int32 `json:"priority,omitempty" json,yaml:"priority,omitempty"`

	// If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	PriorityClassName string `json:"priorityClassName,omitempty" json,yaml:"priorityClassName,omitempty"`

	// If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md
	ReadinessGates []*IoK8sAPICoreV1PodReadinessGate `json:"readinessGates" json,yaml:"readinessGates"`

	// Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	RestartPolicy string `json:"restartPolicy,omitempty" json,yaml:"restartPolicy,omitempty"`

	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of Kubernetes v1.14.
	RuntimeClassName string `json:"runtimeClassName,omitempty" json,yaml:"runtimeClassName,omitempty"`

	// If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
	SchedulerName string `json:"schedulerName,omitempty" json,yaml:"schedulerName,omitempty"`

	// SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.
	SecurityContext *IoK8sAPICoreV1PodSecurityContext `json:"securityContext,omitempty" json,yaml:"securityContext,omitempty"`

	// DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
	ServiceAccount string `json:"serviceAccount,omitempty" json,yaml:"serviceAccount,omitempty"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	ServiceAccountName string `json:"serviceAccountName,omitempty" json,yaml:"serviceAccountName,omitempty"`

	// If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.
	SetHostnameAsFQDN bool `json:"setHostnameAsFQDN,omitempty" json,yaml:"setHostnameAsFQDN,omitempty"`

	// Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.
	ShareProcessNamespace bool `json:"shareProcessNamespace,omitempty" json,yaml:"shareProcessNamespace,omitempty"`

	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Subdomain string `json:"subdomain,omitempty" json,yaml:"subdomain,omitempty"`

	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds,omitempty" json,yaml:"terminationGracePeriodSeconds,omitempty"`

	// If specified, the pod's tolerations.
	Tolerations []*IoK8sAPICoreV1Toleration `json:"tolerations" json,yaml:"tolerations"`

	// TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.
	TopologySpreadConstraints []*IoK8sAPICoreV1TopologySpreadConstraint `json:"topologySpreadConstraints" json,yaml:"topologySpreadConstraints"`

	// List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes
	Volumes []*IoK8sAPICoreV1Volume `json:"volumes" json,yaml:"volumes"`
}

// Validate validates this io k8s api core v1 pod spec
func (m *IoK8sAPICoreV1PodSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeralContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImagePullSecrets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverhead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadinessGates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTolerations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopologySpreadConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.Affinity) { // not required
		return nil
	}

	if m.Affinity != nil {
		if err := m.Affinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("affinity")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	for i := 0; i < len(m.Containers); i++ {
		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {
			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateDNSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSConfig) { // not required
		return nil
	}

	if m.DNSConfig != nil {
		if err := m.DNSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateEphemeralContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.EphemeralContainers) { // not required
		return nil
	}

	for i := 0; i < len(m.EphemeralContainers); i++ {
		if swag.IsZero(m.EphemeralContainers[i]) { // not required
			continue
		}

		if m.EphemeralContainers[i] != nil {
			if err := m.EphemeralContainers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ephemeralContainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ephemeralContainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateHostAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.HostAliases) { // not required
		return nil
	}

	for i := 0; i < len(m.HostAliases); i++ {
		if swag.IsZero(m.HostAliases[i]) { // not required
			continue
		}

		if m.HostAliases[i] != nil {
			if err := m.HostAliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostAliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostAliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateImagePullSecrets(formats strfmt.Registry) error {
	if swag.IsZero(m.ImagePullSecrets) { // not required
		return nil
	}

	for i := 0; i < len(m.ImagePullSecrets); i++ {
		if swag.IsZero(m.ImagePullSecrets[i]) { // not required
			continue
		}

		if m.ImagePullSecrets[i] != nil {
			if err := m.ImagePullSecrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePullSecrets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("imagePullSecrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateInitContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.InitContainers) { // not required
		return nil
	}

	for i := 0; i < len(m.InitContainers); i++ {
		if swag.IsZero(m.InitContainers[i]) { // not required
			continue
		}

		if m.InitContainers[i] != nil {
			if err := m.InitContainers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initContainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("initContainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateOverhead(formats strfmt.Registry) error {
	if swag.IsZero(m.Overhead) { // not required
		return nil
	}

	for k := range m.Overhead {

		if val, ok := m.Overhead[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateReadinessGates(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadinessGates) { // not required
		return nil
	}

	for i := 0; i < len(m.ReadinessGates); i++ {
		if swag.IsZero(m.ReadinessGates[i]) { // not required
			continue
		}

		if m.ReadinessGates[i] != nil {
			if err := m.ReadinessGates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("readinessGates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("readinessGates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateSecurityContext(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityContext) { // not required
		return nil
	}

	if m.SecurityContext != nil {
		if err := m.SecurityContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateTolerations(formats strfmt.Registry) error {
	if swag.IsZero(m.Tolerations) { // not required
		return nil
	}

	for i := 0; i < len(m.Tolerations); i++ {
		if swag.IsZero(m.Tolerations[i]) { // not required
			continue
		}

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateTopologySpreadConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.TopologySpreadConstraints) { // not required
		return nil
	}

	for i := 0; i < len(m.TopologySpreadConstraints); i++ {
		if swag.IsZero(m.TopologySpreadConstraints[i]) { // not required
			continue
		}

		if m.TopologySpreadConstraints[i] != nil {
			if err := m.TopologySpreadConstraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api core v1 pod spec based on the context it is used
func (m *IoK8sAPICoreV1PodSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEphemeralContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImagePullSecrets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverhead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReadinessGates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTolerations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopologySpreadConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.Affinity != nil {
		if err := m.Affinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("affinity")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Containers); i++ {

		if m.Containers[i] != nil {
			if err := m.Containers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateDNSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSConfig != nil {
		if err := m.DNSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateEphemeralContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EphemeralContainers); i++ {

		if m.EphemeralContainers[i] != nil {
			if err := m.EphemeralContainers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ephemeralContainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ephemeralContainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateHostAliases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HostAliases); i++ {

		if m.HostAliases[i] != nil {
			if err := m.HostAliases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostAliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostAliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateImagePullSecrets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImagePullSecrets); i++ {

		if m.ImagePullSecrets[i] != nil {
			if err := m.ImagePullSecrets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePullSecrets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("imagePullSecrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateInitContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InitContainers); i++ {

		if m.InitContainers[i] != nil {
			if err := m.InitContainers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initContainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("initContainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateOverhead(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Overhead {

		if val, ok := m.Overhead[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateReadinessGates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReadinessGates); i++ {

		if m.ReadinessGates[i] != nil {
			if err := m.ReadinessGates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("readinessGates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("readinessGates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateSecurityContext(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityContext != nil {
		if err := m.SecurityContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateTolerations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tolerations); i++ {

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateTopologySpreadConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopologySpreadConstraints); i++ {

		if m.TopologySpreadConstraints[i] != nil {
			if err := m.TopologySpreadConstraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSpec) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
