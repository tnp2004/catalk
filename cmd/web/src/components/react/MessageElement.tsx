export const MessageElement = ({ message, role }: MessageHistory) => {
    return (
        <li className={`flex justify-${role === "user" ? "end" : "start"}`}>
            <label
                className={`border-2 border-slate-700 bg-slate-100 shadow-md rounded-xl rounded-${role == "user" ? "br" : "bl"}-sm block py-2 px-3`}
            >{message}</label>
        </li>
    )
}