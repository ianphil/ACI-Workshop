# AKS-ACI-Workshop
Workshop to setup AKS with ACI offload

```docker run --network="app" -p 27017:27017 --name some-mongo -d mongo:4.0-xenial```

```docker run --network="app" -d -p 5672:5672 --hostname my-rabbit --name some-rabbit rabbitmq:3```

```docker exec -it some-mongo bash```

```mongo```

```use DATABASE_NAME```

```db.events.find().pretty()```

```docker network create app```

```docker run --rm -it --network="app" tripdubroot/aaw-show:0.0.2```

```docker run --rm -it --network="app" tripdubroot/aaw-receive:0.0.2```

```docker run --rm -it --network="app" -p 8080:8080 tripdubroot/aaw-show:0.0.2```