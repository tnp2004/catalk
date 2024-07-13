import { useState, type FormEvent } from "react"
import { MessageElement } from "../react/MessageElement";

export const ChatBox = ({ breed }: { breed: string }) => {
    const [chat, setChat] = useState<MessageHistory[]>([])
    const [msgInput, setMsgInput] = useState<string>("")
    const { PUBLIC_SERVER_API } = import.meta.env

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
            `${PUBLIC_SERVER_API}/gemini/cats/${breed}`,
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
                {chat.map(({ message, role }, i) => <MessageElement key={`${role}message-${i + 1}`} message={message} role={role} />)}
            </ul>

            <form onSubmit={SendMessageToAI} className="flex mb-2">
                <input
                    type="text"
                    value={msgInput}
                    onChange={e => setMsgInput(e.target.value)}
                    className="w-full px-2 rounded-l-md shadow-sm"
                    placeholder="Type a message . . ."
                    required
                />
                <button className="rounded-r font-bold py-1 px-2 bg-slate-600 hover:bg-slate-700 text-slate-200 hover:text-slate-100" type="submit">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="m22 2-7 20-4-9-9-4Z" /><path d="M22 2 11 13" /></svg>
                </button>
            </form>
        </>
    )
}