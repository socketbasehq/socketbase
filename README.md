# SocketBase

SocketBase is an open-source, self-hosted alternative to Pusher Channels that enables real-time bidirectional communication in your applications. Built with Go, it provides a reliable and scalable solution for developers who want to maintain control over their real-time infrastructure.

## Features

- 🚀 **Self-Hosted**: Full control over your WebSocket infrastructure
- 🔌 **Pusher Compatible**: Drop-in replacement for Pusher Channels
- 📱 **Multiple Apps**: Create and manage multiple applications under one instance
- 🔒 **Secure**: Built-in authentication and private channels support
- 🐳 **Easy Deployment**: Deploy via Docker or build from source
- ⚡ **High Performance**: Built with Go for optimal performance
- 🛠 **Developer Friendly**: Simple setup process with minimal configuration

## Quick Start

### Using Docker

```bash
docker run -p 8000:8000 socketbase/socketbase
```

### From Source

```bash
# Clone the repository
git clone https://github.com/socketbasehq/socketbase.git

# Navigate to the project directory
cd socketbase

# Install dependencies
go mod download

# Start the server
go run cmd/main.go
```

## Configuration

Create a `.env` file in your project root:

```env
PORT=8000
HOST=localhost
DATABASE_URL=postgresql://localhost:5432/socketbase
```

## Client Usage

SocketBase is compatible with Pusher client libraries. Here's how to connect using the Pusher JavaScript client:

```javascript
const pusher = new Pusher('your-app-key', {
  wsHost: '127.0.0.1',
  wsPort: 3000,
  forceTLS: false,
});
```

### Subscribe to a Channel

```javascript
const channel = pusher.subscribe('my-channel');
channel.bind('my-event', data => {
  console.log('Received event:', data);
});
```

## Development

```bash
# Run tests
go test ./...

# Build the project
go build cmd/main.go
```

## Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- Create an issue in the [Issue Tracker](https://github.com/socketbasehq/socketbase/pkg/issues)
- Join our [Discord community](https://discord.gg/zSQyK6nmM9) for discussions
- Star the repository to show your support

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/mxvsh">mxvsh</a>
</p>
