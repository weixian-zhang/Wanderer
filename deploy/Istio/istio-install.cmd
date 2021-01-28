istioctl operator init
kubectl create ns istio-system
kubectl apply -f istio-operator-profile-demo.yaml
kubectl get all -n istio-system
kubectl logs -n istio-operator -l name=istio-operator -f
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.8/samples/addons/extras/zipkin.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.8/samples/addons/prometheus.yaml