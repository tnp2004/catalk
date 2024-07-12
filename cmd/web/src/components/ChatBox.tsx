import { useState, type FormEvent } from "react"
import { MessageElement } from "./MessageElement";

export const ChatBox = ({ breed }: { breed: string }) => {
    const [chat, setChat] = useState<MessageHistory[]>([])
    const [msgInput, setMsgInput] = useState<string>("")

    const SendMessageToAI = async (e: FormEvent) => {
        e.preventDefault();
    
        const headers: Headers = new Headers();
        headers.set("Content-Type", "application/json");
        headers.set("Accept", "application/json");

        const bodyReq: MessageRequest = {
            messageHistory: chat,
            newUserMessage: msgInput,
        };

        const request: RequestInfo = new Request(
            `http://localhost:8080/api/v1/gemini/cats/${breed}`,
            {
                method: "POST",
                headers: headers,
                body: JSON.stringify(bodyReq),
            },
        );

        const res = await fetch(request);
        const bodyRes: ResponseData<MessageResponse> = await res.json();
        setChat(bodyRes.data.newMessageHistory)
        setMsgInput("")
    };

    return (
        <>
            <ul className="h-[60vh] flex flex-col px-2 py-3 gap-2 overflow-auto">
                {chat.map(({message, role}, i) => <MessageElement key={`${role}message-${i+1}`} message={message} role={role} />)}
            </ul>
            
            <form onSubmit={SendMessageToAI} className="flex mb-2">
                <input
                    type="text"
                    value={msgInput}
                    onChange={e => setMsgInput(e.target.value)}
                    className="w-full px-2 rounded-md shadow-sm"
                    placeholder="Type a message . . ."
                    required
                />
                <button className="font-bold p-1" type="submit">send</button>
            </form>
        </>
    )
}