package machine

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	machineapiv1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	awsproviderapi "sigs.k8s.io/cluster-api-provider-aws/pkg/apis/awsproviderconfig/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	_ reconcile.Reconciler = &ReconcileMachine{}
)

func TestReconcileRequest(t *testing.T) {
	awsCodec, err := awsproviderapi.NewCodec()
	if err != nil {
		log.Error(err, "Error creating AWSProviderConfigCodec")
	}

	// create awsproviderapi AWSMachineProviderStatus objects for each mock node
	workerNodeProviderStatus := awsproviderapi.AWSMachineProviderStatus{
		InstanceID:    aws.String("worker001"),
		InstanceState: aws.String("running"),
	}
	infraNodeProviderStatus := awsproviderapi.AWSMachineProviderStatus{
		InstanceID:    aws.String("infra001"),
		InstanceState: aws.String("running"),
	}
	master1NodeProviderStatus := awsproviderapi.AWSMachineProviderStatus{
		InstanceID:    aws.String("master001"),
		InstanceState: aws.String("running"),
	}
	master2NodeProviderStatus := awsproviderapi.AWSMachineProviderStatus{
		InstanceID:    aws.String("master002"),
		InstanceState: aws.String("running"),
	}
	master3NodeProviderStatus := awsproviderapi.AWSMachineProviderStatus{
		InstanceID:    aws.String("master003"),
		InstanceState: aws.String("running"),
	}

	// Encode each AWSMachineProviderStatus object to get runtime.RawExtension
	// that machineapiv1 Machine object needs for Provider Status
	workerNodeProviderStatusObj := interface{}(workerNodeProviderStatus).(*runtime.Object)
	workerProviderStatus, err := awsCodec.EncodeProviderStatus(*workerNodeProviderStatusObj)
	if err != nil {
		log.Error(err, "Error creating workerProviderStatus")
	}
	infraNodeProviderStatusObj := interface{}(infraNodeProviderStatus).(*runtime.Object)
	infraProviderStatus, err := awsCodec.EncodeProviderStatus(*infraNodeProviderStatusObj)
	if err != nil {
		log.Error(err, "Error creating infraProviderStatus")
	}
	master1NodeProviderStatusObj := interface{}(master1NodeProviderStatus).(*runtime.Object)
	master1ProviderStatus, err := awsCodec.EncodeProviderStatus(*master1NodeProviderStatusObj)
	if err != nil {
		log.Error(err, "Error creating master1ProviderStatus")
	}
	master2NodeProviderStatusObj := interface{}(master2NodeProviderStatus).(*runtime.Object)
	master2ProviderStatus, err := awsCodec.EncodeProviderStatus(*master2NodeProviderStatusObj)
	if err != nil {
		log.Error(err, "Error creating master2ProviderStatus")
	}
	master3NodeProviderStatusObj := interface{}(master3NodeProviderStatus).(*runtime.Object)
	master3ProviderStatus, err := awsCodec.EncodeProviderStatus(*master3NodeProviderStatusObj)
	if err != nil {
		log.Error(err, "Error creating master3ProviderStatus")
	}

	// Create each node as machineapiv1 Machine object, using encoded
	// AWSMachineProviderStatus objects from above
	workerNode := machineapiv1.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "worker001",
			Namespace: "default",
			Labels: map[string]string{
				"machine.openshift.io/cluster-api-machine-type": "worker",
			},
		},
		Spec: machineapiv1.MachineSpec{
			ProviderSpec: machineapiv1.ProviderSpec{
				Value: &runtime.RawExtension{
					Raw: []byte("{}"),
				},
			},
		},
		Status: machineapiv1.MachineStatus{
			ProviderStatus: workerProviderStatus,
		},
	}
	infraNode := machineapiv1.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "infra001",
			Namespace: "default",
			Labels: map[string]string{
				"machine.openshift.io/cluster-api-machine-type": "infra",
			},
		},
		Spec: machineapiv1.MachineSpec{
			ProviderSpec: machineapiv1.ProviderSpec{
				Value: &runtime.RawExtension{
					Raw: []byte("{}"),
				},
			},
		},
		Status: machineapiv1.MachineStatus{
			ProviderStatus: infraProviderStatus,
		},
	}
	master1Node := machineapiv1.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "master001",
			Namespace: "default",
			Labels: map[string]string{
				"machine.openshift.io/cluster-api-machine-type": "master",
			},
		},
		Spec: machineapiv1.MachineSpec{
			ProviderSpec: machineapiv1.ProviderSpec{
				Value: &runtime.RawExtension{
					Raw: []byte("{}"),
				},
			},
		},
		Status: machineapiv1.MachineStatus{
			ProviderStatus: master1ProviderStatus,
		},
	}
	master2Node := machineapiv1.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "master002",
			Namespace: "default",
			Labels: map[string]string{
				"machine.openshift.io/cluster-api-machine-type": "master",
			},
		},
		Spec: machineapiv1.MachineSpec{
			ProviderSpec: machineapiv1.ProviderSpec{
				Value: &runtime.RawExtension{
					Raw: []byte("{}"),
				},
			},
		},
		Status: machineapiv1.MachineStatus{
			ProviderStatus: master2ProviderStatus,
		},
	}
	master3Node := machineapiv1.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "master003",
			Namespace: "default",
			Labels: map[string]string{
				"machine.openshift.io/cluster-api-machine-type": "master",
			},
		},
		Spec: machineapiv1.MachineSpec{
			ProviderSpec: machineapiv1.ProviderSpec{
				Value: &runtime.RawExtension{
					Raw: []byte("{}"),
				},
			},
		},
		Status: machineapiv1.MachineStatus{
			ProviderStatus: master3ProviderStatus,
		},
	}

	// TO-DO: Create mock NLB instance to register instanceIDs to/deregister instanceIDs from
	// not sure if below is what I'm looking for?
	svc := elb.New(session.New())
	lbInput := &elb.CreateLoadBalancerInput{
		AvailabilityZones: []*string{
			aws.String("us-east-1a"),
		},
		Listeners: []*elb.Listener{
			{
				InstancePort:     aws.Int64(80),
				InstanceProtocol: aws.String("HTTP"),
				LoadBalancerPort: aws.Int64(80),
				Protocol:         aws.String("HTTP"),
			},
		},
		LoadBalancerName: aws.String("test-load-balancer"),
	}

	lbResult, err := svc.CreateLoadBalancer(lbInput)
	if err != nil {
		log.Error(err, "Error creating LB")
	}

	// TO-DO: How to add master instance IDs to NLB before tests run?
	// is this what I'm looking for? I'm concerned because the code above and
	// below both return funky objects that I'm not sure I can work with, and
	// I don't know if giving each function the definition of
	// `LoadBalancerName: aws.String("test-load-balancer")` is enough to
	// correlate them
	instanceInput := &elb.RegisterInstancesWithLoadBalancerInput{
		Instances: []*elb.Instance{
			{
				InstanceId: aws.String("master001"),
			},
			{
				InstanceId: aws.String("master002"),
			},
			{
				InstanceId: aws.String("master003"),
			},
		},
		LoadBalancerName: aws.String("test-load-balancer"),
	}

	instanceResult, err := svc.RegisterInstancesWithLoadBalancer(instanceInput)
	if err != nil {
		log.Error(err, "Error registering instance IDs with LB")
	}

	// TO-DO: How do I get _test.go to work through lines 194-216 of
	// machine_controller.go code?

	// TO-DO: How do I change nodes' instance IDs in order to trigger reconcile
	// function? This may be difficult as reconcile function works with
	// machineapiv1 Machine objects and instance IDs are found in awsproviderapi
	// AWSMachineProviderStatus objects
}
