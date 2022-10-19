# Email Indexer

**`email-indexer`** is a fullstack application that can index and search emails details and content within seconds.

It has a monorepo structure and was developed following some principles of hexagonal architecture.

|                       Web version                       |                        Mobile version                         |
| :-----------------------------------------------------: | :-----------------------------------------------------------: |
| ![Deepsearch-web](./images/deepsearch-web.png?raw=true) | ![Deepsearch-mobile](./images/deepsearch-mobile.png?raw=true) |

## Tech Stack

- Backend: [Go](https://go.dev/)
- API Router: [Chi](https://github.com/go-chi/chi)
- Search Engine: [Zinc](https://github.com/zinclabs/zinc)
- Frontend: [Vue](https://vuejs.org/) + [TypeScipt](https://www.typescriptlang.org/)
- Tooling: [Docker](https://docs.docker.com/)

## Instalation

- Download the data with the `scripts/etl.sh` shell
  script

```sh
    chmod +x scripts/etl.sh
    sh scripts/etl.sh
```

- Build the docker images and start the container

```sh
    make start
```

### See the logs of the containers

```sh
    make logs
```

### Stop the containers

```bash
    make stop
```

### Clean the containers

```bash
    make clean
```
