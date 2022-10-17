#!/bin/sh
# This script is used to run the ETL process

cd server
echo "Downloading the data from the source..."
out=$(curl -O http://download.srv.cs.cmu.edu/\~enron/enron_mail_20110402.tgz 2>&1)

if [[ $? -eq 0 && $out ]]; then
    echo "Downloaded successfully!"
    echo "Extracting the data..."
    tar -xzf enron_mail_20110402.tgz
    echo "Extracted successfully!"
else
    echo "Download failed"
fi