# Scaleway Function for Automated Instance Backups

This is a Scaleway Function in Golang that will create a backup for a given Scaleway instance


## How to use it
1. Clone the repository:
2. Compress the content:

`cd automated-instance-backup`

`zip archive.zip vendor main.go go.mod go.sum`

3. Create a Scaleway function with the compressed file using the environment variables: ([see the documentation](https://www.scaleway.com/en/docs/compute/functions/quickstart/) for further help)

`ORGANIZATION_ID: your organization ID`

`ACCESS_KEY: your API access key`

`SECRET_KEY: your API secret key` preferably as a secret

4. Once your function is deployed, you can take your function endpoint and call it by: 

`curl -X POST https://YOUR_FUNCTION_ENDPOINT -d '{"server": "SERVER_UUID",
   "zone": "SERVER_ZONE"}'`

## If you would like to automate the process

Assuming that you deploy your function in Scaleway Serverless, you can activate a CRON trigger (backing up every day at a specific hour for example)