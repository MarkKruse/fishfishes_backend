mage build
docker buildx build --platform linux/amd64 -t fishfishes .
docker tag fishfishes:latest europe-west3-docker.pkg.dev/tourguide-388412/fishfishes/ff:latest
docker push europe-west3-docker.pkg.dev/tourguide-388412/fishfishes/ff:latest