# dutil - DevOps CLI Tool

**dutil** is a CLI tool written in Go for DevOps tasks, providing functionality to manage insecure registries. It allows you to add and remove insecure registries from your Docker daemon configuration.

## Usage

**dutil** provides functionality to add and remove insecure registries from your Docker daemon configuration.

### Adding an Insecure Registry

You can use the `addins` command to add an insecure registry to your Docker daemon configuration.

- `<registry_hostname>`: The hostname of the insecure registry to add.
- `--daemon-file <path_to_daemon.json>`: (Optional) A custom path to the daemon.json file. Default: `/etc/docker/daemon.json`.
- `--port <port_number>`: (Optional) A custom port to access the registry.

example: 
```sh
dutil addins example.com --daemon-file /path/to/custom/daemon.json --port 8080
```

### Removing an Insecure Registry

You can use the `rmins` command to add an insecure registry to your Docker daemon configuration.

- `<registry_hostname>`: The hostname of the insecure registry to remove.
- `--daemon-file <path_to_daemon.json>`: (Optional) A custom path to the daemon.json file. Default: `/etc/docker/daemon.json`.

example:
```sh
dutil rmins example.com --daemon-file /path/to/custom/daemon.json
```
## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
