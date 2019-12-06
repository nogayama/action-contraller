# action Controller 作業ログ

## operator-sdk のテンプレート

```bash
export GO111MODULE="on"
export GOPATH="${HOME}/go"
# export GOROOT="" # 不要だが，問題が起きたらやってみる価値はある

operator-sdk new action-controller --repo github.com/nogayama/action-controller

cd action-controller

go mod download
# ${HOME}/go にキャッシュがあれば何も表示されない

$ operator-sdk add api --api-version=actioncontroller.k8s.io/v1alpha1 --kind=actionPolicy
    INFO[0000] Generating api version actioncontroller.k8s.io/v1alpha1 for kind actionPolicy. 
    INFO[0000] Created pkg/apis/actioncontroller/group.go 
    INFO[0001] Created pkg/apis/actioncontroller/v1alpha1/actionpolicy_types.go 
    INFO[0001] Created pkg/apis/addtoscheme_actioncontroller_v1alpha1.go 
    INFO[0001] Created pkg/apis/actioncontroller/v1alpha1/register.go 
    INFO[0001] Created pkg/apis/actioncontroller/v1alpha1/doc.go 
    INFO[0001] Created deploy/crds/actioncontroller.k8s.io_v1alpha1_actionpolicy_cr.yaml 
    INFO[0005] Created deploy/crds/actioncontroller.k8s.io_actionpolicies_crd.yaml 
    INFO[0005] Running deepcopy code-generation for Custom Resource group versions: [actioncontroller:[v1alpha1], ] 
    INFO[0014] Code-generation complete.                    
    INFO[0014] Running OpenAPI code-generation for Custom Resource group versions: [actioncontroller:[v1alpha1], ] 
    INFO[0023] Created deploy/crds/actioncontroller.k8s.io_actionpolicies_crd.yaml 
    INFO[0023] Code-generation complete.                    
    INFO[0023] API generation complete.

$ operator-sdk add controller --api-version=actioncontroller.k8s.io/v1alpha1 --kind=actionPolicy
    INFO[0000] Generating controller version actioncontroller.k8s.io/v1alpha1 for kind actionPolicy. 
    INFO[0000] Created pkg/controller/actionpolicy/actionpolicy_controller.go 
    INFO[0000] Created pkg/controller/add_actionpolicy.go 
    INFO[0000] Controller generation complete.
```

`/pkg/apis/actioncontroller/v1alpha1/actionpolicy_types.go` を編集して，カスタムリソースの構造体を修正する



```bash
$ operator-sdk generate k8s
    # actionpolicy_types.goを編集しないなら，何も起きない

$ operator-sdk generate openapi
    # actionpolicy_types.goを編集しないなら，何も起きない

```



### (Optional) 

IKS で動かしたい場合は，github.comにコードを置いてある必要がある？

1. https://github.com/nogayama/action-contraller に作り，
2. add, commit, push


### 実際にうごかす 



```bash
$ kubectl apply -f deploy/crds/actioncontroller.k8s.io_actionpolicies_crd.yaml
$ kubectl apply -f deploy/crds/actioncontroller.k8s.io_v1alpha1_actionpolicy_cr.yaml

$ kubectl get actionpolicies
NAME                        AGE
example-actionpolicy   3s

```



