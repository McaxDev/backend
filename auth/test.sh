cp Dockerfile .. &&
cd .. &&
docker build -t mcaxdev/auth . &&
rm Dockerfile &&
docker-compose -f /srv/axo/auth/docker-compose.yml up -d &&
docker logs -f axo-auth
