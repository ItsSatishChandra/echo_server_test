# echo_server_test

This repository contains a small Go-based echo server implementation with both TCP and UDP protocols and helper test tooling.

The project provides:

- A TCP echo server (`server/tcp/tcp_server.go`).
- A UDP echo server (`server/udp/udp_server.go`).
- A test server and example entrypoint in `cmd/echo_test_server`.
- Utilities for accepting connections in `internal/connection_listener.go`.
- Docker and Docker Compose files for containerized runs: `Dockerfile`, `compose.yaml`, and `compose.debug.yaml`.

**Quick Overview**

- **TCP port:** 20001
- **UDP port:** 20002

**Build (native)**

Install Go (1.20+ recommended), then build or run the server directly.

- Build the main server binary (example using `cmd/Echo`):

  `go build -o bin/echo cmd/Echo/echo.go`

- Run the test server directly:

  `go run ./cmd/echo_test_server`

**Run with Docker (recommended: use included PowerShell helper)**

Build the container image locally, then use the included PowerShell helper `docker_run_powershell.ps1` to run the container with sane defaults and a mounted `data/` folder.

1) Build the image (one-time):

```powershell
docker build -t echoservertest .
```

2) Run the helper script (from the repository root) — this script creates a `data/` directory if missing, runs the container detached, and prints the assigned host ports and container id:

```powershell
.\docker_run_powershell.ps1
```

The helper reads environment variables to override ports if needed. Example (set environment variables for the current session before running):

```powershell
$env:TCP_PORT = 30001; $env:UDP_PORT = 30002; .\docker_run_powershell.ps1
```

Notes about the helper script:

- The script mounts the repository `data/` directory into the container at `/app/data`.
- It runs the container detached and uses `docker port` to show which host ports were assigned (the script maps host ports dynamically to avoid conflicts).
- To see live logs after the script prints the container id, run `docker logs -f <container_id>`.

You can still use Docker Compose directly if you prefer:

```powershell
docker compose up --build
```

**Usage Examples**

TCP echo test (uses `nc` / netcat):

```powershell
# Connect to TCP echo server
nc localhost 20001
# Type a line and the server will echo it back
```

UDP echo test (use `nc -u` or `socat`):

```powershell
# Using netcat (some builds support -u for UDP)
ncat -u -w1 localhost 20002
```

If you run the container with the volumes mapped to `data/`, the server can read or write files there depending on configuration.


**Project Layout**

- `cmd/Echo`  primary server entrypoint (example)
- `cmd/echo_test_server`  helper/test server executable
- `server/tcp`  TCP server implementation
- `server/udp`  UDP server implementation
- `internal`  shared utilities (listener, logger)
- `Dockerfile`, `compose.yaml`, `compose.debug.yaml`  container and compose configs

**Notes & Troubleshooting**

- If ports `20001` or `20002` are in use, change the published ports in `docker run` or `compose.yaml`.
- On Windows, ensure `nc`/`socat` are available or use a small Go client to exercise the servers.
- Use `docker logs <container>` to inspect server startup output.

If you'd like, I can also:

- Add example client programs in `cmd/clients` for TCP and UDP.
- Add health checks and a README section showing log output and expected behavior.

---
Generated README based on the repository layout and existing server implementations.
