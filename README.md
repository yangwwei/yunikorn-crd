# yunikorn-crd

CRD code generation:

```shell script
# this is just a example
github.com/kubernetes/code-generator/generate-groups.sh all github.com/apache/incubator-yunikorn-k8shim/pkg/client github.com/apache/incubator-yunikorn-k8shim/pkg/apis yunikorn:v1alpha1 -v 5
```

Please refer to [code-generator](https://github.com/kubernetes/code-generator), and [deep-deive-code-generation](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/) for details.
