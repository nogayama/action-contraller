package actionpolicy

import (
	"context"

	actioncontrollerv1alpha1 "github.com/nogayama/action-contraller/pkg/apis/actioncontroller/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	
	netv1 "k8s.io/api/networking/v1"
//	"fmt"
)

var log = logf.Log.WithName("controller_actionpolicy")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ActionPolicy Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileActionPolicy{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("actionpolicy-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ActionPolicy
	err = c.Watch(&source.Kind{Type: &actioncontrollerv1alpha1.ActionPolicy{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ActionPolicy
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &actioncontrollerv1alpha1.ActionPolicy{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileActionPolicy implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileActionPolicy{}

// ReconcileActionPolicy reconciles a ActionPolicy object
type ReconcileActionPolicy struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ActionPolicy object and makes changes based on the state read
// and what is in the ActionPolicy.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileActionPolicy) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ActionPolicy")
	
	// Fetch the ActionPolicy instance
	instance := &actioncontrollerv1alpha1.ActionPolicy{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if (false) {
		// Define a new Pod object
		pod := newPodForCR(instance)
	
		// Set ActionPolicy instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}
	
		// Check if this Pod already exists
		found := &corev1.Pod{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
			err = r.client.Create(context.TODO(), pod)
			if err != nil {
				return reconcile.Result{}, err
			}
	
			// Pod created successfully - don't requeue
			return reconcile.Result{}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		}
	}

	var targetRuleName = "nginx.deny-egress"
	

//	netpo := &netv1.NetworkPolicy{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "hoge",
//			Namespace: request.Namespace,
//		},
//	}
//	fmt.Println(netpo)

	
	// Check if this Pod already exists
	found2 := &netv1.NetworkPolicy{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: targetRuleName, Namespace: instance.Namespace}, found2)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new NetworkPolicy", "Pod.Namespace", instance.Namespace, "Pod.Name", targetRuleName)
		
		netpo := &netv1.NetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      targetRuleName,
				Namespace: instance.Namespace,
			},
			Spec: netv1.NetworkPolicySpec{
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": "nginx",
					},
				},
				PolicyTypes: []netv1.PolicyType {
					"Ingress",
				},
			},
		}
//		fmt.Println(netpo)
		
		err = r.client.Create(context.TODO(), netpo)
//		if err != nil {
//			reqLogger.Info("Error when create", err)
//			return reconcile.Result{}, err
//		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", instance.Namespace, "Pod.Name", found2.Name)
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *actioncontrollerv1alpha1.ActionPolicy) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
