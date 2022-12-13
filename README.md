# slack-api

General purpose Slack Pashouses Bot running on Google Cloud Run.

## Feeatures
### Shuffle
[Shufflet](https://pastechnologies.slack.com/apps/A020Q5ZF49H-shufflet)-like capability with additional function to exclude yourself from the shuffle result by default.
```
# Input
/shuffle @backend-reviewers
# Output
@Rafel Permata (@backend-reviewers), nominated by elia
```

Assuming you are a member of `@backend-reviewers` and you want to include yourself, you can use:
```
# Input
/shuffle @backend-reviewers 1
# Output
@elia (@backend-reviewers), nominated by elia
```

Return more than 1 user, like shufflet:
```
# Input
/shuffle @backend-reviewers 1 3
# Output
@Rafel Permata, @elia, @Irvan Putra (@backend-reviewers), nominated by elia
```

You can also use @here or @channel in the channels where you have invited @pasbot into. It will automatically exclude bots in the channel:
```
/shuffle @here
```
