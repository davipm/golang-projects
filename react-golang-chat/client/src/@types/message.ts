type MessageType = {
  type?: number;
  body: string;
};

export interface PropsMessage {
  message: MessageType;
}
