APP_NAME='account'

go build -o app . &&
docker build -t mcaxdev/${APP_NAME} . &&
docker-compose -f /srv/axo/${APP_NAME}/docker-compose.yml up -d &&
docker logs -f axo-${APP_NAME}
