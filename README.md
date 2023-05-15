# mrpack-install

[![Go Reference](https://pkg.go.dev/badge/github.com/nothub/mrpack-install.svg)](https://pkg.go.dev/github.com/nothub/mrpack-install)

A cli application for installing Minecraft servers and [Modrinth](https://modrinth.com/) [modpacks](https://docs.modrinth.com/docs/modpacks/format_definition/).

---

#### modpack deployment
```
Deploys a Modrinth modpack including Minecraft server.

Usage:
  mrpack-install (<filepath> | <url> | <slug> [<version>] | <id> [<version>]) [flags]
  mrpack-install [command]

Examples:
  mrpack-install https://example.org/data/cool-pack.mrpack
  mrpack-install downloads/cool-pack.mrpack --proxy socks5://127.0.0.1:7890
  mrpack-install hexmc-modpack --server-file server.jar
  mrpack-install yK0ISmKn 1.0.0-1.18 --server-dir mcserver
  mrpack-install communitypack9000 --host api.labrinth.example.org
  mrpack-install --version

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ping        Ping a Labrinth instance
  server      Prepare a plain server environment
  update      Update the deployed modpack
  version     Print version infos

Flags:
      --dl-retries int       Retries when download fails (default 3)
      --dl-threads int       Concurrent download threads (default 8)
  -h, --help                 help for mrpack-install
      --host string          Labrinth host address (default "api.modrinth.com")
      --proxy string         Proxy url for http connections
      --server-dir string    Server directory path (default "mc")
      --server-file string   Server jar file name
  -V, --version              Print version and exit

Use "mrpack-install [command] --help" for more information about a command.
```

---

#### modpack update
```
Update the deployed modpacks files, creating backups if necessary.

Usage:
  mrpack-install update [<version>] [flags]

Flags:
      --backup-dir string   Backup directory path
  -h, --help                help for update

Global Flags:
      --dl-retries int       Retries when download fails (default 3)
      --dl-threads int       Concurrent download threads (default 8)
      --host string          Labrinth host address (default "api.modrinth.com")
      --proxy string         Proxy url for http connections
      --server-dir string    Server directory path (default "mc")
      --server-file string   Server jar file name
```

---

#### plain server deployment
```
Download and configure one of several Minecraft server flavors.

Usage:
  mrpack-install server (vanilla | fabric | quilt | forge | paper) [flags]

Examples:
  mrpack-install server fabric --server-dir fabric-srv
  mrpack-install server paper --minecraft-version 1.18.2 --server-file srv.jar

Flags:
      --flavor-version string      Flavor version (default "latest")
  -h, --help                       help for server
      --minecraft-version string   Minecraft version (default "latest")

Global Flags:
      --dl-retries int       Retries when download fails (default 3)
      --dl-threads int       Concurrent download threads (default 8)
      --host string          Labrinth host address (default "api.modrinth.com")
      --proxy string         Proxy url for http connections
      --server-dir string    Server directory path (default "mc")
      --server-file string   Server jar file name
```
