import { Container } from "./styles";
import { PropsInput } from "../../@types/chat-input";

export default function ChatInput({ send }: PropsInput) {
  return (
    <Container>
      <input type="text" onKeyDown={send} placeholder="Message" />
    </Container>
  );
}
