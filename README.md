# ctrlPad daemon 

![GitHub stars](https://img.shields.io/github/stars/ctrlPad/firmware?style=for-the-badge&logo=github) ![GitHub issues](https://img.shields.io/github/issues/ctrlPad/firmware?style=for-the-badge&logo=github) ![Last commit](https://img.shields.io/github/last-commit/ctrlPad/firmware?style=for-the-badge&logo=github)

The daemon is a background service that listens for and executes CtrlPad button actions.

> [!WARNING]  
> This repository and the daemon currently support Linux only. Windows and macOS are not supported at this time.


## Table of Content

- [Quick Start](#quick-start)
- [Development Setup](#development-setup)
- [Contributing](#contributing)

## Quick Start

Download and install the binaries:

```bash
# macOS / Linux
curl -f https://raw.githubusercontent.com/${OWNER}/${REPO}/main/install.sh | sudo sh


# Windows / macOS
Soon...
```

After that run the following command, if you want to run always the daemon: 
```bash
# Create the service file
cat <<EOF > /etc/systemd/system/ctrlpad-daemon.service
[Unit]
Description=CtrlPad Daemon
After=network.target

[Service]
ExecStart=/usr/local/bin/ctrlpad-daemon
Restart=on-failure
User=root

EOF

# Reload, start, and enable the service
systemctl daemon-reload
systemctl start ctrlpad-daemon.service
systemctl enable ctrlpad-daemon.service
```

Verify the installation:
```bash
systemctl status ctrlpad-daemon.service
```

## Development Setup

To contribute to this project, please ensure you have [devenv](https://deven.sh) installed on your system.

```
# Clone the repository
git clone https://github.com/CtrlPad/desktop.git
cd desktop

# Enter the development environment
devenv shell
```

## Contributing

1. **Fork** the repository
2. **Clone** your fork: `git clone https://github.com/ctrlPad/daemon.git`
3. **Branch**: `git checkout -b feature/your-feature`
4. **Commit**: `git commit -m 'feat: add some feature'`
5. **Push**: `git push origin feature/your-feature`
6. **Open** a PR

Please follow the existing code style. Thanks!
