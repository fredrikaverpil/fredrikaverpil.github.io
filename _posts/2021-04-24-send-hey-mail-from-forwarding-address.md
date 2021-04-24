---
layout: post
title: 'Send HEY mail with forwarding address using Office365 SMTP'
tags: [workflow]
---

[HEY mail](https://hey.com/) announced a while back you can [send emails from external email addresses using SMTP](https://hey.com/features/send-as/). Just the other day, they now finally also announced that [their take on custom domains](https://hey.com/domains/).

I've been waiting to hook up my existing HEY mail address with my custom domain so that I don't have to read emails in HEY but compose/reply in GMail (where my custom domain works fine).

In my case, I was hoping for something more similar to what GMail's offering is like, where you can choose to send email from other email addresses which you can receive email from (e.g. forwarding addresses).

This guide outlines a solution I've found, which works great for me, where I can now finally send email from a custom domain forward address, using the [Office365 SMTP](https://docs.microsoft.com/en-us/exchange/mail-flow-best-practices/how-to-set-up-a-multifunction-device-or-application-to-send-email-using-microsoft-365-or-office-365).

<!--more-->

## What you'll need

* A domain registered with a registrar where you can set up email addresses with forwarding.
* An Office365 account

## Pros over HEY mail's official custom domains feature

* Keep your `<username>@hey.com` address and all mail you already have in HEY
* Choose whether to send email from `<username>@hey.com` or your `<username>@yourdomain.com`

## Setup

1. In your domain registrar's website, setup your desired email address `<username>@yourdomain.com` and forward it to your `<username>@hey.com`
1. In your Microsoft account, set up `<username>@yourdomain.com` as ["Microsoft Account Alias"](https://account.live.com/AddAssocId)
1. Once this has been done, verify that you can send email using `<username>@yourdomain.com`, from the web UI at https://outlook.live.com/
1. In your Microsoft account, set up an ["App password"](https://support.microsoft.com/en-us/account-billing/using-app-passwords-with-apps-that-don-t-support-two-step-verification-5896ed9b-4263-e681-128a-a6f2979a7944).
1. In HEY mail's account settings, go to "Forwarding & SMTP setup" and add `<username>@yourdomain.com`. Also add SMTP details, and use the ones from Office365 (you can find these details under the "Sync email" mail setting):
    * Outgoing (SMTP) server: `smtp.office365.com`
    * Port: `587`
    * Username: `<Microsoft username>`
    * Password: `<app password>`

Now, when composing emails in HEY mail, you can choose whether to send the email as `<username>@hey.com` or `<username>@yourdomain.com`!

## Closing comments

All your emails are now stored in HEY, but you will also have all sent emails stored in Office365, where you have 5 GB of free space.
