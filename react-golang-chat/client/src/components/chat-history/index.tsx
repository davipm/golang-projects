import Message from "../message";

import { Container } from "./styles";
import { PropsHistory } from "../../@types/chat-history";

export default function ChatHistory({ chatHistory }: PropsHistory) {
  return (
    <Container>
      <h2>Chat History</h2>
      {chatHistory.map((msg, index) => {
        const parseMessage = JSON.parse(msg.data);
        return <Message message={parseMessage} key={index} />;
      })}
    </Container>
  );
}
