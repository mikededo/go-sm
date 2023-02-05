# Server

The server is written in go and following the hexagonal architecture. It uses
gin as the REST framework.

## Set up

1. Clone the repository and cd into the `/server` directory:

```shell
git clone https://github.com/mikededo/go-sm.git --depth 1
cd go-sm/server
```

2. Create an `.env` file in the root of the `/server` directory:

```shell
touch .env
```

3. Fill the `.env` file with the following fields

```shell
JWT_SIGN_KEY=<jwt signing key>
```

4. Run `make help` for a list of the available commands.
