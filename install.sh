#!/bin/bash

# Default port
PORT=${1:-8080}

# Download and install binary
echo "Downloading systemd-healthcheck binary..."
wget -O /usr/local/bin/systemd-healthcheck https://github.com/Saalin/systemd-healthcheck/releases/latest/download/systemd-healthcheck
chmod +x /usr/local/bin/systemd-healthcheck

# Create systemd service
echo "Creating systemd service..."
cat <<EOF | sudo tee /etc/systemd/system/systemd-healthcheck.service > /dev/null
[Unit]
Description=Systemd Healthcheck Service
After=network.target

[Service]
ExecStart=/usr/local/bin/systemd-healthcheck -port=$PORT
Restart=always
User=nobody
Group=nobody

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd, enable and start service
echo "Reloading systemd and starting systemd-healthcheck service..."
sudo systemctl daemon-reload
sudo systemctl enable systemd-healthcheck
sudo systemctl start systemd-healthcheck

echo "Installation complete! Service running on port $PORT"
