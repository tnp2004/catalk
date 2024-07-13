export const MessageElement = ({ message, role }: MessageHistory) => {
    return (
        <li className={`flex ${role == "user" ? "justify-end" : "justify-start"}`}>
            <label htmlFor="message"
                className={`border-2 border-slate-700 bg-slate-100 shadow-md max-w-[32rem] rounded-xl ${role == "user" ? "rounded-br-sm" : "rounded-bl-sm"} block py-2 px-3`}
            >{message}</label>
        </li>
    )
}