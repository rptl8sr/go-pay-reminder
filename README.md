# Reminder for payments

Go script to check table in Google Sheet and send messages as Telegram bot. Run with Yandex Cloud Functions

### Environments
Use '.env' file for local development
```shell
TG_API_TOKEN=
CHAT_ID=
SHEET_ID=
LOG_LEVEL=
```
Add vars to Yandex Cloud Functions environments to production use.


### Google Authorization
1. Create project on Google Cloud https://console.cloud.google.com/projectcreate
2. Enable Google Sheets API https://console.cloud.google.com/apis/library/sheets.googleapis.com
3. Create service account https://console.cloud.google.com/apis/credentials
4. Get service account key https://console.cloud.google.com/iam-admin/serviceaccounts/details/
5. Save it as serviceCreds.json in project root
6. Create Google Sheet and get Sheet ID (like https://docs.google.com/spreadsheets/d/THERE_IS_A_SHEET_ID/)
7. Add service account email (like 'service-account-name@service-name.iam.gserviceaccount.com') to the Google Sheet Viewers/Editor 


### Google Sheet table
Headers of table
```shell
name, url, ip, llc, login, commen, pay day, amount
```

### Yandex Cloud Function
Zip project
```shell
make yzip
```
Upload to Functions

