import { Container } from "./styles";
import { PropsInput } from "../../@types/chat-input";

export default function ChatInput({ onSend }: PropsInput) {
  return (
    <Container>
      <input type="text" onKeyDown={onSend} placeholder="Message" />
    </Container>
  );
}
