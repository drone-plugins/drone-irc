Use this plugin for sending build status notifications via IRC. You can override
the default configuration with the following parameters:

* `prefix` - Prefix for the sent notifications
* `nick` - Nickname used by the bot
* `channel` - Messages sent are posted here
* `recipient` - Alternatively you can send it to a specific user
* `host` - IRC server host to connect to
* `port` - IRC server port, defaults to 6667
* `password` - Password for IRC server, optional
* `enable_tls` - Enable TLS, defaults to false
* `debug` - Enable the ability to debug
* `use_sasl` - When sasl is required by the server
* `sasl_password` - the required sasl password


## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  irc:
    prefix: build
    nick: drone
    channel: my-channel
    host: chat.freenode.net
    port: 6667
```
