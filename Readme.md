# go-home-or-away

`go-home-or-away` is a conditional proxy for SSH, written in Go. It's designed to be used with the `ProxyCommand` option in your SSH client configuration.

## What it does

This tool helps you manage SSH connections in environments with dynamic network access.

-   **When you're "home"**: If you can reach a target host and port directly over TCP, it will connect you directly.
-   **When you're "away"**: If the direct connection fails, it will automatically proxy your connection through a specified SSH server, using `ssh -W "$host:$port" $proxyhost`.

This is useful when you move between a network that has direct access to your servers (e.g., an office network or VPN) and one that doesn't (e.g., a public Wi-Fi network).

## Building

The recommended way to build this application is with Docker to ensure a consistent build environment.

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/go-home-or-away.git
    cd go-home-or-away
    ```

2.  **Build the executable**:
    The following command will build the Windows executable (`go-home-or-away.exe`) in the current directory.
    ```bash
    docker run --rm -v $(pwd):/app -e GOOS=windows -e GOARCH=amd64 -w /app golang go build -o go-home-or-away.exe
    ```
    To build for other operating systems, change the `GOOS` and `GOARCH` environment variables. For example, for Linux, you would use `GOOS=linux` and `GOARCH=amd64`.

## Usage

To use `go-home-or-away`, add it as a `ProxyCommand` in your `~/.ssh/config` file.

1.  Place the `go-home-or-away.exe` executable in a known location on your machine (e.g., `C:\Users\your-user\bin`).

2.  Edit your `~/.ssh/config` file and add an entry for your host:

    ```
    Host my-remote-machine
      HostName internal.server.com
      User your-user
      # Use the Go app as the proxy command
      # The %h and %p are automatically replaced by SSH with the HostName and port
      ProxyCommand C:\path\to\your\app\go-home-or-away.exe %h %p your-ssh-proxy.example.com
    ```

    Replace `C:\path\to\your\app` with the actual path to the directory containing the executable, and `your-ssh-proxy.example.com` with your actual proxy server.

Now, when you run `ssh my-remote-machine`, `go-home-or-away` will automatically handle the connection for you.

## License

This project is licensed under the MIT License.
