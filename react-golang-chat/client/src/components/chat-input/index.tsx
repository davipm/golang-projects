import { memo } from "react";
import { Container } from "./styles";
import { PropsInput } from "../../@types/chat-input";

function ChatInput({ onSend }: PropsInput) {
  return (
    <Container>
      <input type="text" onKeyDown={onSend} placeholder="Message" />
    </Container>
  );
}

export default memo(ChatInput);
