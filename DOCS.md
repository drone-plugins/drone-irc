Use the IRC plugin to send a message to an IRC channel when a build completes.
The following parameters are used to configuration the notification:

* **host** - connects to this host
* **port** - connects to this port
* **password** - authenicates using this password
* **channel** - messages sent to the above server are posted here
* **nick** - choose the nickname this plugin will post as
* **prefix** - choose the prefix for the sent notifications

The following is a sample IRC configuration in your .drone.yml file:

```yaml
notify:
  irc:
    prefix: drone
    nick: drone
    channel: #dev
    server:
      host: chat.freenode.net
      port: 6697
```
