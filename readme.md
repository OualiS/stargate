# Stargate

🚪 A minimal, hot-reloadable reverse proxy written in Go.  
🔁 Route traffic based on host and path rules using a YAML config.

---

## Features

- 🧠 Smart host/path routing
- ♻️ Hot-reload config on change
- 🐳 Docker-ready

---

## TO-DO

### 🔧 Features
- [x] Load configuration from YAML
- [x] Host and path prefix matching
- [x] Basic reverse proxy using `httputil`
- [x] Live config reload with `fsnotify`
- [ ] Improved proxy error handling (logs + proper status)
- [ ] HTTPS support (manual or auto TLS via ACME)
- [ ] Web dashboard to visualize current routes
- [ ] CLI tool (`stargate`) to run / reload / validate config

---

### 🧪 Testing & Quality
- [ ] Unit tests for routing logic
- [ ] Integration tests with simulated backend services
- [ ] Stricter config validation (YAML schema or rules)

---

### 📦 Deployment & Developer Experience
- [x] Dockerfile
- [ ] Multi-arch Docker image (ARM / x86)
- [ ] GitHub Actions CI for automatic builds
- [ ] Publish image to Docker Hub

---

### 💡 Future Ideas
- [ ] Custom middleware support (auth, rewrite, caching, etc.)
- [ ] WebSocket proxy support
- [ ] Support for config via JSON or ENV variables
- [ ] Maintenance mode toggle per route or globally

## Quick Start

```bash
docker compose up --build
