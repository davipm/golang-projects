import { KeyboardEvent } from "react";

export interface PropsInput {
  send: (event: KeyboardEvent<HTMLInputElement>) => void;
}
