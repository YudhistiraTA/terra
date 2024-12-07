# Terra

This is a simple application which only has a login feature, a note list for logged in user, and the ability to create, update, and delete the notes.

## How to run

### Docker deployment

A `docker-compose.yml` file is provided. The docker compose will run a **postgres instance, a go backend service, and a nextjs frontend**.

Please refer to the [this documentation](https://docs.docker.com/compose/install/) on how to **install docker compose**. Once it is installed, run the following command to build the images and run them.

```
docker compose up --build
```

A user should already be inserted into the database with the following credentials:
```
email: example@mail.com
password: p@ssw0rd1234
```

#### Important notes

The docker containers ran expose the following ports to the host machine for easy debugging during development:

1. `:3000` for the frontend client
2. `:8000` for the backend service
3. `:5432` for the postgres instance

If any of these ports are already used, **most commonly due to an already existing postgres service running in the host machine**, the docker compose may fail. In which case, please **modify the exposed ports** as needed in the `docker-compose.yml` file.

### Manual development run

#### Frontend Client

Make sure to **install the dependencies** with your preferred package manager. a `pnpm-lock.yaml` file is provided as it is the package manager that was used during development.

Create a `.env` file by following the provided `.env.example` file, and then use the `dev` script with your chosen package manager to start the nextjs client like so:

```
pnpm run dev
```

#### Database Setup

This project uses postgres as the database. You may **initialize the tables** using the schema located at `./backend/internal/infrastructure/db/sqlc/schema`.

It is important to note that there is **no register feature** in the application, as such, you must **manually insert a user** into the `users` table. The password should be **encrypted using bcrypt** before insertion.

#### Backend Server

Please create a `.env` file by following the `.env.example` file provided.

This project uses [air](https://github.com/air-verse/air) for hot reloading. Please refer to its documentation on how to **install air**. An `.air.toml` file is already provided, as such **initializing it again won't be necessary**. You may simply run the following command to start the server:

```
air
```

