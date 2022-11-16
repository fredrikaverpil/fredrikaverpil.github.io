---
title: Async and await with subprocesses
draft: true
tags: [python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2017-06-20
updated: 2022-11-16T23:22:33+01:00
---

A boilerplate which can be used on Windows and Linux/macOS in order to asynchronously run subprocesses. This requres Python 3.6.

**Update 2019-06-28:** Fixed a problem where the loop got closed prematurely, added better progress messages, tested on Python 3.7.3.

```python
"""Async and await example using subprocesses

Note:
    Requires Python 3.6.
"""

import asyncio
import platform
import sys
import time
from pprint import pprint


async def run_command(*args):
    """Run command in subprocess.

    Example from:
        http://asyncio.readthedocs.io/en/latest/subprocess.html
    """
    # Create subprocess
    process = await asyncio.create_subprocess_exec(
        *args, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE
    )

    # Status
    print("Started: %s, pid=%s" % (args, process.pid), flush=True)

    # Wait for the subprocess to finish
    stdout, stderr = await process.communicate()

    # Progress
    if process.returncode == 0:
        print(
            "Done: %s, pid=%s, result: %s"
            % (args, process.pid, stdout.decode().strip()),
            flush=True,
        )
    else:
        print(
            "Failed: %s, pid=%s, result: %s"
            % (args, process.pid, stderr.decode().strip()),
            flush=True,
        )

    # Result
    result = stdout.decode().strip()

    # Return stdout
    return result


async def run_command_shell(command):
    """Run command in subprocess (shell).

    Note:
        This can be used if you wish to execute e.g. "copy"
        on Windows, which can only be executed in the shell.
    """
    # Create subprocess
    process = await asyncio.create_subprocess_shell(
        command, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE
    )

    # Status
    print("Started:", command, "(pid = " + str(process.pid) + ")", flush=True)

    # Wait for the subprocess to finish
    stdout, stderr = await process.communicate()

    # Progress
    if process.returncode == 0:
        print("Done:", command, "(pid = " + str(process.pid) + ")", flush=True)
    else:
        print("Failed:", command, "(pid = " + str(process.pid) + ")", flush=True)

    # Result
    result = stdout.decode().strip()

    # Return stdout
    return result


def make_chunks(l, n):
    """Yield successive n-sized chunks from l.

    Note:
        Taken from https://stackoverflow.com/a/312464
    """
    if sys.version_info.major == 2:
        for i in xrange(0, len(l), n):
            yield l[i : i + n]
    else:
        # Assume Python 3
        for i in range(0, len(l), n):
            yield l[i : i + n]


def run_asyncio_commands(tasks, max_concurrent_tasks=0):
    """Run tasks asynchronously using asyncio and return results.

    If max_concurrent_tasks are set to 0, no limit is applied.

    Note:
        By default, Windows uses SelectorEventLoop, which does not support
        subprocesses. Therefore ProactorEventLoop is used on Windows.
        https://docs.python.org/3/library/asyncio-eventloops.html#windows
    """
    all_results = []

    if max_concurrent_tasks == 0:
        chunks = [tasks]
        num_chunks = len(chunks)
    else:
        chunks = make_chunks(l=tasks, n=max_concurrent_tasks)
        num_chunks = len(list(make_chunks(l=tasks, n=max_concurrent_tasks)))

    if asyncio.get_event_loop().is_closed():
        asyncio.set_event_loop(asyncio.new_event_loop())
    if platform.system() == "Windows":
        asyncio.set_event_loop(asyncio.ProactorEventLoop())
    loop = asyncio.get_event_loop()

    chunk = 1
    for tasks_in_chunk in chunks:
        print("Beginning work on chunk %s/%s" % (chunk, num_chunks), flush=True)
        commands = asyncio.gather(*tasks_in_chunk)  # Unpack list using *
        results = loop.run_until_complete(commands)
        all_results += results
        print("Completed work on chunk %s/%s" % (chunk, num_chunks), flush=True)
        chunk += 1

    loop.close()
    return all_results


def main():
    """Main program."""
    start = time.time()

    if platform.system() == "Windows":
        # Commands to be executed on Windows
        commands = [
            ["hostname"],
        ]
    else:
        # Commands to be executed on Unix
        commands = [
            ["du", "-sh", "/var/tmp"],
            ["hostname"],
        ]

    tasks = []
    for command in commands:
        tasks.append(run_command(*command))

    # # Shell execution example
    # tasks = [run_command_shell('copy c:/somefile d:/new_file')]

    # # List comprehension example
    # tasks = [
    #     run_command(*command, get_project_path(project))
    #     for project in accessible_projects(all_projects)
    # ]

    results = run_asyncio_commands(
        tasks, max_concurrent_tasks=20
    )  # At most 20 parallel tasks
    print("Results:")
    pprint(results)

    end = time.time()
    rounded_end = "{0:.4f}".format(round(end - start, 4))
    print("Script ran in about %s seconds" % (rounded_end), flush=True)


if __name__ == "__main__":
    main()
```