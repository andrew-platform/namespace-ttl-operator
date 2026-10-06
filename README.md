# Namespace TTL Operator

A Kubernetes operator that creates namespaces with a time-to-live (TTL). 
When the TTL expires, the operator automatically deletes the namespace.

Useful for preview environments, test sandboxes, and cost governance.

## How It Works

1. Apply a `TemporaryNamespace` resource with a `ttl` field
2. The operator creates a real namespace
3. When the TTL expires, the operator deletes it

## Example

```yaml
apiVersion: infra.example.com/v1alpha1
kind: TemporaryNamespace
metadata:
  name: demo-sandbox
spec:
  ttl: "30m"
```

## Running Locally

```bash
make install
make run
kubectl apply -f config/samples/infra_v1alpha1_temporarynamespace.yaml
```

## Built With

- Go
- Kubebuilder
- controller-runtime
