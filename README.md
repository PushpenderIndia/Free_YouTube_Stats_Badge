# Free YouTube Stats Badge
This project will help you to create Live img.shields.io Badges which will Count YouTube Stats (Subscriber, Views, Videos) without YouTube API

## No Longer Maintained 

## Motive of this project
There is no way to get YouTube Stats (Subscriber Count, Videos Count, Views Count) without using YouTube API, that's why there is no public `img.shield.io` badge which will count Live YouTube Stats for you.

That's why, I've created this project, which will help you to create Live YouTube Stats Badges for absolutely Free, which you can then use on your GitHub page.

## Example Badges

### Urls

| Name                    | Example                                                                                                     |
| ----------------------- | ------------------------------------------------------------------------------------------------------------|
| Live Subscriber count   | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-badge-deploy.vercel.app/api/subscriber) |
| Live Total Views count  | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-badge-deploy.vercel.app/api/views)      | 
| Live Total Videos count | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-badge-deploy.vercel.app/api/videos)     | 

### Styles

| Name          | Examples                                                                                                                        |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| plastic       | ![Custom badge](https://img.shields.io/endpoint?style=plastic&url=https://youtube-badge-deploy.vercel.app/api/subscriber)       |
| flat          | ![Custom badge](https://img.shields.io/endpoint?url=https://youtube-badge-deploy.vercel.app/api/subscriber)                     |
| flat-square   | ![Custom badge](https://img.shields.io/endpoint?style=flat-square&url=https://youtube-badge-deploy.vercel.app/api/subscriber)   |
| for-the-badge | ![Custom badge](https://img.shields.io/endpoint?style=for-the-badge&url=https://youtube-badge-deploy.vercel.app/api/subscriber) |
| social        | ![Custom badge](https://img.shields.io/endpoint?style=social&url=https://youtube-badge-deploy.vercel.app/api/subscriber)        |


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
It is neccessary to deploy this repo by yourself to get a domain for the desired youtube channel.
An easy and free way to deploy this repo is by using vercel. Just use the button below.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?envLink=https%3A%2F%2Fgithub.com%2FPushpenderIndia%2FFree_YouTube_Stats_Badge%23configuration&envDescription=Find%20information%20on%20how%20to%20get%20these%20in%20the%20readme&env=CHANNEL_ID%2CCACHE_TIME&repository-url=https%3A%2F%2Fgithub.com%2FPushpenderIndia%2FFree_YouTube_Stats_Badge)

### Step by step

1. Click on the above **deploy** button or simply click [here](https://vercel.com/new/clone?envLink=https%3A%2F%2Fgithub.com%2FPushpenderIndia%2FFree_YouTube_Stats_Badge%23configuration&envDescription=Find%20information%20on%20how%20to%20get%20these%20in%20the%20readme&env=CHANNEL_ID%2CCACHE_TIME&repository-url=https%3A%2F%2Fgithub.com%2FPushpenderIndia%2FFree_YouTube_Stats_Badge)

2. Click on **GitHub** Button

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/deploy1.PNG?raw=true)

3. Click on **Add GitHub Namespace** and Authorize vercel

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/deploy2.png?raw=true)

4. After Adding the GitHub Namespace, Select your **Namespace** from the Drop down
    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/deploy3.png?raw=true)

5. Type Any **Repository Name** of your choice & then click on **Create** Button

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/deploy4.png?raw=true)

6. Type your Channel ID and Cache Time, then finally click on **Deploy** Button

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/deploy5.png?raw=true)

## Usage

1. Deploy the repo and get your url from **domains**

    Example in vercel:

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/usage1.png?raw=true)

2. Visit this url: [Shields.io](https://shields.io/endpoint) and scroll down

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/usage2.PNG?raw=true)

3. Add your vercel url to the **url** field
4. Add your wanted url suffix:
    - to get the subscriber count: add api/subscriber
    - to get the view count: add api/views
    - to get the video count: add api/videos

    It should look like this:

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/usage3.png?raw=true)

5. Override the labels/colors/logo if you want (optional)
6. Click on the button to copy the badge url

    ![url image](https://github.com/PushpenderIndia/Free_YouTube_Stats_Badge/blob/main/img/usage4.png?raw=true)



