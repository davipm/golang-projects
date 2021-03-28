import { memo } from "react";
import { Container } from "./style";
import { PropsMessage } from "../../@types/message";

function Message({ message }: PropsMessage) {
  return <Container>{message.body}</Container>;
}

export default memo(Message);
