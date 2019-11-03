# GoQuote

A simple CLI printing the quote of the day via theysaidso.com API

Roadmap:
- (completed) Version flag (-v or --version) to display application version number.
- (completed) Help flag (-h or --help) to display name, usage, version, author, commands, and global options.
- (completed) ~~Use flags for keyword quotes (ex: happy, inspire, funny, love)~~
- Return only the quote and author
- Use flags to return other metadata


## Installation

    go get -u github.com/NathanielRand/goQuote

## Navigate to application folder

    cd /github.com/NathanielRand/goQuote

## Build the application

    go build quote.go

## Global usage

    cp quote /usr/local/bin

## Test default (outside of current project directory):

    quote

## Test with keyword (outside of current project directory):

    quote funny

## Help

    quote -h

## Version

    quote -v


### Warning: Limit 10 uses per hour

    
## Author

* [Nathaniel Rand](https://oneware.io)

## More

* [NathanielRand's Github](https://github.com/NathanielRand/)

## How to Contribute

Follow this guide (preferred):
https://github.com/firstcontributions/first-contributions

or

1. Fork it
2. Create your feature branch (`git checkout -b a-new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin a-new-feature`)
5. Create new Pull Request