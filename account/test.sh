cp Dockerfile .. &&
cd .. &&
docker build -t mcaxdev/account . &&
rm Dockerfile &&
docker-compose -f /srv/axo/account/docker-compose.yml up -d &&
docker logs -f axo-account
