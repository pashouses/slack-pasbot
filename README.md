# slack-api

General purpose Slack Pashouses Bot running on Google Cloud Run.

## Feeatures
### Shuffle
[Shufflet](https://pastechnologies.slack.com/apps/A020Q5ZF49H-shufflet)-like capability with additional function to exclude yourself from the shuffle result by default.
```
# Input
/shuffle @backend-reviewers
# Output
@Rafel Permata from @backend-reviewers, nominated by elia
```

Assuming you are a member of `@backend-reviewers` and you want to include yourself, you can use
```
# Input
/shuffle @backend-reviewers 1
# Output
@elia from @backend-reviewers, nominated by elia
```

You can also return more than 1 user, like shufflet
```
# Input
/shuffle @backend-reviewers 1 3
# Output
@Rafel Permata, @elia, @Irvan Putra from @backend-reviewers, nominated by elia
```
