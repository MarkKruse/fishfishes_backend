# fishfishes_backend

## Upload Docker-Image
https://cloud.google.com/build/docs/build-push-docker-image?hl=de


MongoDB PW
mongodb+srv://markkruse92:SZ9WuHjDrr4KgcoP@fishfishescluster.byjsxja.mongodb.net/?retryWrites=true&w=majority
markkruse92
SZ9WuHjDrr4KgcoP

Update to GCP

einmal ausführen: gcloud auth configure-docker europe-west3-docker.pkg.dev

europe-west3-docker.pkg.dev/tourguide-388412/fishfishes



docker build -t fishfishes .
docker tag fishfishes:latest europe-west3-docker.pkg.dev/tourguide-388412/fishfishes/ff:latest
docker push europe-west3-docker.pkg.dev/tourguide-388412/fishfishes/ff:latest

