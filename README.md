# fishfishes_backend

### Build and Upload Docker-Image for GCP
``` bash
docker build -t fishfishes .
docker tag fishfishes:latest europe-west3-docker.pkg.dev/tourguide-388412/fishfishes/ff:latest
docker push europe-west3-docker.pkg.dev/tourguide-388412/fishfishes/ff:latest
```
