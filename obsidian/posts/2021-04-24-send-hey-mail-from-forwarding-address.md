---
title: 'Send HEY mail with forwarding address using Office365 SMTP'
tags: [workflow]
draft: true

# PaperMod
ShowToc: true
TocOpen: true

date: 2021-04-24T02:00:00+02:00
---

[HEY mail](https://hey.com/) announced a while back you can [send emails from external email addresses using SMTP](https://hey.com/features/send-as/). Just the other day, they now finally also announced that [their take on custom domains](https://hey.com/domains/).

I've been waiting to hook up my existing HEY mail address with my custom domain's forwarding addresses, so that I don't have to set up MX records for email on my custom domain's DNS.

In my case, I was hoping for something more similar to what GMail's offering is like, where you can choose to send email from any other email addresses which you can receive email from (e.g. forwarding addresses, registered directly in your domain registrar's DNS settings).

This guide outlines a solution I've found, which works great for me, where I can now finally send HEY email from a custom domain forward address, using the [Office365 SMTP](https://docs.microsoft.com/en-us/exchange/mail-flow-best-practices/how-to-set-up-a-multifunction-device-or-application-to-send-email-using-microsoft-365-or-office-365).



## Update 2021-09-25

I've realized this is not a great workaround for using Hey's custom domains. The reason is that when emails are sent using the Microsoft 365 SMTP server, the sender address does not match that of the SMTP server.

This results in the "via outlook.com" notice in e.g. Gmail:

![](fredrikaverpil.github.io/obsidian/static/hey/outlook_smtp.png)

It was recently brought to my attention that my emails often end up in the spam folder for recipients who does not already have me in their address book. And I think this may be because of the sender address not matching my domain name.

Therefore I am moving away from using Hey with Microsoft's SMTP server and the setup outlined in this blog post.

With this warning in mind, you are welcome to continue reading this post although I no longer recommend this setup.

## What you'll need

* A domain registered with a registrar where you can set up email addresses with forwarding.
* An [Microsoft/Office365/Outlook](http://outlook.live.com/) account

## Pros over setting up MX records on your DNS

* You can set up forwarding addresses in the DNS settings at your domain registrar (MAILFW rules).

## Pros over HEY mail's official custom domains feature

* Keep your `<username>@hey.com` address and all mail you already have in HEY
* Choose whether to send email from `<username>@hey.com` or your `<username>@yourdomain.com`

## Setup

1. In your domain registrar's website, setup your desired email address `<username>@yourdomain.com` and forward it to your `<username>@hey.com`
1. In your Microsoft account, set up `<username>@yourdomain.com` as ["Microsoft Account Alias"](https://account.live.com/AddAssocId)
1. Once this has been done, verify that you can send email using `<username>@yourdomain.com`, from the web UI at [outlook.live.com](https://outlook.live.com/)
1. In your Microsoft account, set up an ["App password"](https://support.microsoft.com/en-us/account-billing/using-app-passwords-with-apps-that-don-t-support-two-step-verification-5896ed9b-4263-e681-128a-a6f2979a7944).
1. In HEY mail's account settings, go to "Forwarding & SMTP setup" and add `<username>@yourdomain.com`. Also add SMTP details, and use the ones from Office365 (you can find these details under the "Sync email" mail setting):
    * Outgoing (SMTP) server: `smtp.office365.com`
    * Port: `587`
    * Username: `<Microsoft username>`
    * Password: `<app password>`

Now, when composing emails in HEY mail, you can choose whether to send the email as `<username>@hey.com` or `<username>@yourdomain.com`!

## Closing comments

All your emails are now stored in HEY, but you will also have all sent emails stored in Office365, where you have 5 GB of free space.
