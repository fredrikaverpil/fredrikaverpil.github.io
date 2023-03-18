---
ShowToc: false
TocOpen: false
date: 2019-05-27 02:00:12+02:00
draft: false
summary: Quick notes on getting set up with remote development in Visual Studio code
  on Windows 10.
tags:
- windows
title: Visual Studio Code remote development with Windows 10
---

## Remote development over SSH

First make sure the following commands are executed from `C:\System32\OpenSSH`: `ssh`, `ssh-keygen`, `scp`. If they are located in e.g. Chocolatey's bin folder, keys will be searched for in weird places. This can be verifed by installing `which`:

```bash
choco install which
which ssh
```

### Create key pair in Windows 10 client

```bash
cd ~
mkdir .ssh
cd ssh
ssh-keygen -t rsa -b 4096
```

- RSA is the [recommended](https://security.stackexchange.com/questions/5096/rsa-vs-dsa-for-ssh-authentication-keys) authentication security.
- The -b flag instructs `ssh-keygen` to increase the number of bits used to generate the key pair, and is suggested for additional security.

### Prepare server

I am using a CentOS 7 server.

```bash
ssh user@server-hostname
mkdir -p ~/.ssh
touch ~/.ssh/authorized_keys
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
exit
```

### Transfer public key from Windows 10 onto server

```bash
scp ~/.ssh/id_rsa.pub user@server-hostname:~/.ssh/authorized_keys
```

**Note**: this will just _copy_ the file onto the server, overwriting an already existing `authorized_keys` file if it exists.

Replace `id_rsa.pub` with your .pub file name, if you named it explicitly.

### Test SSH keys

```bash
ssh user@server-hostname  # should not prompt for password
```

### SSH config in Windows 10

If you haven't already, install the "Remote development" extension in vscode.

Click the lower left `><` icon in vscode and choose "Remote-SSH: Open configuration file...". Then enter something like:

```
# Read more about SSH config files: https://linux.die.net/man/5/ssh_config
Host PROFILE
    HostName HOSTNAME
    User USERNAME
```

- Replace `PROFILE` with a name you will use from a dropdown menu to choose the server later on.
- Replace `HOSTNAME` with the server's hostname or IP address.
- Replace `USERNAME` with the username used to log in over SSH.

More options are outlined in the [docs](https://code.visualstudio.com/docs/remote/ssh).

### Connect to server over SSH and code away

Click the lower left `><` icon in vscode and choose "Remote-SSH: Connect to Host...". Choose your newly created profile.

A new vscode session will launch, and from the "Open folder" dialog, choose your project's folder. Voila!