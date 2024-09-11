## roadrunner-demo-fiber

## Load tests

go only
```bash
artillery quick \
    --count 20 \
    --num 10 \
    http://localhost:8080/api/v2/profile/b1302b2a-795e-4faa-a730-a278e218b72b
```

php only
```bash
artillery quick \
    --count 20 \
    --num 10 \
    http://localhost:8080/api/v1/profile/b1302b2a-795e-4faa-a730-a278e218b72b
```

go with slow downsteam service
```bash
artillery quick \
    --count 20 \
    --num 10 \
    "http://localhost:8080/api/v2/profile/b1302b2a-795e-4faa-a730-a278e218b72b?include-history=true"
```

php with slow downsteam service
```bash
artillery quick \
    --count 20 \
    --num 10 \
    "http://localhost:8080/api/v1/profile/b1302b2a-795e-4faa-a730-a278e218b72b?include-history=true"
```


