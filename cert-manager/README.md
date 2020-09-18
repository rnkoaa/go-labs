Cron and Server Service

This simple application demonstrates the ability to run a few go modules together.

1. We can run a server as daemon by calling `./cert-manager serve`. This will also kickoff a background cron job with a configured crontab.

2. Secondly, this allows us to use multiple config files, a `config.yaml` and a `secret.yaml` files together without overriding one another.
