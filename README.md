# DEV ENVIRONMENT SETUP:

## LOCAL SETUP (on your machine)

Regardless of OS, dowload ```air``` with:

```
go install github.com/air-verse/air@latest
```
Make sure that your ```GOBIN``` path is set properly, so the ```air``` command can be found.

Make sure you also have ```make``` installed.

Run the API with:

```
make run-local
```

This will detect your OS and run ```air``` with the proper ```.toml``` configuration so you can have hot reload on your local machine.

For local runs you'll need a ```.env.local``` environment file with all the keys described in ```.env.dist``` and ```AIS_ENV``` set to ```local```.

## DOCKER COMPOSE

Make sure you have ```docker``` installed.

Start the API container with:

```
make docker-start
```

This will start a container with ```air``` pre-installed and create a volume attached to your local project folder so hot-reload can
detect changes from your local machine and trigger a rebuild of the project inside the container.

To stop the project and remove the attached volume run:

```
make docker-stop
```

For containerized runs you'll need a ```.env.dev``` environment file with all the keys described in ```.env.dist``` and ```AIS_ENV``` set to ```development```.


## OTHER ENVIRONMENTS

There are two more ```AIS_ENV``` types of environment that are supported, ```staging``` and ```production```. These environments assume that the application is deployed on a provider and so you'll need to setup the environment variables according to the provider specification. In any environment other tha ```local```, the app is reading the variables directly from the OS (bare metal, container, or vm). For containers, use those environment types when running your image with the production ```Dockerfile``` that has multistage builds enabled, and does not use ```air```. 

- To build the production or staging image run ```make build```. 

Once you publish your image to a registry, you'll be able to deploy it with the environment variables of your choice.

### Note:

Nothing stops you to pass a different ```AIS_ENV``` value to the production container. However, if set to ```production``` it will serve https from :443 
and you will need to install a ```cert.pem``` and ```cert.key```. ```AIS_ENV=local``` reads the variables from a ```.env``` file and not from the OS.
