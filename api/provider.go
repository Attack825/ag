package api

type Provider interface {
    Name() string
    CreateChatCompletion(prompt string, model string, stream bool) (chan string, error)
}

var providers = make(map[string]Provider)

func RegisterProvider(name string, p Provider) {
    providers[name] = p
}

func GetProvider(name string) Provider {
    return providers[name]
}