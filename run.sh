docker run -it --rm \
  -v `pwd`/secrets:/app/secrets:ro \
  -v `pwd`/cache:/app/cache \
  -v `pwd`/masterdata:/app/masterdata \
  inspix-hailstorm:latest
