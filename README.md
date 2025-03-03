# Systemd Healthcheck

[![Build](https://github.com/Saalin/systemd-healthcheck/actions/workflows/build.yml/badge.svg)](https://github.com/Saalin/systemd-healthcheck/actions/workflows/build.yml)

A lightweight HTTP server that exposes the status of systemd services as a health check endpoint.

## Features
- ✅ Query multiple systemd services via an HTTP API
- ✅ Returns HTTP `200 OK` if all services are active, `503 Service Unavailable` if any are down
- ✅ Lightweight, single binary with no dependencies
- ✅ Runs as a systemd service for automatic startup
- ✅ Simple installation script for quick setup
- ✅ Configurable port via command-line argument (default: `8080`)

---

## Installation

### **1. Download & Install**
Clone the repository and build the binary:
```sh
git clone https://github.com/Saalin/systemd-healthcheck.git
cd systemd-healthcheck
go build -o systemd-healthcheck
```

Or download a prebuilt binary:
```sh
wget -O systemd-healthcheck https://github.com/Saalin/systemd-healthcheck/releases/latest/download/systemd-healthcheck
chmod +x systemd-healthcheck
```

### **2. Run the Binary**
```sh
./healthcheck -port=8080
```
By default, it starts on **port 8080**.

---

## Usage

### **Check Multiple Services**
Query multiple services using a comma-separated list in the `services` query parameter.

> **Note:** The `systemd-healthcheck` service uses `systemctl status` to check the status of the specified services.

#### ✅ **Example: All services running (`200 OK`)**
```sh
curl "http://localhost:8080/health?services=nginx,sshd"
```
**Response:**
```json
{
  "services": {
    "nginx": true,
    "sshd": true
  }
}
```

#### ❌ **Example: A service is down (`503 Service Unavailable`)**
```sh
curl "http://localhost:8080/health?services=nginx,nonexistent"
```
**Response:**
```json
{
  "services": {
    "nginx": true,
    "nonexistent": false
  }
}
```

### **HTTP Status Codes**
- **`200 OK`** → All services are active
- **`503 Service Unavailable`** → At least one service is inactive or does not exist
- **`400 Bad Request`** → No `services` parameter provided

---

## Automatic Installation as a Systemd Service

### **1. Install the Service**
Run the installation script with an optional port argument (default: `8080`):
```sh
wget -O install.sh https://raw.githubusercontent.com/Saalin/systemd-healthcheck/main/install.sh
chmod +x install.sh
sudo ./install.sh 9090  # Change 9090 to desired port
```

### **2. Enable & Start Service**
```sh
sudo systemctl enable systemd-healthcheck
sudo systemctl start systemd-healthcheck
```

### **3. Check Status**
```sh
sudo systemctl status systemd-healthcheck
```

---

## Manual Systemd Setup
Alternatively, you can manually create the systemd service file:
```sh
sudo cp systemd-healthcheck /usr/local/bin/
sudo nano /etc/systemd/system/systemd-healthcheck.service
```
Paste the following:
```ini
[Unit]
Description=Systemd Healthcheck Service
After=network.target

[Service]
ExecStart=/usr/local/bin/systemd-healthcheck -port=8080
Restart=always
User=nobody
Group=nobody

[Install]
WantedBy=multi-user.target
```
Then reload systemd and start the service:
```sh
sudo systemctl daemon-reload
sudo systemctl enable systemd-healthcheck
sudo systemctl start systemd-healthcheck
```

---

## Configuration
You can modify the listening port by passing the `-port` argument:
```sh
./systemd-healthcheck -port=9090
```
Or modify the systemd service file:
```ini
ExecStart=/usr/local/bin/systemd-healthcheck -port=9090
```

---

## Troubleshooting

### **Check Logs**
If the service is not working, check the logs:
```sh
journalctl -u systemd-healthcheck -f
```

### **Restart the Service**
```sh
sudo systemctl restart systemd-healthcheck
```

### **Verify Binary Works Manually**
If systemd is failing, try running the binary manually:
```sh
/usr/local/bin/systemd-healthcheck -port=8080
```

## License
MIT License
