## roadrunner-demo-fiber

## Load tests

go
```bash
artillery quick \
    --count 20 \
    --num 100 \
    http://localhost:8080/api/v2/profile/b1302b2a-795e-4faa-a730-a278e218b72b
```

php
```bash
artillery quick \
    --count 20 \
    --num 100 \
    http://localhost:8080/api/v2/profile/b1302b2a-795e-4faa-a730-a278e218b72b
```


