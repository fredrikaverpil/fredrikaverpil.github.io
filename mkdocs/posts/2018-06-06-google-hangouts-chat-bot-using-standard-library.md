---
date: 2018-06-06
tags:
- python
---

# Google Hangouts Chat incoming webhook using Python standard library only

The [official docs on setting up an incoming webhook](https://developers.google.com/hangouts/chat/quickstart/incoming-bot-python) uses the third-party [httplib2](https://github.com/httplib2/httplib2), which is not part of the Python 3.6 standard library. Here's a quick snippet using only the standard library instead:

<!-- more -->

```python
import json
import urllib.parse
import urllib.request


def main():
    # python 3.6

    url = '<INCOMING-WEBHOOK-URL>'
    bot_message = {'text': 'Hello from Python script!'}
    message_headers = {'Content-Type': 'application/json; charset=UTF-8'}

    byte_encoded = json.dumps(bot_message).encode('utf-8')
    req = urllib.request.Request(url=url, data=byte_encoded, headers=message_headers)
    response = urllib.request.urlopen(req)

    print(response.read())


if __name__ == '__main__':
    main()
```