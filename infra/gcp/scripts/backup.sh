#!/bin/bash
# Dump database
pg_dump -h ec2-54-155-87-214.eu-west-1.compute.amazonaws.com -p 5432 -U jcatltbukgvvhp d9o2hvngiaghkt > "/srv/backu
ps/list/backup_d9o2hvngiaghkt_$(date +"%Y-%m-%d_%H-%M-%S")"
# Delete older backups
find /srv/backups/list/backup_d9o2hvngiaghkt_* -mtime +3 -exec rm {} \;
# Copy on the GCS bucket
rclone copy /srv/backups/list/ rclone-therentyapp-backups-d9o2hvngiaghkt:therentyapp-backups/d9o2hvngiaghkt
