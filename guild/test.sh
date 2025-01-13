cp Dockerfile .. &&
cd .. &&
docker build -t mcaxdev/guild . &&
rm Dockerfile &&
docker-compose -f /srv/axo/guild/docker-compose.yml up -d &&
docker logs -f axo-guild
