# go-home-or-away

`go-home-or-away` is a conditional proxy for SSH, written in Go. It's designed to be used with the `ProxyCommand` option in your SSH client configuration.

## What it does

This tool helps you manage SSH connections in environments with dynamic network access.

-   **When you're "home"**: If you can reach a target host and port directly over TCP, it will connect you directly.
-   **When you're "away"**: If the direct connection fails, it will automatically proxy your connection through a specified SSH server, using `ssh -W "$host:$port" $proxyhost`.

This is useful when you move between a network that has direct access to your servers (e.g., an office network or VPN) and one that doesn't (e.g., a public Wi-Fi network).

## Building

You can build this application from source if you have a Go environment set up.

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/cfullelove/go-home-or-away.git
    cd go-home-or-away
    ```

2.  **Build the executable**:
    To build the executable for your current operating system, run:
    ```bash
    go build -o go-home-or-away
    ```
    To build for a specific operating system, use the `GOOS` and `GOARCH` environment variables. For example, to build for Windows:
    ```bash
    GOOS=windows GOARCH=amd64 go build -o go-home-or-away.exe
    ```

Alternatively, you can download pre-compiled binaries for Linux and Windows from the [Releases](https://github.com/cfullelove/go-home-or-away/releases) page.

## Usage

To use `go-home-or-away`, add it as a `ProxyCommand` in your SSH client configuration file.

### Windows

1.  Place the `go-home-or-away.exe` executable in a known location on your machine (e.g., `C:\Users\your-user\bin`).

2.  Edit your `%USERPROFILE%\.ssh\config` file and add an entry for your host:

    ```
    Host my-remote-machine
      HostName internal.server.com
      User your-user
      # Use the Go app as the proxy command
      # The %h and %p are automatically replaced by SSH with the HostName and port
      ProxyCommand C:\path\to\your\app\go-home-or-away.exe %h %p your-ssh-proxy.example.com
    ```

    Replace `C:\path\to\your\app` with the actual path to the directory containing the executable, and `your-ssh-proxy.example.com` with your actual proxy server.

### Linux

1.  Place the `go-home-or-away` executable in a known location on your machine (e.g., `/usr/local/bin`).

2.  Make sure the executable has execute permissions:
    ```bash
    chmod +x /usr/local/bin/go-home-or-away
    ```

3.  Edit your `~/.ssh/config` file and add an entry for your host:

    ```
    Host my-remote-machine
      HostName internal.server.com
      User your-user
      # Use the Go app as the proxy command
      # The %h and %p are automatically replaced by SSH with the HostName and port
      ProxyCommand /usr/local/bin/go-home-or-away %h %p your-ssh-proxy.example.com
    ```

    Replace `your-ssh-proxy.example.com` with your actual proxy server.

Now, when you run `ssh my-remote-machine`, `go-home-or-away` will automatically handle the connection for you.

## License

This project is licensed under the MIT License.
