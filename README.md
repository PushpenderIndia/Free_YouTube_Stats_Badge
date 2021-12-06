# Free YouTube Stats Badge
This project will help you to create Live img.shields.io Badges which will Count YouTube Stats (Subscriber, Views, Videos) without YouTube API

## Motive of this project
There is no way to get YouTube Stats (Subscriber Count, Videos Count, Views Count) without using YouTube API, that's why there is no public `img.shield.io` badge which will count Live YouTube Stats for you.

That's why, I've created this project, which will help you to create Live YouTube Stats Badges for absolutely Free, which you can then use on your GitHub page.

## Example Badges

### Urls

| Name                    | Example                                                                                                               |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------- |
| Live Subscriber count   | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-channel-badge.ngoldack.vercel.app/api/subscriber) |
| Live Total Views count  | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-channel-badge.ngoldack.vercel.app/api/views)      | 
| Live Total Videos count | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-channel-badge.ngoldack.vercel.app/api/videos)     | 

### Styles

| Name          | Examples                                                                                                                                  |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| plastic       | ![Custom badge](https://img.shields.io/endpoint?style=plastic&url=https://youtube-channel-badge.ngoldack.vercel.app/api/subscriber)       |
| flat          | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-channel-badge.ngoldack.vercel.app/api/subscriber)                     |
| flat-square   | ![Custom badge](https://img.shields.io/endpoint?style=flat-square&url=https://youtube-channel-badge.ngoldack.vercel.app/api/subscriber)   |
| for-the-badge | ![Custom badge](https://img.shields.io/endpoint?style=for-the-badge&url=https://youtube-channel-badge.ngoldack.vercel.app/api/subscriber) |
| social        | ![Custom badge](https://img.shields.io/endpoint?style=social&url=https://youtube-channel-badge.ngoldack.vercel.app/api/subscriber)        |


## Features
* [X] Caching Functionality to Prevent Rate Limiting
* [X] Using Free API for fetching LIVE YouTube Stats
* [X] Converts Integer (e.g. 10700) to K-Thousand, M-Million, B-Billion, T-Trillion Number (e.g. 10.7K)
* [X] Very Low Latency as Whole Project is Written in GoLang

## Commands For Local Testing
* How to use on Windows Server for testing purposes
```
# Export Cache Time Global Variable
SET CACHE_TIME=10

# Export CHANNEL_ID Global Variable containing your YouTube Channel ID
SET CHANNEL_ID=UCRv-wp0CWXXXXXX3NkTIXXX  

# Running API
go run .

# Visit: http://localhost:8090/subscribers
```

* How to use on Linux Server for testing purposes
```
# Export Cache Time Global Variable
export CACHE_TIME=10

# Export CHANNEL_ID Global Variable containing your YouTube Channel ID
export CHANNEL_ID=UCRv-wp0CWXXXXXX3NkTIXXX

# Running API
go run .

# Visit: http://localhost:8090/subscribers
```

## Deployment