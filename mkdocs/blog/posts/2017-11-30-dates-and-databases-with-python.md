---
date: 2017-11-30
authors:
  - fredrikaverpil
comments: true
tags:
- python
---

# Dates and databases with Python

It's a bit tricky to deal with dates, timezones and daylight savings when you need to store dates in e.g. a database for later reading.

<!-- more -->

To me, it's a bit perplexing that all tools required to deal with this _doesn't_ come with the Python standard library (meaning; batteries are not included). Instead we need to use three different modules: [datetime](https://docs.python.org/3/library/datetime.html), [pytz](http://pytz.sourceforge.net) and [tzlocal](https://github.com/regebro/tzlocal) where the two latter ones are not part of the standard library and must be installed separately via e.g. `pip`.

Here follows some personal notes on how to store and read back dates with reliability and control of the timezones and daylight savings.

When reading the below code it helps to think that the `tz` abbreviation means "timezone".


### Getting the current local time

You can use `now` to get the local time right now:

```python
>>> import datetime
>>> datetime.datetime.now()
>>> print(now)
2017-11-30 09:01:19.676817
```

However, this does not tell us whether daylight savings was in effect or not. And we don't know from which timezone this date was taken if we were just given these numbers. This becomes a problem when storing the date in e.g. a database and later want to show this date in a user's local timezone.


### Storing dates as UTC-aware

When storing dates in e.g. a database, store them aware of [UTC](https://en.wikipedia.org/wiki/Coordinated_Universal_Time) (or "UTC-aware"):

```python
>>> import datetime
>>> import pytz
>>> now = datetime.datetime.now(pytz.utc)  # store this in db
>>> print(now)
2017-11-30 08:01:19.676817+00:00
```

Please note, the actual local time for me (who is UTC+1) is `09:01:19` and not `08:01:19`. But instead of storing my local UTC+1 datetime, we store a datetime which is just aware of UTC, thanks to `pytz.utc`.


### Reading UTC-aware dates back and showing them accurately

Later, when reading the dates back from e.g. a database, apply the user's local timezone and any daylight savings (in my case UTC+1 right now since I'm in Sweden).

```python
>>> import datetime
>>> import pytz
>>> timezone = 'Europe/Stockholm'
>>> local_tz = pytz.timezone(timezone)
>>> local_dt = now.replace(tzinfo=pytz.utc).astimezone(local_tz)
>>> now_local = local_tz.normalize(local_dt)  # show this to the user
>>> print(now_local)
2017-11-30 09:01:19.676817+01:00
```

Now I get the time I was expecting, my local time `09:01:19`, again thanks to `pytz` and the [tz database](https://en.wikipedia.org/wiki/Tz_database) which it is using.

For a list of all timezones supported by `pytz`, see [here](https://github.com/newvem/pytz/blob/f137ff00112a9682bc4e4945067b3b88f158d010/pytz/zone.tab).


### Avoid hardcoding the local timezone

In the previous code block, I'm hardcoding the `timezone` variable. You may want to read the client system's timezone and just use that. This can be achieved using the `tzlocal` module.

```python
>>> import datetime
>>> import pytz
>>> import tzlocal
>>> local_tz = tzlocal.get_localzone()
>>> local_dt = now.replace(tzinfo=pytz.utc).astimezone(local_tz)
>>> now_local = local_tz.normalize(local_dt)  # show this to the user
>>> print(now_local)
2017-11-30 09:01:19.676817+01:00
```

### Warning: datetime's timedelta does not consider DST

This was added on 2022-11-06.

When using `datetime.timedelta` to add or subtract time, it does not consider daylight savings. I created a gist [here](https://gist.github.com/fredrikaverpil/0cde09c624824ebafe0cb94a6cca9e1e) to illustrate the problem.

### Closing comments

Having all this code finally assembled and condensed in a blog post like this is nice and neat, but why is this so hard to do currently, and why does this require three separate modules of which two are not included in the standard library?

I guess the answer is I should shut up because I'm already spoiled with Python? ;)

I would like to include a link to [this fantastic article](https://zachholman.com/talk/utc-is-enough-for-everyone-right), written by Zach Holman. This touches on the history and complexity of managing timezones, which is not only educational but also quite funny reading. For those interested, here's also the [Hacker News thread](https://news.ycombinator.com/item?id=17181046) which followed after this article was published.