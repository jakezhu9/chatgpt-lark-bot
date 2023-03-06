# ChatGPT Lark Bot

Your personal AI assistant in lark/feishu

![1](assets/1.jpeg)

## How to use

1. Download the file corresponding to your OS in the [releases page](https://github.com/jakezhu9/chatgpt-lark-bot/releases).
2. After you download the file, extract it into a folder and rename `config_example.yaml` to `config.yaml`.
3. Create your bot app in [Lark Open Platform](https://open.larksuite.com/app). Get the`app_id`, `app_secret` in the **Credentials & Basic Info** page, `verification_token` and `event_encrypt_key` in the **Event Subscriptions** page.
4. Edit `config.yaml` and fill in your credentials:
    - `open_ai_key`: [Click here](https://platform.openai.com/account/api-keys) to get your OpenAI key.
5. Run the bot. Go to the **Event Subscriptions** page to configure the Request URL. If you don't have public network, you may need a reverse proxy tool such as [ngrok](https://ngrok.com/download). Add `Message received` event in the **Events added** part.
6. Enable bot in **Features**. Add the following scopes to the bot in the **Permissions & Scopes** page.
    - im:message.group_msg:readonly
    - im:message.p2p_msg:readonly
    - im:message:send_as_bot
7. Release your bot in the **Version Management & Release** page.
8. Enjoy! Welcome star if like it😄



## License

This repository is licensed under the [MIT License](LICENSE).