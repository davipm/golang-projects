import { KeyboardEvent } from "react";

export interface PropsInput {
  onSend: (event: KeyboardEvent<HTMLInputElement>) => void;
}
