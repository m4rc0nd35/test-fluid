# test-fluid

- [x] 1. Get lead by repository X
- [x] 2. Add endpoint setting params
- [x] 3. Define time by cron
- [x] 4. Lead received queue H new status send to H
- [x] 5. Received H, send to I, status processing
- [x] 6. Waiting T, send to K with status processed
- [x] 7. Case receive in J, is recused status
- [x] 8. Logger all events
- [x] 9. Add endpoint find lead by uuid
- [x] 10. Add endpoint find all lead
- [x] 11. Add endpoint find history status changed by uuid
- [x] 12. Add endpoint find stats of status
- [x] 13. Unique instance
- [x] 14. A/B accept multiple instance


### Microservice A
Link set for params cron, pause and lead limit by job [Settings](http://localhost:8080/setting)
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

### Microservice B
```

```