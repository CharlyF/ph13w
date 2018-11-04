## Upload and Share images

### Status

Under development

### Build

clone the project

dep ensure # Will make the vendor dir
docker build -t gcr.io/<GCP_Project_ID>/thymeline:latest
docker push gcr.io/<GCP_Project_ID>/thymeline:latest
kubectl apply ...

### Requirements

- go: version `go1.10.3`
- dep: version `v0.5.0`
- docker:
```
Client:
 Version:           18.06.1-ce
 API version:       1.38
 Go version:        go1.10.3

Server:
 Engine:
  Version:          18.06.1-ce
  API version:      1.38 (minimum version 1.12)
  Go version:       go1.10.3

```

- kubectl:
```
Client Version: version.Info{Major:"1", Minor:"10", GitVersion:"v1.10.7"[...]}
Server Version: version.Info{Major:"1", Minor:"10+", GitVersion:"v1.10.7-gke.6"[...]}
```

- A cluster in GKE.