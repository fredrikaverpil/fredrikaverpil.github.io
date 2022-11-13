---
layout: page
title: Archive
---

{% include filter_by_tag.html %}

{% for post in site.posts %}<span class="archivemono">{{ post.date | date_to_string }}</span> &raquo; [ {{ post.title }} ]({{ post.url }})  
{% endfor %}
