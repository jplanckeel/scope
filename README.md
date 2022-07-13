# scope

cli to sync Helmchart on private registry with oci compatibility

## build 

```bash
 go build -ldflags "-s -w"
```
## requierment 

- Regsitry with OCI compatibility
- Repository created on regsitry with named format :  helm-mirrors/{{ registry }}/{{chart}}

example : helm-mirrors/aws.github.io/eks-charts/aws-load-balancer-controller

## cli 

```bash
> scope -h
a cli to sync helmchart to private registry

Usage:
  scope [flags]

Flags:
  -b, --binary string     alias for binary helm3 (default "helm")
  -c, --config string     path to configfile
  -d, --dryrun            enable dry-run mode
  -h, --help              help for scope
  -r, --registry string   destination chart registry
  -v, --version           version for scope

```

example : 

```bash
scope -c config.yaml -r 000000000000.dkr.ecr.eu-west-3.amazonaws.com
```

## Configuration example 

```yaml 
## example.yaml
apache.github.io/superset:
  charts:
    superset:
    - 0.1.0
    - 0.1.1
prometheus-community.github.io/helm-charts:
  charts:
    prometheus:
    - ~11.1.0
    prometheus-node-exporter:
    - 2.0.0
    - 2.0.1


```
