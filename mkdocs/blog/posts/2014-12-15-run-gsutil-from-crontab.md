---
date: 2014-12-15
tags:
- bash
---

# Run gsutil from crontab

Setting up gsutil to sync files over to Google Cloud Storage (on CentOS 6) requires some environment variables to be set, which seems oddly undocumented at the moment. Turning to [Stackoverflow](http://stackoverflow.com/questions/27439326/how-to-properly-run-gsutil-from-crontab/27480249) was the solution, and here’s a summary from that.

<!-- more -->

Also, the `-d` flag is required to truly catch all errors/warnings, but on very large data volumes you may want to omit this flag as it generates quite a large log file. The `> gsutil.log 2>&1` pipe is necessary to pipe both stdout and stderr to file. You can change this to `>> gsutil.log 2>&1` if you wish to append to file rather than overwrite a previous sync log.

In the example below, I’m trying to detect any already running instances of gsutil for this particular folder sync, and if such an instance is detected the gsutil command won’t execute.

Please note, you can find the path to your .boto file by running `gsutil -D ls 2>&1 | grep config_file_list` in a regular shell, outside of crontab. Make sure you run this command as the same user you intend to run gsutil with from within crontab.

Before commencing gsutil, a gcloud components update is being performed to ensure that we are using the latest and greatest version of gsutil.

```
PATH=/sbin:/bin:/usr/sbin:/usr/bin:/home/fredrik/google-cloud-sdk/bin
HOME=/home/fredrik
BOTO_CONFIG="/home/fredrik/.config/gcloud/legacy_credentials/[your-email-address]/.boto"

# Example of job definition:
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
# |  |  |  |  |
# *  *  *  *  * user-name command to be executed

  0  0 */1 *  * fredrik gcloud components update -q
  5  0 */1 *  * fredrik if (ps -ef | grep -v grep | grep "gs://my-bucket/my-folder/"); then echo "Skipping gsutil sync, it is already running."; else gsutil -d -m rsync -r -C /local-folder/ gs://my-bucket/my-folder/ > /gsutil.log 2>&1; fi
```