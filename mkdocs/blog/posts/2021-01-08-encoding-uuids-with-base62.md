---
date: 2021-01-08
authors:
  - fredrikaverpil
comments: true
tags:
- python
---

# Encoding UUIDs with base62

This was just a fun experiment.

<!-- more -->

## Base62 vs base64

In the very common [base64](https://en.wikipedia.org/wiki/Base64) encoding scheme, 64 characters are used for binary-to-text encoding:

```text
0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/
```

Sometimes, the `+` and `/` characters can be undesired in the encoded string. An example could be a base64-encoded [universally unique id](https://en.wikipedia.org/wiki/Universally_unique_identifier) to be used in a URL.

So, if we remove those two characters, we end up with 62 characters, which indeed makes out the characters used in the [base62](https://en.wikipedia.org/wiki/Base62) encoding scheme:

```text
0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
```

## Simple implementation

Here's `base62.py`:

```python
"""Base62 encoding/decoding.

Example usage:

    >>> import uuid
    >>> import base62  # this module
    >>> i = uuid.uuid4()
    >>> i
    UUID('a40362fe-9aac-4587-90cb-fb6cf306da75')
    >>> i.int
    218010976047265199223445579130814782069
    >>> b = base62.encode(i.int)
    >>> b
    '4ZuhWgpKFPPpOD21d5urcx'
    >>> e = base62.decode('4ZuhWgpKFPPpOD21d5urcx')
    >>> e
    218010976047265199223445579130814782069
    >>> i2 = uuid.UUID(int=218010976047265199223445579130814782069)
    >>> i2
    UUID('a40362fe-9aac-4587-90cb-fb6cf306da75')

Test:

    >>> base62.verify(i, e)
    True
"""

import uuid

BASE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


def encode(num, alphabet=BASE62):
    """Encode a positive number in Base X.

    Arguments:
        `num` (int): The number to encode
        `alphabet` (str): The alphabet to use for encoding
    """
    if num == 0:
        return alphabet[0]
    arr = []
    base = len(alphabet)
    while num:
        num, rem = divmod(num, base)
        arr.append(alphabet[rem])
    arr.reverse()
    return ''.join(arr)


def decode(string, alphabet=BASE62):
    """Decode a Base X encoded string into the number.

    Arguments:
        `string` (str): The encoded string
        `alphabet` (str): The alphabet to use for encoding
    """
    base = len(alphabet)
    strlen = len(string)
    num = 0

    idx = 0
    for char in string:
        power = (strlen - (idx + 1))
        num += alphabet.index(char) * (base**power)
        idx += 1

    return num


def verify(original_uuid, base62_str):
    """Verify that the base62-encoded string can be decoded back."""
    reconstructed_uuid = uuid.UUID(int=decode(base62_str))

    assert str(original_uuid) == str(reconstructed_uuid)

    return True
```