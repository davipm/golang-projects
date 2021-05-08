import Message from "../message";

import { Container } from "./styles";
import { PropsHistory } from "../../@types/chat-history";

export default function ChatHistory({ chatHistory }: PropsHistory) {
  const parseHistory = chatHistory.map((msg) => JSON.parse(msg.data));

  return (
    <Container>
      <h2>Chat History</h2>
      {parseHistory.map((msg, index) => {
        return <Message message={msg} key={index} />;
      })}
    </Container>
  );
}
