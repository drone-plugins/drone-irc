Use this plugin for sending build status notifications via IRC. You can override
the default configuration with the following parameters:

* `prefix` - Prefix for the sent notifications
* `nick` - Nickname used by the bot
* `channel` - Messages sent are posted here
* `recipient` - Alternatively you can send it to a specific user
* `server` - Connection information for the server
  * `host` - IRC server host to connect to
  * `port` - IRC server port, defaults to 6667
  * `password` - Password for IRC server, optional
  * `tls` - Enable TLS, defaults to false

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  irc:
    prefix: build
    nick: drone
    channel: my-channel
    server:
      host: chat.freenode.net
      port: 6667
```
