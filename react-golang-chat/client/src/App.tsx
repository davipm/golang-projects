import React, {
  useState,
  useEffect,
  useCallback,
  KeyboardEvent,
  ChangeEvent,
} from "react";

import Header from "./components/header";
import ChatHistory from "./components/chat-history";
import ChatInput from "./components/chat-input";
import { connect, sendMsg } from "./service/websocket";

import Global from "./styles/global";
import { HistoryType } from "./@types/chat-history";

type SendMessageType =
  | KeyboardEvent<HTMLInputElement>
  | ChangeEvent<HTMLInputElement>;

export default function App() {
  const [chatHistory, setChatHistory] = useState<HistoryType[]>([]);

  useEffect(() => {
    connect((msg: any) => {
      setChatHistory([...chatHistory, msg]);
    });
  }, [chatHistory]);

  const sendMessage = useCallback((event: SendMessageType) => {
    const { key } = event as KeyboardEvent<HTMLInputElement>;
    const { target } = event as ChangeEvent<HTMLInputElement>;

    if (key === "Enter") {
      sendMsg(target?.value);
      target.value = "";
    }
  }, []);

  return (
    <div id="app">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <ChatInput onSend={sendMessage} />

      <Global />
    </div>
  );
}
