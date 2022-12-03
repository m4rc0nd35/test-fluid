# test-fluid



### Microservice A
Link set for params cron, pause and limit lead by job [Settings](localhost:8080/setting)
```
{
    "pause": false,
    "getLimit": 500,
    "cron": "43 * * * *"
}
// pause true | false
// getLimit int (limit conquest lead by job)
// Cron:
// * * * * *
// | | | | |
// | | | | |
// | | | | +---- Day of week  (interval: 1-7)
// | | | +------ Month        (interval: 1-12)
// | | +-------- Day          (interval: 1-31)
// | +---------- Hour         (interval: 0-23)
// +------------ Minute       (interval: 0-59)
```