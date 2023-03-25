---
date: 2017-06-28
tags:
- linux
---

# Find missing libs with repoquery

Not sure how I didn't learn about this until today. Anyways, if you end up with missing libs on CentOS/RedHat, use [repoquery](http://yum.baseurl.org/wiki/RepoQuery) to find missing libs.

<!-- more -->

For example, I had an issue today where docker required `libltdl.so.7` but this wasn't installed.


```bash
repoquery -q -f */libltdl.so.*
```

The above is quering (`-q`) for files (`-f`) matching the pattern (`*/libltdl.so*`). The first star is important since the query is looking for matches against the full paths of the files within the RPMs stored on the various YUM repos your system is aware of.

...and the response was:

```
libtool-ltdl-0:2.4.2-21.el7_2.i686
libtool-ltdl-0:2.4.2-22.el7_3.i686
libtool-ltdl-0:2.4.2-21.el7_2.x86_64
libtool-ltdl-0:2.4.2-22.el7_3.x86_64
```

The solution in this case was to `yum install libtool-ltdl`. Thanks `repoquery`!