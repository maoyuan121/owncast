<br />
<p align="center">
  <a href="https://github.com/owncast/owncast" alt="Owncast">
    <img src="https://owncast.online/images/logo.png" alt="Logo" width="200">
  </a>


  <p align="center">
    Take control over your content and stream it yourself.
    <br />
    <a href="http://owncast.online"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://watch.owncast.online/">View Demo</a>
    ·
    <a href="https://broadcast.owncast.online/">Use Our Server for Testing</a>
    ·
    <a href="https://owncast.online/docs/faq/">FAQ</a>
    ·
    <a href="https://github.com/owncast/owncast/issues">Report Bug</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
* [Getting Started](#getting-started)
  * [Getting Started](#getting-started)
  * [Configuration](#configuration)
  * [Web Interface & Chat](#web-interface--chat)
* [Use with your broadcasting software](#use-with-your-existing-broadcasting-software)
* [Video storage and distribution options](#video-storage-options)
* [Building from source](#building-from-source)
* [License](#license)
* [Contact](#contact)


<!-- ABOUT THE PROJECT -->
## About The Project

<p align="center">
  <a href="https://owncast.online/images/owncast-screenshot.png">
    <img src="https://owncast.online/images/owncast-screenshot.png" width="70%">
  </a>
</p>

2020年，当每个人都被困在家里时，世界发生了变化，
人们在他们的卧室里寻找创造性的出口来分享他们的艺术、技能和他们自己。

这种情况使得 Facebook live，YouTube Live Instagram 还有 Twitch 上造成了直播的爆炸式增长，
这些服务提供了他们需要的一切，一种向世界直播的简单方法，让用户成为他们社区的一部分。

那时我想为人们提供更好的选择。你可以自己运行，获得这些服务的所有功能，你可以对观众进行直播，让他们参与到聊天中来，就像他们在其他服务上习惯的那样。
 **There should be a independent, standalone _Twitch in a Box_.**

 
**记住，虽然大型社交公司提供的流媒体服务永远是免费的，你用你的身份和数据，以及每个收听的人的身份和数据来支付。
当你自己举办任何活动时，你就得自己掏钱了。但是，运行一个自托管的流媒体服务器只需每月 5 美元，
这比把你的灵魂卖给 Facebook、谷歌或亚马逊要好得多。**

---

<!-- GETTING STARTED -->

## Getting Started

我们的目标是拥有一个可以运行并开箱即用的服务。
 
**Visit the [Quickstart](https://owncast.online/docs/quickstart/) to get up and running.**

## Configuration

许多方面都可以根据您的喜好进行调整和定制。
[查看 Configuration](https://owncast.online/docs/configuration/) 更新web界面，视频设置等等。

## Web interface + chat

Owncast 包括一个内置聊天的视频 web 界面，一旦你启动服务器就可以使用。

web 界面是专门为任何喜欢调整 web 页面的人而设计的。它没有捆绑或编译成任何东西，它只是 HTML + Javascript + CSS，你可以开始编辑。

请阅读 [web 文档](https://owncast.online/docs/website/) 中提供的特性和如何配置它们的更多信息。


## Use with your existing broadcasting software

一般来说，Owncast 兼容任何使用 “RTMP” 向远程服务器广播的软件。“RTMP” 是所有主要的流媒体服务使用的，
所以如果你目前正在使用其中的一个，你很可能可以将你现有的软件指向你自己的 cast 实例。

OBS、Streamlabs、Restream 和其他许多工具都能与 Owncast 一起使用。[阅读更多关于与现有软件的兼容性](https://owncast.online/docs/broadcasting/)。


## Video storage options

支持两种存储和分发视频的方式。

1. 通过 Owncast 服务器存储在本地。
2. [S3-compatible storage](https://owncast.online/docs/s3/).

### Local file distribution

这是最简单的，可以开箱即用。在这个场景中，视频将从运行服务器的计算机向公众提供。
如果你有一个快速的互联网连接，足够的带宽给你，和较少的观众，这可能对许多人来说是好的。

### S3-Compatible Storage

与其直接从您的个人服务器上提供视频，您可以使用 S3 兼容存储提供商来减轻其他地方的带宽和存储需求。

阅读 [more detailed documentation about configuration of S3-compatible services](https://owncast.online/docs/s3/).


## Building from Source

1. 确保你正确配置了 gcc complier 
1. 安装 [Go toolchain](https://golang.org/dl/)
1. 克隆仓储  `git clone https://github.com/owncast/owncast`
1. 从源代码的根目录运行 `go run main.go pkged.go` ，这个 pkged.go 用于将静态文件打包到 Go 的二进制文件
1. 在您的新服务器上指出你的[广播软件](https://owncast.online/docs/broadcasting/)，并开始流媒体。

这里还提供了一个 `Dockerfile`，所以你可以用很少的努力从源代码启动它。[阅读更多关于从源代码运行](https://owncast.online/docs/building/)。


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Project chat: [Join us on Rocket.Chat](https://owncast.rocket.chat/home) if you want to contribute, follow along, or if you have questions.

Gabe Kangas - [@gabek@mastodon.social](https://mastodon.social/@gabek) - email [gabek@real-ity.com](mailto:gabek@real-ity.com)

Project Link: [https://github.com/owncast/owncast](https://github.com/owncast/owncast)
