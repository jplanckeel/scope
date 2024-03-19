<p align="center" style="margin-top: 120px">

  <h3 align="center">Scope</h3>

  <p align="center">
  <img src="https://cdn.rawgit.com/jplanckeel/scope/main/images/banner.png" style="width:66%" alt="Scope">
  </p>
  <p align="center">
    An Open-Source cli to Sync (Helm)Chart On Private OCI Registry
    <br />
  </p>
</p>
<p align="center">
  <a href="https://github.com/jplanckeel/scope/releases"><img title="Release" src="https://img.shields.io/github/v/release/jplanckeel/scope"/></a>
  <a href=""><img title="Downloads" src="https://img.shields.io/github/downloads/jplanckeel/scope/total.svg"/></a>
  <a href=""><img title="Docker pulls" src="https://img.shields.io/docker/pulls/jplanckeel/scope"/></a>
  <a href=""><img title="Go version" src="https://img.shields.io/github/go-mod/go-version/jplanckeel/scope"/></a>
  <a href=""><img title="Docker builds" src="https://img.shields.io/docker/automated/jplanckeel/scope"/></a>
  <a href=""><img title="Code builds" src="https://img.shields.io/github/actions/workflow/status/jplanckeel/scope/build.yml"/></a>
  <a href=""><img title="apache licence" src="https://img.shields.io/badge/License-Apache-yellow.svg"/></a>
  <a href="https://github.com/jplanckeel/scope/releases"><img title="Release date" src="https://img.shields.io/github/release-date/jplanckeel/scope"/></a>
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
- [x] Add CI example for Gitlab
- [ ] Add CI example Github Action

## Regsitry Tested 

| Resgistry         | Tested     | Scope Version |
|--------------|-----------|------------|
| AWS ECR | (/)      | v0.0        |
| Nexus     | (/) | v0.2      |
| Docker Hub     | (/) | v0.4      |
| Scaleway Registry     | (/) | v0.4      |

## requierment 

- Regsitry with OCI compatibility
- Repository created on regsitry with named format :  {{ registry }}/{{ namespace }}

example : rg.fr-par.scw.cloud/helm-mirrors

## cli 

```bash
> scope -h
Sync Chart On Private Registry

Usage:
  scope [flags]

Flags:
      --append-source              append repository to chart name
      --ca-file string             verify certificates of HTTPS-enabled servers using this CA bundle
      --cert-file string           identify HTTPS client using this SSL certificate file
  -h, --help                       help for scope
      --insecure-skip-tls-verify   skip tls certificate checks
      --key-file string            identify HTTPS client using this SSL key file
  -n, --namespace string           namespace destination chart registry
  -p, --password string            chart destination repository password
      --password-stdin             read password or identity token from stdin
  -r, --registry string            destination chart registry
  -s, --source-file string         path to source file
  -t, --type string                type for registry (nexus or oci) (default "oci")
  -u, --user string                chart destination repository user
```

example : 

```bash
scope -s config.yaml -r 000000000000.dkr.ecr.eu-west-3.amazonaws.com
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

## Docker Image

You can find a docker image with Helm and Scope cli here :

https://hub.docker.com/r/jplanckeel/scope

```bash
docker run jplanckeel/scope scope -h                                                                               
```                                                                                                                                     

## CI

### Gitlab-ci

```bash
sync-charts:
  tags:
    - docker
  image: jplanckeel/scope
  stage: sync
  script: scope -s ./scope_config.yml -t nexus -u $REGISTRY_USER -p $REGISTRY_USER_TOKEN -r https://docker.nexus-jplanckeel.com
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
      when: manual
    - if: $CI_PIPELINE_SOURCE == "schedule"
      when: always
```

## Renovate Configuration

You can update file with renovate, an example to package rules custom :

INFO : We have a limitation with renovate in this example, we can use mutliple version for a same chart.

```json
"packageRules": [
      {
        "customType": "regex",
        "fileMatch": ["scope_config.yml$"],
        "matchStrings": [
          "(?<registryUrl>.*):\\s+charts:\\s+((?<depName>.*):\\s+-\\s(?<currentValue>[\\w+\\.]*))"
        ],
        "datasourceTemplate": "helm"
      }
    ]
```

<!> In this example you need to specify protocole to regsitry url.

```yaml
https://stakater.github.io/stakater-charts:
  charts:
    reloader:
    - 1.0.67
```

## build 

```bash
 go build -ldflags "-s -w" -o bin/scope 
```
