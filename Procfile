web: thin start -p 5000
resque-web: resque-web --foreground --no-launch
resque: rake resque:work QUEUE='*' --trace
