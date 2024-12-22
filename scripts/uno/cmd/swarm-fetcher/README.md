# Swarm History Downloader

This is a very simple tool that allows you to download your [Swarm](https://www.swarmapp.com/)
history as a JSON file with some information. To use it, you need to:

1. Create an application on https://foursquare.com/developers/home and set the
   Redirect URI to `http://localhost:8050/callback`. Then, store the client ID
   and the client secret.
2. Copy `config.example.yaml` into `config.yaml` and fill it according to
   instructions.
3. Run `go run *.go login` to login and follow command.
4. Run `go run *.go fetch` to fetch the history.

## License

MIT © Henrique Dias
