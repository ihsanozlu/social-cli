# social-cli

A simple command-line interface (CLI) to post content to **Instagram** (images and reels) using the **Instagram Graph API**.  
Built with [Cobra](https://github.com/spf13/cobra) in Go.  

## ✨ Features
- Configure Instagram App credentials (IG ID, Access Token, App Version).
- Post images to Instagram feed.
- Post videos as **Reels**.
- Manage and inspect configuration (`set`, `get`, `delete`).

## 📦 Installation
Clone and build:

```bash
git clone https://github.com/ihsanozlu/social-cli.git
cd social-cli
go build -o social-cli


### Make it available globally:

```
sh move_social-cli_to_bin.sh
```

## ⚙️ Configuration
Before using, you need to set your credentials:

```
social-cli config set --id <IG_USER_ID> --token <ACCESS_TOKEN> --version v23.0
```

Check current config:
```
social-cli config get
```
Delete config:
```
social-cli config delete
```

## 🚀 Usage

### Post an Image
```
social-cli post --url "https://example.com/image.jpg" --caption "Hello world!"
```
### Post a Reel (video)

```
social-cli post --url "https://example.com/video.mp4" --caption "test 🎥" --type video
```

⚠️ Notes:

For video uploads, Instagram only supports Reels (media_type=REELS).

You must provide a direct video file URL (.mp4), YouTube/Vimeo links will not work.

🤝 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss.


