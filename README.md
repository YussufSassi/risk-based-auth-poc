# Risk based auth proof of concept

This is a simple proof of concecpt of risk based authentication using fingerprinting. Below is a flowchart that explains this process.

![flowchar](./images/Simple%20risk%20based%20auth%20flow.png)

## Installing

Run `setup.sh` to install the required dependencies and build the frontend application.

Now run `docker compose up -d` to deploy the project using docker.

*Note:* use the username "user" and the password "password" for testing

## Tech stack

- Golang (With gin web framework)
- redis
- Docker
- Svelte
- Typescript
