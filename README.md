# Dagins

## A command line utility to keep you accountable about your daily habits cultivation

It helps to keep you 100 with yourself.

### Usage

Set _habit creation_ goal:

    $ dagins goal create "learn go" --daily
    # --daily is the default flag

Then check if your bills are paid:

    $ dagins goal status "learn go"
    ➜ Do it!

Go do your thing and let `dagins` know that you were a trooper about it:

    $ dagins goal done "learn go"
    ➜ Way to go. I will be here for you tomorrow.

## Should I use it?

Not yet. :)

## How it works

Information is stored in two JSON files: goals.json and progress.json.

### goals.json

```json
{
  "goals": [
    {
      "name": "learn go",
      "periodicity: "daily"
    }
  ]
}
```

### progress.json

```json
{
  "goals": [
    {
      "name": "learn go",
      "streak": 3,
      "completions": [
        "2020-01-01 4:20:00",
        "2020-01-02 4:20:00",
        "2020-01-06 4:20:00",
        "2020-01-07 4:20:00",
        "2020-01-08 4:20:00",
      ]
    }
  ]
}
```
