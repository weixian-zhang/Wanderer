kubectl delete istiooperator istio-control-plane -n istio-system
kubectl delete crd istiooperators.install.istio.io
istioctl operator remove
kubectl delete -f istio-operator-demo.yaml
kubectl delete ns istio-system --grace-period=0 --force
kubectl delete ns istio-operator --grace-period=0 --force

REM remove finalizer if crd deletion stuck. usually happens when remove namespace istio-system
REM before removing crd istiooperators.install.istio.io
REM kubectl get istiooperator -n istio-system istio-control-plane -o=json | jq '.metadata.finalizers = null' | kubectl apply -f -