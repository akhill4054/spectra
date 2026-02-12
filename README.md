# Spectra

Spectra is an **in-progress, event-driven, privacy-preserving backend
system** for computing real-time audience metrics for digital signage
using camera-derived observation events from edge devices.

---

## Project Status

🚧 **Actively in progress**

Spectra is currently being built as a learning exercise focused on
streaming-style data processing in Go. It may evolve into a full-fledged
event-driven analytics engine for real-world signage systems over time.

---

## Roadmap

See the full roadmap in [`ROADMAP.md`](ROADMAP.md).

---

## Non-Goals

-   Face recognition or biometric identification
-   Image or video storage
-   Frontend dashboards
-   Model training or ML experimentation

---

## Running the Project

You can start the Spectra service either directly via Go/Make, or using
Docker Compose.

### Using Make (local Go toolchain)

Prerequisites:

-   Go 1.22+ installed locally

Common commands:

-   **Build**: `make build`
-   **Run**: `make run`
-   **Test**: `make test`

By default the service listens on port `8080`. You can customize the
port in three ways (in order of precedence):

1.  **CLI flag**: `make run ARGS="--p=9090"`
2.  **Environment variable**: `PORT=9090 make run`
3.  **Default**: falls back to `8080`

The health endpoint is available at `http://localhost:<port>/healthz`.

### Using Docker Compose

Prerequisites:

-   Docker and Docker Compose installed

Create or edit the `.env` file in the project root (an example is
already present):

``` bash
PORT=8080
```

Then start the service:

``` bash
docker-compose up
```

Docker Compose will:

-   Mount the local project into the container
-   Use `PORT` from `.env` (defaulting to `8080` if not set)
-   Run `go run ./cmd/spectra`

Access the health endpoint at:

-   `http://localhost:8080/healthz` (or your configured `PORT`)

---

## License

MIT
