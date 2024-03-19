## Sync
### AWS ECR

```bash

```

### Scaleway Regsitry

```bash
scope -s example.yaml -u nologin -n <NAMESPACE> -r rg.fr-par.scw.cloud --password-stdin <<< "$SCW_SECRET_KEY"
```

### Docker Hub

```bash
scope -s example.yaml -u <USER> -p <PASSWORD> -n <NAMESPACE> -r registry-1.docker.io
```


### Nexus Registry

```bash
scope -s example.yaml -t nexus -u <USER> -p <PASSWORD> -r https://nexus.registry.com
```
