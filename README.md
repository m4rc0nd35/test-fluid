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
- Only processing of queue
```
Get all lead data GET [Link](http://localhost:8081/lead/all)
Get data lead by uuid GET [Link](http://localhost:8081/lead/fff693bc-f5e9-48de-ab2b-9e83d39fa2c3)
### Microservice C

```json
{
        "gender": "male",
        "name": {
            "title": "Mr",
            "first": "سام",
            "last": "کریمی"
        },
        "statusFlow": "processed",
        "location": {
            "street": {
                "number": 7983,
                "name": "فاطمی"
            },
            "city": "همدان",
            "state": "خوزستان",
            "country": "Iran",
            "postcode": 84532,
            "coordinates": {
                "latitude": "17.3658",
                "longitude": "-166.6431"
            },
            "timezone": {
                "offset": "-8:00",
                "description": "Pacific Time (US & Canada)"
            }
        },
        "email": "sm.khrymy@example.com",
        "login": {
            "uuid": "6d2c8c5d-ad34-4a69-8c18-9b282abf96db",
            "username": "beautifulleopard657",
            "password": "sanchez",
            "salt": "OZorQALc",
            "md5": "2827781ac3ac3b0331a21ebdded2c4ea",
            "sha1": "312c89d83559d0229ee183622c27d599ca265aa7",
            "sha256": "67006e948879b7152d4350726fecdd4247555a7367e6e72fe5ff55c492844f95"
        },
        "dob": {
            "date": "1986-10-30T13:34:47.963Z",
            "age": 36
        },
        "registered": {
            "date": "2005-03-29T19:26:32.421Z",
            "age": 17
        },
        "phone": "016-79548777",
        "cell": "0909-087-8563",
        "id": {
            "name": "",
            "value": "086028195"
        },
        "picture": {
            "large": "https://randomuser.me/api/portraits/men/55.jpg",
            "medium": "https://randomuser.me/api/portraits/med/men/55.jpg",
            "thumbnail": "https://randomuser.me/api/portraits/thumb/men/55.jpg"
        },
        "nat": "IR",
        "createdAt": "2022-12-03T02:44:26Z"
    }
```