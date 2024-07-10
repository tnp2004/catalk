type MessageRequest = {
    messageHistory: MessageHistory[]
    newUserMessage: string
}

type MessageResponse = {
    newMessageHistory: MessageHistory[]
    responseMessage: string
}

type Roles = "user" | "model"

type MessageHistory = {
    message: string
    role: Roles
}