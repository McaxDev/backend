cp Dockerfile .. &&
cd .. &&
docker build -t mcaxdev/wiki . &&
rm Dockerfile &&
docker-compose -f /srv/axo/wiki/docker-compose.yml up -d &&
docker logs -f axo-wiki
