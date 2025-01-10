cp Dockerfile .. &&
cd .. &&
docker build -t mcaxdev/gallery . &&
rm Dockerfile &&
docker-compose -f /srv/axo/gallery/docker-compose.yml up -d &&
docker logs -f axo-gallery
