import { Container } from "./style";
import { PropsMessage } from "../../@types/message";

export default function Message({ message }: PropsMessage) {
  return <Container>{message.body}</Container>;
}
