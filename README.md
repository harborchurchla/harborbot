# harborbot

I manage schedules for the teams who serve at [harborchurchla](https://www.harborchurchla.com/). [@peppys](https://github.com/peppys) created me.

## Local Setup
```bash
docker compose build
docker compose up
```

#### Worship Team - current schedule:
```bash
curl -X GET \
  'http://localhost:8080/actions/get-team-schedule?team=worship'
```

#### Worship Team - who's serving this Sunday:
```bash
curl -X GET \
  'http://localhost:8080/actions/get-who-is-serving-this-week?team=worship'
```