```bash
$ OPERATOR_NAME=foo-controller operator-sdk up local
    INFO[0000] Running the operator locally.                
    INFO[0000] Using namespace default.                     
    {"level":"info","ts":1575336988.608072,"logger":"cmd","msg":"Operator Version: 0.0.1"}
    {"level":"info","ts":1575336988.608135,"logger":"cmd","msg":"Go Version: go1.13.4"}
    {"level":"info","ts":1575336988.608143,"logger":"cmd","msg":"Go OS/Arch: darwin/amd64"}
    {"level":"info","ts":1575336988.6081479,"logger":"cmd","msg":"Version of operator-sdk: v0.12.0"}
    {"level":"info","ts":1575336988.609798,"logger":"leader","msg":"Trying to become the leader."}
    {"level":"info","ts":1575336988.609812,"logger":"leader","msg":"Skipping leader election; not running in a cluster."}
    {"level":"info","ts":1575336989.065544,"logger":"controller-runtime.metrics","msg":"metrics server is starting to listen","addr":"0.0.0.0:8383"}
    {"level":"info","ts":1575336989.06573,"logger":"cmd","msg":"Registering Components."}
    {"level":"info","ts":1575336989.065954,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"actionpolicy-controller","source":"kind source: /, Kind="}
    {"level":"info","ts":1575336989.066093,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"actionpolicy-controller","source":"kind source: /, Kind="}
    {"level":"info","ts":1575336989.0662131,"logger":"cmd","msg":"Could not generate and serve custom resource metrics","error":"operator run mode forced to local"}
    {"level":"info","ts":1575336989.53053,"logger":"metrics","msg":"Skipping metrics Service creation; not running in a cluster."}
    {"level":"info","ts":1575336989.989225,"logger":"cmd","msg":"Could not create ServiceMonitor object","error":"no ServiceMonitor registered with the API"}
    {"level":"info","ts":1575336989.9892569,"logger":"cmd","msg":"Install prometheus-operator in your cluster to create ServiceMonitor objects","error":"no ServiceMonitor registered with the API"}
    {"level":"info","ts":1575336989.9892638,"logger":"cmd","msg":"Starting the Cmd."}
    {"level":"info","ts":1575336989.989557,"logger":"controller-runtime.manager","msg":"starting metrics server","path":"/metrics"}
    {"level":"info","ts":1575336990.0936162,"logger":"controller-runtime.controller","msg":"Starting Controller","controller":"actionpolicy-controller"}
    {"level":"info","ts":1575336990.197492,"logger":"controller-runtime.controller","msg":"Starting workers","controller":"actionpolicy-controller","worker count":1}
    {"level":"info","ts":1575336990.197728,"logger":"controller_actionpolicy","msg":"Reconciling actionPolicy","Request.Namespace":"default","Request.Name":"example-actionpolicy"}
    {"level":"info","ts":1575336990.1978939,"logger":"controller_actionpolicy","msg":"Creating a new Pod","Request.Namespace":"default","Request.Name":"example-actionpolicy","Pod.Namespace":"default","Pod.Name":"example-actionpolicy-pod"}
    {"level":"info","ts":1575336990.230028,"logger":"controller_actionpolicy","msg":"Reconciling actionPolicy","Request.Namespace":"default","Request.Name":"example-actionpolicy"}
    {"level":"info","ts":1575336990.230135,"logger":"controller_actionpolicy","msg":"Skip reconcile: Pod already exists","Request.Namespace":"default","Request.Name":"example-actionpolicy","Pod.Namespace":"default","Pod.Name":"example-actionpolicy-pod"}
    {"level":"info","ts":1575336990.242269,"logger":"controller_actionpolicy","msg":"Reconciling actionPolicy","Request.Namespace":"default","Request.Name":"example-actionpolicy"}
    {"level":"info","ts":1575336990.2423182,"logger":"controller_actionpolicy","msg":"Skip reconcile: Pod already exists","Request.Namespace":"default","Request.Name":"example-actionpolicy","Pod.Namespace":"default","Pod.Name":"example-actionpolicy-pod"}
    {"level":"info","ts":1575336990.270487,"logger":"controller_actionpolicy","msg":"Reconciling actionPolicy","Request.Namespace":"default","Request.Name":"example-actionpolicy"}
    {"level":"info","ts":1575336990.270541,"logger":"controller_actionpolicy","msg":"Skip reconcile: Pod already exists","Request.Namespace":"default","Request.Name":"example-actionpolicy","Pod.Namespace":"default","Pod.Name":"example-actionpolicy-pod"}
    {"level":"info","ts":1575336997.700701,"logger":"controller_actionpolicy","msg":"Reconciling actionPolicy","Request.Namespace":"default","Request.Name":"example-actionpolicy"}
    {"level":"info","ts":1575336997.700744,"logger":"controller_actionpolicy","msg":"Skip reconcile: Pod already exists","Request.Namespace":"default","Request.Name":"example-actionpolicy","Pod.Namespace":"default","Pod.Name":"example-actionpolicy-pod"}

```

```bash
# 別のターミナルでみると，カスタムリソースの名前+"-pod" が存在しない場合，作成する．
$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
example-actionpolicy-pod   1/1     Running   0          4m24s
```



### 調査

```yaml
apiVersion: networking.k8s.io/v1
```

https://github.com/kubernetes/kubernetes

https://github.com/kubernetes/kubernetes/blob/3ea3cfc3bea08c425593a25533bf14c703f90420/staging/src/k8s.io/client-go/kubernetes/typed/networking/v1/networkpolicy.go



```go
import (
	"time"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)
...
```


トリエアズこれは動いた
```go
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
```



### 編集



### テーブル表示に新たな列を足す

```yaml
# deploy/crds/actioncontroller.k8s.io_actionpolicies_crd.yaml
...
spec:
  ...
  additionalPrinterColumns:
  - name: AGE
    type: date
    JSONPath: .metadata.creationTimestamp
  - name: size
    type: string
    description: message content which want to show
    JSONPath: .spec.size

```



```bash

$ kubectl apply -f deploy/crds/actioncontroller.k8s.io_actionpolicies_crd.yaml

$ kubectl get actionpolicies
NAME                        AGE   SIZE
example-actionpolicy   20m   3

```



### 実際のフィールドに変更

```yaml
# deploy/crds/actioncontroller.k8s.io_actionpolicies_crd.yaml
  additionalPrinterColumns:
  - name: AGE
    type: date
    JSONPath: .metadata.creationTimestamp
  - name: targetType
    type: string
    description: message content which want to show
    JSONPath: .spec.targetType
  - name: target
    type: string
    description: message content which want to show
    JSONPath: .spec.target
  - name: action
    type: string
    description: message content which want to show
    JSONPath: .spec.action
    
```

```yaml
# deploy/crds/actioncontroller.k8s.io_v1alpha1_actionpolicy_cr.yaml
apiVersion: actioncontroller.k8s.io/v1alpha1
kind: actionPolicy
metadata:
  name: actionpolicy-for-abcbankweb
spec:
  targetType: deployment
  target: abcbankweb
  action: auto-update

```



### 対象の構造体を編集



```go
# pkg/apis/actioncontroller/v1alpha1/actionpolicy_types.go

type actionPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	
	// +kubebuilder:validation:Format=string
	Target string `json:"target"`
	
	// +kubebuilder:validation:Format=string
	TargetType string `json:"targetType"`
	
	// +kubebuilder:validation:Format=string
	Action string `json:"action"`
}
```





```bash
$ operator-sdk generate k8s
    # pkg/apis/actioncontroller/v1alpha1/zz_generated.deepcopy.go が変更される

$ operator-sdk generate openapi
    # pkg/apis/actioncontroller/v1alpha1/zz_generated.openapi.go が変更される

```

