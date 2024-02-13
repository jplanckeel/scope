<p align="center" style="margin-top: 120px">

  <h3 align="center">Scope</h3>

  <p align="center">
    An Open-Source cli to Sync (Helm)Chart On Private OCI Registry
    <br />
  </p>
</p>



## About Scope 
 
Synchronize Your Helm Charts Seamlessly with OCI-Compatible Registries

In the ever-evolving landscape of container orchestration, Helm diagrams have become a cornerstone for simplifying application deployment and management. To enhance this experience, we present Scope, a simple and effective tool for effortlessly synchronizing Helm diagrams with OCI-compliant registries such as Nexus or ECR (Elastic Container Registry). This allows you to maintain control over the sources you use in a simple YAML

### Key Features:

* Seamless Integration: Scope seamlessly integrates with Helm charts, providing a smooth and efficient synchronization process with OCI-compatible registries.

* OCI Compatibility: Embracing the Open Container Initiative (OCI) standards, ensures that your Helm charts align perfectly with OCI-compatible registries, promoting interoperability and industry standards.

* Multi-Registry Support: Scope doesn't limit you to a single registry. Whether you prefer Nexus or ECR, the tool supports multiple OCI-compatible registries, offering flexibility and choice.

* Automated Sync: Say goodbye to manual interventions. Scope automates the synchronization process, ensuring that your Helm charts are always up-to-date in the designated OCI-compatible registry.

* Version Control: Manage and track versions effortlessly. Scope supports versioning, allowing you to synchronize specific chart versions with precision.

*  Logging: Gain insights into synchronization activities with detailed logging. Scope provides comprehensive logs, making it easy to troubleshoot and monitor the synchronization process.

## Roadmap 

- [ ] Add regex usage for version
- [ ] Add CI example for Gitlab and Github Action

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
  -b, --binary string          alias for binary helm3 (default "helm")
  -c, --config string          path to configfile
  -d, --dryrun                 enable dry-run mode
  -h, --help                   help for scope
  -p, --password string        password for nexus registry
  -r, --registry string        destination chart registry
  -t, --registry-type string   registry nexus or ecr (default: oci) (default "oci")
  -u, --user string            user for nexus registry
  -v, --version                version for scope

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


## build 

```bash
 go build -ldflags "-s -w"
```
