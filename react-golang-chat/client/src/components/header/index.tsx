import { memo } from "react";
import { Container } from "./style";

function Header() {
  return (
    <Container>
      <h2>Realtime Chat App</h2>
    </Container>
  );
}

export default memo(Header);
