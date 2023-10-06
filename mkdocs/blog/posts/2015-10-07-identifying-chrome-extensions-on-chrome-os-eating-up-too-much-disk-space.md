---
date: 2015-10-07
authors:
  - fredrikaverpil
comments: true
tags:
- chromeos
---

# Identifying Chrome extensions on Chrome OS eating up too much disk space

I just recently realized I had only 300 MB of free disk space on my Chromebook, when I should have about 8 GB of free disk space. After some investigation, I found that one of my extensions where not deleting its previous versions when it got updated. Here’s how you can detect and fix such an issue.

<!-- more -->

Type <strong>chrome://system</strong> into a Chrome tab, scroll down to “user_files” and click the “expand” button for it. This will generate a (quite long) list of files with file size. Then search for the following in in Chrome using <strong>alt+f</strong>:

    G	/home/

Please note, that’s a [tab character](https://en.wikipedia.org/wiki/Tab_key) between “G” and “/home/” in there. You may want to copy (<strong>ctrl+c</strong>) this from above in order to easily paste it (<strong>ctrl+v</strong>) into the search field. This will find any occurrences of gigabyte-sized folders and may reveal stuff which is taking up a lot of disk space.

If there is a result which looks something like this;

    /home/chronos/user/Extensions/mfaihdlpglflfgpfjcifdjdjcckigekc

...you may do a Google search for that last string (`mfaihdlpglflfgpfjcifdjdjcckigekc`). You’ll probably find out which extension that is. You can also try to construct your own Chrome Web Store link, by placing that long string at the end of `https://chrome.google.com/webstore/detail/` which in this case would result in this link: [https://chrome.google.com/webstore/detail/mfaihdlpglflfgpfjcifdjdjcckigekc](https://chrome.google.com/webstore/detail/mfaihdlpglflfgpfjcifdjdjcckigekc)

So, this reveals we are dealing with a certain extension. Now, to fix this issue, remove that extension from Chrome (<strong>chrome://extensions</strong>) to see those gigabytes get free’d up.

You can re-install the extension, but perhaps you’ll see it keep eating up your disk space over time. It would be a good idea to notify the developers of the extension in question about the problem so that it will get fixed.